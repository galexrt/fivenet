package auth

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/galexrt/arpanet/api"
	"github.com/galexrt/arpanet/model"
	"github.com/galexrt/arpanet/pkg/auth"
	"github.com/galexrt/arpanet/pkg/helpers"
	"github.com/galexrt/arpanet/pkg/session"
	"github.com/galexrt/arpanet/query"
	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	AccountServiceServer
}

func NewServer() *Server {
	return &Server{}
}

// AuthFuncOverride is called instead of exampleAuthFunc
func (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// Skip authentication for the login endpoint
	if fullMethodName == "/gen.auth.AccountService/Login" {
		return ctx, nil
	}

	return auth.GRPCAuthFunc(ctx)
}

func (s *Server) createTokenForAccount(account *model.Account, activeCharIdentifier string) (string, error) {
	return session.Tokens.NewWithClaims(&session.CitizenInfoClaims{
		AccountID:  account.ID,
		Username:   account.Username,
		ActiveChar: activeCharIdentifier,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "arpanet",
			Subject:   account.License,
			ID:        strconv.FormatUint(uint64(account.ID), 10),
			Audience:  []string{"arpanet"},
		},
	})
}

func (s *Server) getAccountFromDB(userID string) (*model.Account, error) {
	accounts := query.Account
	account, err := accounts.Where(accounts.Enabled.Is(true), accounts.Username.Eq(userID)).Limit(1).First()
	if err != nil {
		return nil, err

	}
	return account, nil
}

func (s *Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	resp := &LoginResponse{}
	account, err := s.getAccountFromDB(req.Username)
	if err != nil {
		return nil, err
	}

	if !account.CheckPassword(req.Password) {
		return &LoginResponse{}, errors.New("wrong username or password")
	}

	token, err := s.createTokenForAccount(account, "")
	if err != nil {
		return nil, err
	}
	resp.Token = token

	return resp, nil
}

func (s *Server) GetCharacters(ctx context.Context, req *GetCharactersRequest) (*GetCharactersResponse, error) {
	resp := &GetCharactersResponse{}

	claims, err := session.Tokens.ParseWithClaims(auth.MustGetTokenFromGRPCContext(ctx))
	if err != nil {
		return resp, nil
	}

	// Load chars and add them to the response
	chars, err := api.Auth.GetCharsByLicense(claims.Subject)
	if err != nil {
		return resp, err
	}

	for _, char := range chars {
		resp.Chars = append(resp.Chars, helpers.ConvertModelUserToCommonCharacter(char))
	}

	return resp, nil
}

func (s *Server) ChooseCharacter(ctx context.Context, req *ChooseCharacterRequest) (*ChooseCharacterResponse, error) {
	resp := &ChooseCharacterResponse{}

	claims, err := session.Tokens.ParseWithClaims(auth.MustGetTokenFromGRPCContext(ctx))
	if err != nil {
		return resp, nil
	}

	account, err := s.getAccountFromDB(claims.Username)
	if err != nil {
		return nil, err
	}

	token, err := s.createTokenForAccount(account, req.Identifier)
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return resp, nil
}

func (s *Server) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	// TODO till we have a JWT token manager "blocking" users when they logout, nothing todo here
	resp := &LogoutResponse{
		Success: true,
	}
	return resp, nil
}
