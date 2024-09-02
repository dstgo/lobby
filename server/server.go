package server

import (
	"context"
	"fmt"
	"github.com/dstgo/lobby/server/conf"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/server/pkg/logh"
	"github.com/dstgo/lobby/server/types"
	_ "github.com/ginx-contribs/ent-sqlite"
	"github.com/ginx-contribs/ginx"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"net/http"
)

// NewApp returns a new http app server
func NewApp(ctx context.Context, appConf *conf.App) (*ginx.Server, error) {
	// initialize database
	slog.Debug(fmt.Sprintf("connecting to %s(%s)", appConf.DB.Driver, appConf.DB.Address))
	db, err := NewDBClient(ctx, appConf.DB)
	if err != nil {
		return nil, err
	}

	// initialize redis client
	slog.Debug(fmt.Sprintf("connecting to redis(%s)", appConf.Redis.Address))
	redisClient, err := NewRedisClient(ctx, appConf.Redis)
	if err != nil {
		return nil, err
	}

	// initialize email client
	slog.Debug(fmt.Sprintf("establish email client(%s:%d)", appConf.Email.Host, appConf.Email.Port))
	emailClient, err := NewEmailClient(ctx, appConf.Email)
	if err != nil {
		return nil, err
	}

	// initialize lobby client
	slog.Debug("initialize lobby client")
	var lobbyClient *lobbyapi.Client
	if appConf.Dst.ProxyUrl != "" {
		lobbyClient = lobbyapi.NewWith(appConf.Dst.KeliToken, resty.New().SetProxy(appConf.Dst.ProxyUrl))
	} else {
		lobbyClient = lobbyapi.NewWith(appConf.Dst.KeliToken,
			resty.NewWithClient(&http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}))
	}

	tc := types.Context{
		AppConf: appConf,
		Ent:     db,
		Redis:   redisClient,
		Email:   emailClient,
		Lobby:   lobbyClient,
	}

	// initialize ginx server
	server, err := NewHttpServer(ctx, appConf, tc)
	if err != nil {
		return nil, err
	}
	tc.Router = server.RouterGroup().Group(appConf.Server.BasePath)
	slog.Debug("setup api router")

	// initialize api router
	sc, err := setup(tc)
	if err != nil {
		return nil, err
	}

	// register cron job
	cronJob, err := NewCronJob(ctx, tc, sc)
	if err != nil {
		return nil, err
	}
	started := cronJob.Start()
	slog.Info(fmt.Sprintf("started %d cron jobs", started))

	// shutdown hook
	onShutdown := func(ctx context.Context) error {
		slog.Info(fmt.Sprintf("stopped %d jobs", cronJob.Stop()))
		logh.ErrorNotNil("db closed failed", db.Close())
		logh.ErrorNotNil("redis closed failed", redisClient.Close())
		logh.ErrorNotNil("email client closed failed", emailClient.Close())
		return nil
	}
	server.OnShutdown = append(server.OnShutdown, onShutdown)

	return server, nil
}
