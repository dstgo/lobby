package dst

import (
	"github.com/dstgo/lobby/server/handler/dst"
	dstype "github.com/dstgo/lobby/server/types/dst"
	"github.com/gin-gonic/gin"
	"github.com/ginx-contribs/ginx"
	"github.com/ginx-contribs/ginx/pkg/resp"
)

func NewLobbyAPI(lobbyHandler *dst.LobbyHandler) *LobbyAPI {
	return &LobbyAPI{lobbyHandler: lobbyHandler}
}

type LobbyAPI struct {
	lobbyHandler *dst.LobbyHandler
}

// Search
// @Summary      Search
// @Description  return a list of servers filtered by search parameters
// @Tags         dst/lobby
// @Accept       json
// @Produce      json
// @Param        SearchOptions  query dstype.SearchOptions  true "SearchOptions"
// @Success      200  {object}  types.Response{data=dstype.QueryListResult}
// @Router       /lobby/search [GET]
func (l *LobbyAPI) Search(ctx *gin.Context) {
	var opt dstype.SearchOptions
	if err := ginx.ShouldValidateQuery(ctx, &opt); err != nil {
		return
	}
	result, err := l.lobbyHandler.SearchByPage(ctx, opt)
	if err != nil {
		resp.Fail(ctx).Error(err).JSON()
	} else {
		resp.Ok(ctx).Data(result).JSON()
	}
}

// Details
// @Summary      Details
// @Description  returns details information about the server
// @Tags         dst/lobby
// @Accept       json
// @Produce      json
// @Param        QueryDetailsOptions  query  dstype.QueryDetailsOptions  true "QueryDetailsOptions"
// @Success      200  {object}  types.Response{data=dstype.QueryDetailsResult}
// @Router       /lobby/info [GET]
func (l *LobbyAPI) Details(ctx *gin.Context) {
	var opt dstype.QueryDetailsOptions
	if err := ginx.ShouldValidateQuery(ctx, &opt); err != nil {
		return
	}
	details, err := l.lobbyHandler.GetServerDetails(ctx, opt.Region, opt.RowID)
	if err != nil {
		resp.Fail(ctx).Error(err).JSON()
	} else {
		resp.Ok(ctx).Data(details).JSON()
	}
}

// Version
// @Summary      Version
// @Description  return latest version of server
// @Tags         dst/lobby
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response{data=string}
// @Router       /lobby/version [GET]
func (l *LobbyAPI) Version(ctx *gin.Context) {
	version, err := l.lobbyHandler.LatestVersion(ctx)
	if err != nil {
		resp.Fail(ctx).Error(err).JSON()
	} else {
		resp.Ok(ctx).Data(version).JSON()
	}
}
