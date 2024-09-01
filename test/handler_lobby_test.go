package test

import (
	"context"
	"github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/handler/dst"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/test/testutil"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newLobbyHandler() (*dst.LobbyHandler, error) {
	ctx := context.Background()
	appconf, err := testutil.ReadConf()
	if err != nil {
		return nil, err
	}
	db, err := server.InitializeDB(ctx, appconf.DB)
	if err != nil {
		return nil, err
	}
	serverRepo := repo.NewServerRepo(db)
	client := lobbyapi.NewWith(appconf.Dst.KeliToken, resty.New().SetProxy("http://127.0.0.1:7890"))
	return dst.NewLobbyHandler(serverRepo, client), nil
}

func TestLobbyDetails(t *testing.T) {
	ctx := context.Background()
	handler, err := newLobbyHandler()
	if !assert.NoError(t, err) {
		return
	}
	details, err := handler.GetServerDetails(ctx, lobbyapi.ApSoutheast, "KU_CJH79WSu")
	if !assert.NoError(t, err) {
		return
	}
	t.Logf("%+v\n", details)
}

func TestLobbyVersion(t *testing.T) {
	ctx := context.Background()
	handler, err := newLobbyHandler()
	if !assert.NoError(t, err) {
		return
	}
	version, err := handler.LatestVersion(ctx)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(version)
}

func TestDeleteInBatch(t *testing.T) {
	ctx := context.Background()
	handler, err := newLobbyHandler()
	if !assert.NoError(t, err) {
		return
	}
	deleted, err := handler.DeleteServerBatch(ctx, time.Hour, 2000)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(deleted)
}
