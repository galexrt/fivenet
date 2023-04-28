package oauth2

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/galexrt/fivenet/pkg/auth"
	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/oauth2/providers"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var (
	accounts      = table.FivenetAccounts
	oauthAccounts = table.FivenetOauth2Accounts
)

type OAuth2 struct {
	logger *zap.Logger
	db     *sql.DB
	tm     *auth.TokenManager

	oauthConfigs map[string]providers.IProvider
}

func New(logger *zap.Logger, db *sql.DB, tm *auth.TokenManager) *OAuth2 {
	o := &OAuth2{
		logger:       logger,
		db:           db,
		tm:           tm,
		oauthConfigs: make(map[string]providers.IProvider, len(config.C.OAuth2.Providers)),
	}

	for _, p := range config.C.OAuth2.Providers {
		cfg := &oauth2.Config{
			RedirectURL:  p.RedirectURL,
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			Scopes:       p.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:   p.Endpoints.AuthURL,
				TokenURL:  p.Endpoints.TokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
			},
		}
		var provider providers.IProvider
		switch p.Type {
		case config.OAuth2ProviderDiscord:
			provider = &providers.Discord{
				BaseProvider: providers.BaseProvider{
					Name: p.Name,
				},
			}
		default:
			provider = &providers.Generic{
				BaseProvider: providers.BaseProvider{
					Name:          p.Name,
					UserInfoURL:   p.Endpoints.UserInfoURL,
					DefaultAvatar: p.DefaultAvatar,
				},
			}
		}

		provider.SetOauthConfig(cfg)
		provider.SetMapping(p.Mapping)

		o.oauthConfigs[p.Name] = provider
	}

	return o
}

func (o *OAuth2) GetProvider(c *gin.Context) (providers.IProvider, error) {
	param, ok := c.Params.Get("provider")
	if !ok {
		return nil, fmt.Errorf("no provider found")
	}

	provider, ok := o.oauthConfigs[param]
	if !ok {
		return nil, fmt.Errorf("no provider found")
	}

	return provider, nil
}

const (
	AccountInfoRedirBase string = "/auth/account-info"
	LoginRedirBase       string = "/auth/login"
)

func (o *OAuth2) handleRedirect(c *gin.Context, err error, connectOnly bool, success bool, message string) {
	if !success {
		redirURL := ""
		if connectOnly {
			redirURL = AccountInfoRedirBase + "?oauth-connect=failed"
		} else {
			redirURL = LoginRedirBase + "?oauth-login=failed"
		}

		if message != "" {
			redirURL = redirURL + "&message=" + url.QueryEscape(message)
		}

		c.Redirect(http.StatusTemporaryRedirect, redirURL)
		return
	}

	redirURL := ""
	if connectOnly {
		redirURL = AccountInfoRedirBase + "?oauth-connect=success"
	} else {
		redirURL = LoginRedirBase + "?oauth-login=success"
	}

	c.Redirect(http.StatusTemporaryRedirect, redirURL)
}

func (o *OAuth2) Login(c *gin.Context) {
	sess := sessions.DefaultMany(c, "fivenet_oauth2_state")
	connectOnly := false
	if c.Request.Method == http.MethodPost {
		connectOnlyVal := c.Request.PostFormValue("connect-only")
		if connectOnlyVal != "" {
			var err error
			connectOnly, err = strconv.ParseBool(connectOnlyVal)
			if err != nil {
				o.logger.Error("failed to parse connect only var", zap.Error(err))
				o.handleRedirect(c, err, false, false, "invalid_request")
				return
			}
		}

		tokenVal := c.Request.PostFormValue("token")
		if tokenVal != "" {
			sess.Set("token", tokenVal)
		}
	}

	state, err := utils.GenerateRandomString(64)
	if err != nil {
		o.handleRedirect(c, err, connectOnly, false, "internal_error")
		return
	}

	sess.Set("connect-only", connectOnly)
	sess.Set("state", state)
	sess.Save()

	provider, err := o.GetProvider(c)
	if err != nil {
		o.logger.Error("failed to get provider", zap.Error(err))
		o.handleRedirect(c, err, connectOnly, false, "invalid_provider")
		return
	}

	// Redirect to provider for login
	c.Redirect(http.StatusTemporaryRedirect, provider.GetRedirect(state))
}

