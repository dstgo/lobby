package system

import (
	"github.com/dstgo/lobby/server/types/route"
	"github.com/ginx-contribs/ginx"
)

type Router struct {
	System *SystemAPI
}

func NewRouter(root *ginx.RouterGroup, systemAPI *SystemAPI) Router {
	// test api
	root.MGET("/ping", ginx.M{route.Public}, systemAPI.Ping)
	root.MGET("/pong", ginx.M{route.Private}, systemAPI.Pong)

	return Router{System: systemAPI}
}
