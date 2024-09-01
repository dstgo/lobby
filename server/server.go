package server

import (
	"context"
	"fmt"
	"github.com/dstgo/lobby/server/conf"
	authhandler "github.com/dstgo/lobby/server/handler/auth"
	"github.com/dstgo/lobby/server/mids"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/server/pkg/logh"
	"github.com/dstgo/lobby/server/types"
	"github.com/dstgo/size"
	"github.com/ginx-contribs/ginx"
	"github.com/ginx-contribs/ginx/constant/methods"
	"github.com/ginx-contribs/ginx/contribs/requestid"
	"github.com/ginx-contribs/ginx/middleware"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/ginx-contribs/ent-sqlite"
)

// NewApp returns a new http app server
func NewApp(ctx context.Context, appConf *conf.App) (*ginx.Server, error) {

	slog.Debug("lobby server is initializing")

	// initialize database
	slog.Debug(fmt.Sprintf("connecting to %s(%s)", appConf.DB.Driver, appConf.DB.Address))
	db, err := InitializeDB(ctx, appConf.DB)
	if err != nil {
		return nil, err
	}

	// initialize redis client
	slog.Debug(fmt.Sprintf("connecting to redis(%s)", appConf.Redis.Address))
	redisClient, err := InitializeRedis(ctx, appConf.Redis)
	if err != nil {
		return nil, err
	}

	// initialize email client
	slog.Debug(fmt.Sprintf("establish email client(%s:%d)", appConf.Email.Host, appConf.Email.Port))
	emailClient, err := InitializeEmail(ctx, appConf.Email)
	if err != nil {
		return nil, err
	}

	// initialize http client
	var lobbyClient *lobbyapi.Client
	if appConf.Dst.ProxyUrl != "" {
		lobbyClient = lobbyapi.NewWith(appConf.Dst.KeliToken, resty.New().SetProxy(appConf.Dst.ProxyUrl))
	} else {
		lobbyClient = lobbyapi.NewWith(appConf.Dst.KeliToken,
			resty.NewWithClient(&http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}))
	}

	// initialize ginx server
	server := ginx.New(
		ginx.WithOptions(ginx.Options{
			Mode:               gin.ReleaseMode,
			Address:            appConf.Server.Address,
			ReadTimeout:        appConf.Server.ReadTimeout,
			WriteTimeout:       appConf.Server.WriteTimeout,
			IdleTimeout:        appConf.Server.IdleTimeout,
			MaxMultipartMemory: appConf.Server.MultipartMax,
			MaxHeaderBytes:     int(size.MB * 2),
			MaxShutdownTimeout: time.Second * 5,
		}),
		ginx.WithNoMethod(middleware.NoMethod(methods.Get, methods.Post, methods.Put, methods.Delete, methods.Options)),
		ginx.WithNoRoute(middleware.NoRoute()),
		ginx.WithMiddlewares(
			// reocvery handler
			middleware.Recovery(slog.Default(), nil),
			// request id
			requestid.RequestId(),
			// access logger
			middleware.Logger(slog.Default(), "request-log"),
			// rate limit by counting
			mids.RateLimitByCount(redisClient, appConf.Limit.Public.Limit, appConf.Limit.Public.Window, mids.ByIpPath),
			// jwt authentication
			mids.TokenAuthenticator(authhandler.NewTokenHandler(appConf.Jwt, redisClient)),
		),
	)

	// set validator for gin
	err = setupHumanizedValidator()
	if err != nil {
		return nil, err
	}

	// whether to enable pprof program profiling
	if appConf.Server.Pprof {
		server.Engine().GET("/pprof/profile", gin.WrapF(pprof.Profile))
		server.Engine().GET("/pprof/heap", gin.WrapH(pprof.Handler("heap")))
		server.Engine().GET("/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		slog.Info("pprof profiling enabled")
	}

	tc := types.Context{
		AppConf: appConf,
		Ent:     db,
		Redis:   redisClient,
		Router:  server.RouterGroup().Group("/api"),
		Email:   emailClient,
		Lobby:   lobbyClient,
	}
	slog.Debug("setup api router")

	// initialize api router
	sc, err := setup(tc)
	if err != nil {
		return nil, err
	}

	// register cron job
	cronJob, err := InitializeCronJob(ctx, tc, sc)
	if err != nil {
		return nil, err
	}
	cronJob.Start()

	// register shutdown hook
	onShutdown := func(ctx context.Context) error {
		cronJob.Stop()
		logh.ErrorNotNil("db closed failed", db.Close())
		logh.ErrorNotNil("redis closed failed", redisClient.Close())
		logh.ErrorNotNil("email client closed failed", emailClient.Close())
		return nil
	}
	server.OnShutdown = append(server.OnShutdown, onShutdown)

	return server, nil
}
