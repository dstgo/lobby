package test

import (
	"context"
	"github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/handler/dst"
	"github.com/dstgo/lobby/server/jobs"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/test/testuitl"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newLobbyCollectJob() (*jobs.LobbyCollectJob, error) {
	ctx := context.Background()
	cfg, err := testuitl.ReadConf()
	if err != nil {
		return nil, err
	}
	db, err := server.InitializeDB(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}
	serverRepo := repo.NewServerRepo(db.Debug())
	client := lobbyapi.New(cfg.Dst.KeliToken)
	handler := dst.NewLobbyHandler(serverRepo, client)
	return jobs.NewLobbyCollectJob(handler, client, cfg.Job.Collect), nil
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

func TestLobbyCron(t *testing.T) {
	cronjob := jobs.NewCronJob()
	collectJob, err := newLobbyCollectJob()
	if !assert.NoError(t, err) {
		return
	}
	err = cronjob.AddJob(collectJob)
	if !assert.NoError(t, err) {
		return
	}
	cronjob.Start()
	defer cronjob.Stop()
	select {
	case <-time.After(time.Second * 1):
	}
}