func (o *OAuth2) Callback(c *gin.Context) {
	sess := sessions.DefaultMany(c, "fivenet_oauth2_state")
	sessState := sess.Get("state")
	if sessState == nil {
		o.handleRedirect(c, nil, false, false, "invalid_state")
		return
	}

	state := sessState.(string)
	connectOnly := false
	sessConnectOnly := sess.Get("connect-only")
	if sessConnectOnly != nil {
		connectOnly = sessConnectOnly.(bool)
	}

	var token string
	sessToken := sess.Get("token")
	if sessToken != nil {
		token = sessToken.(string)
	}

	// Remove vars from session
	sess.Delete("connect-only")
	sess.Delete("state")
	sess.Delete("token")
	sess.Save()

	if c.Request.FormValue("state") != state {
		o.handleRedirect(c, nil, connectOnly, false, "invalid_state")
		return
	}

	provider, err := o.GetProvider(c)
	if err != nil {
		o.logger.Error("failed to get provider", zap.Error(err))
		o.handleRedirect(c, nil, connectOnly, false, "invalid_provider")
		return
	}

	userInfo, err := provider.GetUserInfo(c.Request.FormValue("code"))
	if err != nil {
		o.logger.Error("failed to get userinfo from provider", zap.Error(err))
		o.handleRedirect(c, err, connectOnly, false, "provider_failed")
		return
	}

	if connectOnly {
		claims, err := o.tm.ParseWithClaims(token)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth-login=failed&reason=token_invalid")
			return
		}

		if err := o.storeUserInfo(c, claims.AccountID, provider.GetName(), userInfo); err != nil {
			o.handleRedirect(c, err, connectOnly, false, "internal_error")
			return
		}

		o.handleRedirect(c, nil, connectOnly, true, "")
		return
	}

	// Take care of logging user in
	account, err := o.getUserInfo(c, provider.GetName(), userInfo)
	if err != nil {
		o.logger.Error("failed to store userinfo in database", zap.Error(err))
		o.handleRedirect(c, err, connectOnly, false, "internal_error")
		return
	}

	if account == nil {
		c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth-login=failed&reason=unconnected")
		return
	} else if account.ID == 0 {
		o.handleRedirect(c, nil, connectOnly, true, "internal_error")
		return
	}

	newToken, err := o.tm.NewWithClaims(auth.BuildTokenClaimsFromAccount(account, nil))
	if err != nil {
		o.handleRedirect(c, err, connectOnly, true, "internal_error")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth-connect=success&token="+url.QueryEscape(newToken))
}

func (o *OAuth2) getUserInfo(ctx context.Context, provider string, userInfo *providers.UserInfo) (*model.FivenetAccounts, error) {
	stmt := oauthAccounts.
		SELECT(
			accounts.ID,
			accounts.Username,
			accounts.License,
		).
		FROM(oauthAccounts.
			INNER_JOIN(accounts,
				accounts.ID.EQ(oauthAccounts.AccountID),
			),
		).
		WHERE(jet.AND(
			oauthAccounts.Provider.EQ(jet.String(provider)),
			oauthAccounts.ExternalID.EQ(jet.Uint64(uint64(userInfo.ID))),
			accounts.Enabled.IS_TRUE(),
		)).
		LIMIT(1)

	var dest model.FivenetAccounts
	if err := stmt.QueryContext(ctx, o.db, &dest); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, err
		}
		return nil, nil
	}

	return &dest, nil
}

func (o *OAuth2) storeUserInfo(ctx context.Context, accountId uint64, provider string, userInfo *providers.UserInfo) error {
	stmt := oauthAccounts.
		INSERT(
			oauthAccounts.AccountID,
			oauthAccounts.Provider,
			oauthAccounts.ExternalID,
			oauthAccounts.Username,
			oauthAccounts.Avatar,
		).
		VALUES(
			accountId,
			provider,
			userInfo.ID,
			userInfo.Username,
			userInfo.Avatar,
		)

	if _, err := stmt.ExecContext(ctx, o.db); err != nil {
		return err
	}

	return nil
}
