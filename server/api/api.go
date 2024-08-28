package api

import (
	"github.com/dstgo/lobby/server/api/auth"
	"github.com/google/wire"
)

type Router struct {
	Auth auth.Router
}

var Provider = wire.NewSet(
	// auth router
	auth.NewAuthAPI,
	auth.NewRouter,
	// system router

	// build Router struct
	wire.Struct(new(Router), "*"),
)

// RegisterRouter
// @title	                        MaxWell HTTP API
// @version		                    v0.0.0Beta
// @description                     This is lobby swagger generated api documentation, know more information about lobby on GitHub.
// @contact.name                    dstgo
// @contact.url                     https://github.com/dstgo/maxwell
// @BasePath	                    /api/
// @license.name                    MIT LICENSE
// @license.url                     https://mit-license.org/
// @securityDefinitions.apikey      BearerAuth
// @in                              header
// @name                            Authorization
//
//go:generate swag init --ot yaml --generatedTime -g api.go -d ./,../types --output ./ && swag fmt -g api.go -d ./
