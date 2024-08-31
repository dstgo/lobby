package types

import (
	"github.com/dstgo/lobby/server/conf"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/ginx-contribs/ginx"
	"github.com/ginx-contribs/ginx/constant/status"
	"github.com/ginx-contribs/ginx/pkg/resp/statuserr"
	"github.com/redis/go-redis/v9"
	"github.com/wneessen/go-mail"
)

// Response it is only used for documentation, use package 'ginx/resp' to build response.
type Response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Error string `json:"error"`
}

// Context holds all dependent context
type Context struct {
	// app configuration
	AppConf *conf.App
	// ent client
	Ent *ent.Client
	// redis client
	Redis *redis.Client
	// app router
	Router *ginx.RouterGroup
	// email client
	Email *mail.Client
	// lobbyapi client
	Lobby *lobbyapi.Client
}

// custom code is composed of three parts: Order_Status_Code, it will be shown in the response body.
// Order just represents order of package create time, it is used to avoid duplicates error code in different packages.
// Status represents the error will be occurred in which situation, it is corresponds to http status.
// Code is the true error code, whose max capacity is 999.
const customCode = 0_000_000

var (
	ErrBadParams = statuserr.Errorf("bad parameters").SetCode(400_001).SetStatus(status.BadRequest)

	ErrInternal = statuserr.Errorf("internal server error").SetCode(500_000).SetStatus(status.InternalServerError)
)
