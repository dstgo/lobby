package test

import (
	"context"
	"github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/handler/dst"
	"github.com/dstgo/lobby/server/jobs"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/test/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newLobbyCollectJob() (*jobs.LobbyCollectJob, error) {
	ctx := context.Background()
	cfg, err := testutil.ReadConf()
	if err != nil {
		return nil, err
	}
	db, err := server.InitializeDB(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}
	// due to large number of records, do not enable debug sql logging.
	serverRepo := repo.NewServerRepo(db)
	client := lobbyapi.New(cfg.Dst.KeliToken)
	handler := dst.NewLobbyHandler(serverRepo, client)
	return jobs.NewLobbyCollectJob(handler, client, cfg.Job.Collect), nil
}

func newLobbyCleanJob() (*jobs.LobbyCleanJob, error) {
	ctx := context.Background()
	cfg, err := testutil.ReadConf()
	if err != nil {
		return nil, err
	}
	db, err := server.InitializeDB(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}
	// due to large number of records, do not enable debug sql logging.
	serverRepo := repo.NewServerRepo(db)
	client := lobbyapi.New(cfg.Dst.KeliToken)
	handler := dst.NewLobbyHandler(serverRepo, client)
	return jobs.NewLobbyCleanJob(handler, cfg.Job.Clean), nil
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

func TestLobbyCollectCron(t *testing.T) {
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
	case <-time.After(time.Minute * 5):
	}
}

func TestLobbyClean(t *testing.T) {
	job, err := newLobbyCleanJob()
	if !assert.NoError(t, err) {
		return
	}
	cmd := job.Cmd()
	cmd()
}
