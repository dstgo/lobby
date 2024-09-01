package svc

import (
	"github.com/dstgo/lobby/server/api"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/handler/auth"
	"github.com/dstgo/lobby/server/handler/dst"
	"github.com/dstgo/lobby/server/handler/email"
	"github.com/dstgo/lobby/server/handler/job"
	"github.com/dstgo/lobby/server/handler/user"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(Context), "*"),
)

// Context represents holds all handler and repo instances, just for helper
type Context struct {
	// api
	ApiRouter api.Router

	// dst
	LobbyHandler *dst.LobbyHandler
	ServerRepo   *repo.ServerRepo

	// user
	UserHandler *user.UserHandler
	UserRepo    *repo.UserRepo

	// system
	AuthHandler *auth.AuthHandler

	// email
	EmailHandler *email.Handler

	// job
	JobHandler *job.JobHandler
	JobRepo    *repo.JobRepo
}
