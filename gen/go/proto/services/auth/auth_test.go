package auth

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	permsauth "github.com/galexrt/fivenet/gen/go/proto/services/auth/perms"
	"github.com/galexrt/fivenet/internal/tests/proto"
	"github.com/galexrt/fivenet/internal/tests/servers"
	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/config/appconfig"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/mstlystcdata"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/server/audit"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func TestMain(m *testing.M) {
	if err := servers.TestDBServer.Setup(); err != nil {
		fmt.Println("failed to setup mysql test server: %w", err)
		return
	}
	defer servers.TestDBServer.Stop()

	if err := servers.TestNATSServer.Setup(); err != nil {
		fmt.Println("failed to setup nats test server: %w", err)
		return
	}
	defer servers.TestNATSServer.Stop()

	code := m.Run()

	os.Exit(code)
}

func TestFullAuthFlow(t *testing.T) {
	defer servers.TestDBServer.Reset()

	db, err := servers.TestDBServer.DB()
	require.NoError(t, err)

	ctx := context.Background()
	logger := zap.NewNop()
	tp := tracesdk.NewTracerProvider()
	ui := userinfo.NewMockUserInfoRetriever(map[int32]*userinfo.UserInfo{})
	tm := auth.NewTokenMgr("")

	cfg, err := config.LoadBaseConfigTest()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	cfg.NATS.URL = servers.TestNATSServer.GetURL()
	cfg.Cache.RefreshTime = 1 * time.Hour

	appCfg, err := appconfig.LoadTest()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	js, err := servers.TestNATSServer.GetJS()
	require.NoError(t, err)

	fxLC := fxtest.NewLifecycle(t)

	p, err := perms.New(perms.Params{
		LC:        fxLC,
		Logger:    logger,
		DB:        db,
		TP:        tp,
		JS:        js,
		AppConfig: appCfg,
	})
	assert.NoError(t, err)

	aud := &audit.Noop{}

	c, err := mstlystcdata.NewCache(mstlystcdata.Params{
		LC:     fxLC,
		Logger: logger,
		TP:     tp,
		DB:     db,
		Config: cfg,
	})
	assert.NoError(t, err)
	enricher := mstlystcdata.NewEnricher(c, appCfg)

	srv := NewServer(Params{
		Logger:   logger,
		DB:       db,
		Auth:     auth.NewGRPCAuth(ui, tm),
		TM:       tm,
		Perms:    p,
		Enricher: enricher,
		Aud:      aud,
		UI:       ui,
		Config:   cfg,
	})

	fxLC.RequireStart()
	defer fxLC.RequireStop()

	client, _, cancel := NewTestAuthServiceClient(srv)
	defer cancel()

	// First login without credentials
	loginReq := &LoginRequest{}
	loginReq.Username = ""
	loginReq.Password = ""
	res, err := client.Login(ctx, loginReq)
	assert.Error(t, err)
	assert.Nil(t, res)
	proto.CompareGRPCError(t, ErrInvalidLogin, err)

	// Login with invalid credentials
	loginReq.Username = "non-existant-username"
	loginReq.Password = "non-existant-password"
	res, err = client.Login(ctx, loginReq)
	assert.Error(t, err)
	assert.Nil(t, res)
	proto.CompareGRPCError(t, ErrInvalidLogin, err)

	// user-3: Login with valid account that has one char
	loginReq.Username = "user-3"
	loginReq.Password = "password"
	res, err = client.Login(ctx, loginReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	if res == nil {
		assert.FailNow(t, "user-3: Login with valid account failed, response is nil")
	}
	assert.NotEmpty(t, res.GetToken())

	// user-3: Create authenticated metadate and get characters (only has one char)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + res.GetToken()})
	ctx = metadata.NewOutgoingContext(ctx, md)
	getCharsReq := &GetCharactersRequest{}
	getCharsRes, err := client.GetCharacters(ctx, getCharsReq)
	assert.NoError(t, err)
	assert.NotNil(t, getCharsRes)
	if getCharsRes == nil {
		assert.FailNow(t, "user-3: Empty char list returned for valid account that should have 2 chars")
	}
	assert.Len(t, getCharsRes.GetChars(), 1)

	// user-1: Login with valid account (2 chars)
	loginReq.Username = "user-1"
	loginReq.Password = "password"
	res, err = client.Login(ctx, loginReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	if res == nil {
		assert.FailNow(t, "user-1: Login with valid account failed, response is nil")
	}
	assert.NotEmpty(t, res.GetToken())

	// user-1: Create authenticated metadate and get characters
	md = metadata.New(map[string]string{"Authorization": "Bearer " + res.GetToken()})
	ctx = metadata.NewOutgoingContext(ctx, md)
	getCharsReq = &GetCharactersRequest{}
	getCharsRes, err = client.GetCharacters(ctx, getCharsReq)
	assert.NoError(t, err)
	assert.NotNil(t, getCharsRes)
	if getCharsRes == nil {
		assert.FailNow(t, "user-1: Empty char list returned for valid account that should have 2 chars")
	}
	assert.Len(t, getCharsRes.GetChars(), 2)

	// user-1: Choose an invalid character
	chooseCharReq := &ChooseCharacterRequest{}
	chooseCharReq.CharId = 2 // Char id 2 is `user-2`'s char
	chooseCharRes, err := client.ChooseCharacter(ctx, chooseCharReq)
	assert.Error(t, err)
	assert.Nil(t, chooseCharRes)
	proto.CompareGRPCError(t, ErrUnableToChooseChar, err)

	role, err := p.GetRoleByJobAndGrade(ctx, "ambulance", 1)
	assert.NoError(t, err)
	assert.NotNil(t, role)

	perm, err := p.GetPermission(ctx, permsauth.AuthServicePerm, permsauth.AuthServiceChooseCharacterPerm)
	assert.NoError(t, err)
	assert.NotNil(t, perm)

	// user-1: Choose valid character but we don't have permissions on the role
	err = p.RemovePermissionsFromRole(ctx, role.ID, perm.Id)
	assert.NoError(t, err)
	// Disable choose char perm from ambulance rank 1 role, `user-1` is a medic rank 1+
	err = p.UpdateRolePermissions(ctx, role.ID, perms.AddPerm{
		Id:  perm.Id,
		Val: false,
	})
	assert.NoError(t, err)
	chooseCharReq.CharId = 1
	chooseCharRes, err = client.ChooseCharacter(ctx, chooseCharReq)
	assert.Error(t, err)
	assert.Nil(t, chooseCharRes)
	proto.CompareGRPCError(t, ErrUnableToChooseChar, err)

	// user-1: Choose valid character, now we add a permssion
	err = p.UpdateRolePermissions(ctx, role.ID, perms.AddPerm{
		Id:  perm.Id,
		Val: true,
	})
	assert.NoError(t, err)
	chooseCharReq.CharId = 1
	chooseCharRes, err = client.ChooseCharacter(ctx, chooseCharReq)
	assert.NoError(t, err)
	assert.NotNil(t, chooseCharRes)
}
