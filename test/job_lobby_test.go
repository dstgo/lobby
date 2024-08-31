package test

import (
	"context"
	"github.com/dstgo/lobby/pkg/lobbyapi"
	"github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/handler/dst"
	"github.com/dstgo/lobby/server/job"
	"github.com/dstgo/lobby/test/testuitl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newLobbyCollectJob() (*job.LobbyCollectJob, error) {
	ctx := context.Background()
	dbConf, err := testuitl.ReadDBConf()
	if err != nil {
		return nil, err
	}
	db, err := server.InitializeDB(ctx, dbConf)
	if err != nil {
		return nil, err
	}
	serverRepo := repo.NewServerRepo(db)
	client := lobbyapi.New("")
	handler := dst.NewLobbyHandler(serverRepo, client)
	return job.NewLobbyCollectJob(handler, client), nil
}

func TestLobbyCollect(t *testing.T) {
	collectJob, err := newLobbyCollectJob()
	if !assert.NoError(t, err) {
		return
	}
	collected, cost, err := collectJob.Collect(1, 10)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(len(collected), cost)
}

func TestLobbyCollectBatch(t *testing.T) {
	collectJob, err := newLobbyCollectJob()
	if !assert.NoError(t, err) {
		return
	}
	collectJob.CollectBatch(10, 1000)
}
