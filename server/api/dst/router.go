package dst

import (
	"github.com/ginx-contribs/ginx"
)

type Router struct {
	Lobby *LobbyAPI
}

func NewRouter(root *ginx.RouterGroup, lobbyApi *LobbyAPI) Router {

	serverGroup := root.Group("/lobby")
	serverGroup.GET("/search", lobbyApi.Search)
	serverGroup.GET("/info", lobbyApi.Details)
	serverGroup.GET("/version", lobbyApi.Version)

	return Router{}
}
