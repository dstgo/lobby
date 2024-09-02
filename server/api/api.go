package api

import (
	"github.com/dstgo/lobby/server/api/auth"
	"github.com/dstgo/lobby/server/api/dst"
	"github.com/dstgo/lobby/server/api/job"
	"github.com/dstgo/lobby/server/api/system"
	"github.com/dstgo/lobby/server/api/user"
	"github.com/google/wire"
)

// RegisterRouter
// @title	                        Lobby HTTP API
// @version		                    v0.0.0-Beta
// @description                     This is lobby swagger generated api documentation, know more information about lobby on GitHub.
// @contact.name                    dstgo
// @contact.url                     https://github.com/dstgo/lobby
// @BasePath	                    /api/
// @license.name                    MIT LICENSE
// @license.url                     https://mit-license.org/
// @securityDefinitions.apikey      BearerAuth
// @in                              header
// @name                            Authorization
//
//go:generate swag init --ot yaml --generatedTime -g api.go -d ./,../types,../pkg --output ./ && swag fmt -g api.go -d ./

type Router struct {
	Auth   auth.Router
	System system.Router
	User   user.Router
	Dst    dst.Router
	Job    job.Router
}

var Provider = wire.NewSet(
	// auth router
	auth.NewAuthAPI,
	auth.NewRouter,
	// system router
	system.NewSystemAPI,
	system.NewRouter,
	// user router
	user.NewUserAPI,
	user.NewRouter,
	// dst router
	dst.NewLobbyAPI,
	dst.NewRouter,

	// job router
	job.NewJobAPI,
	job.NewRouter,

	// build Router struct
	wire.Struct(new(Router), "*"),
)
