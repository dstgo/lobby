package user

import (
	"github.com/dstgo/lobby/server/types/route"
	"github.com/ginx-contribs/ginx"
)

type Router struct {
	User *UserAPI
}

func NewRouter(root *ginx.RouterGroup, userApi *UserAPI) Router {

	userGroup := root.Group("/user")
	userGroup.MGET("/info", ginx.M{route.Private}, userApi.Info)
	userGroup.MGET("/list", ginx.M{route.Private}, userApi.List)

	return Router{User: userApi}
}
