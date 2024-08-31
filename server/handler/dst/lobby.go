package dst

import (
	"cmp"
	"context"
	"github.com/dstgo/lobby/pkg/geo"
	"github.com/dstgo/lobby/pkg/lobbyapi"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/types/dst"
	"github.com/dstgo/lobby/server/utils/maputil"
	"github.com/ginx-contribs/ginx/pkg/resp/statuserr"
	"net"
	"slices"
)

func NewLobbyHandler(serverRepo *repo.ServerRepo, client *lobbyapi.Client) *LobbyHandler {
	return &LobbyHandler{serverRepo: serverRepo, client: client}
}

type LobbyHandler struct {
	client     *lobbyapi.Client
	serverRepo *repo.ServerRepo
}

func (l *LobbyHandler) SearchByPage(ctx context.Context, options dst.SearchOptions) (dst.QueryListResult, error) {
	options.Size = min(options.Size, 100)
	list, total, err := l.serverRepo.PageQueryByOption(ctx, options)
	if err != nil {
		return dst.QueryListResult{}, statuserr.InternalError(err)
	}
	var servers []dst.ServerInfo
	for _, e := range list {
		servers = append(servers, dst.EntServerToServerInfo(e))
	}
	return dst.QueryListResult{Total: total, List: servers}, nil
}

func (l *LobbyHandler) LatestVersion(ctx context.Context) (int, error) {
	version, err := l.serverRepo.QueryLatestVersion(ctx)
	if err != nil {
		return 0, statuserr.InternalError(err)
	}
	return version, nil
}

func (l *LobbyHandler) GetServerDetails(ctx context.Context, region, rowId string) (dst.QueryDetailsResult, error) {
	serverDetails, err := l.client.GetServerDetails(region, rowId)
	if err != nil {
		return dst.QueryDetailsResult{}, err
	}
	serverInfo := dst.LobbyServerToServerInfo(serverDetails.Server)
	ipAddress, err := geo.City(net.ParseIP(serverInfo.Address))
	if err != nil {
		return dst.QueryDetailsResult{}, err
	}
	serverInfo.CountryCode = ipAddress.Country.IsoCode
	serverInfo.Country = maputil.GetFbMap("zh-CN", "en", ipAddress.Country.Names)
	serverInfo.City = maputil.GetFbMap("zh-CN", "en", ipAddress.City.Names)
	serverInfo.Continent = maputil.GetFbMap("zh-CN", "en", ipAddress.Continent.Names)
	if serverInfo.Platform == "WeGame" {
		serverInfo.CountryCode = "CN"
		serverInfo.Continent = "亚洲"
		serverInfo.Country = "中国"
	}
	return dst.QueryDetailsResult{
		ServerInfo: serverInfo,
		MetaInfo:   serverDetails.Details,
	}, nil
}

// CreateServersBatch creates a list of servers in n batch
func (l *LobbyHandler) CreateServersBatch(ctx context.Context, servers []*ent.Server, batchSize int) (int64, error) {
	// sort by country code
	slices.SortFunc(servers, func(a, b *ent.Server) int {
		return cmp.Compare(a.CountryCode, b.CountryCode)
	})

	var created int64
	for start := 0; start < len(servers); start += batchSize {
		end := start + batchSize
		if end > len(servers) {
			end = len(servers)
		}
		createdBatch, err := l.serverRepo.CreateBulk(ctx, servers[start:end])
		if err != nil {
			return 0, statuserr.InternalError(err)
		}
		created += createdBatch
	}
	return created, nil
}
