package dst

import (
	"context"
	"fmt"
	"github.com/dstgo/lobby/server/conf"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/pkg/geo"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/server/pkg/maputil"
	"github.com/dstgo/lobby/server/pkg/ts"
	"github.com/dstgo/lobby/server/types"
	"github.com/dstgo/lobby/test/testutil"
	"github.com/ginx-contribs/ginx/pkg/resp/statuserr"
	"log/slog"
	"net"
	"time"
)

func NewLobbyHandler(serverRepo *repo.ServerRepo, client *lobbyapi.Client, esRepo *repo.ServerEsRepo, esConf conf.Elasticsearch) *LobbyHandler {
	return &LobbyHandler{serverRepo: serverRepo, client: client, esRepo: esRepo, esConf: esConf}
}

type LobbyHandler struct {
	client     *lobbyapi.Client
	serverRepo *repo.ServerRepo
	esRepo     *repo.ServerEsRepo
	esConf     conf.Elasticsearch
}

func (l *LobbyHandler) SearchByPage(ctx context.Context, options types.LobbyServerSearchOptions) (types.LobbyServerSearchResult, error) {
	options.Size = min(options.Size, 100)
	var (
		list  []*ent.Server
		total int
		err   error
	)

	if l.esConf.Enabled {
		list, total, err = l.esRepo.PageQueryByOption(ctx, options)
	} else {
		list, total, err = l.serverRepo.PageQueryByOption(ctx, options)
	}

	if err != nil {
		return types.LobbyServerSearchResult{}, statuserr.InternalError(err)
	}

	// process
	var servers []types.LobbyServerInfo
	for _, e := range list {
		servers = append(servers, types.EntServerToServerInfo(e))
	}
	return types.LobbyServerSearchResult{Total: total, List: servers}, nil
}

func (l *LobbyHandler) LatestVersion(ctx context.Context) (int, error) {
	version, err := l.serverRepo.QueryLatestVersion(ctx)
	if err != nil {
		return 0, statuserr.InternalError(err)
	}
	return version, nil
}

func (l *LobbyHandler) GetServerDetails(ctx context.Context, region, rowId string) (types.LobbyServerDetails, error) {
	serverDetails, err := l.client.GetServerDetails(region, rowId)
	if err != nil {
		return types.LobbyServerDetails{}, err
	}
	serverInfo := types.LobbyServerToServerInfo(serverDetails.Server)
	ipAddress, err := geo.City(net.ParseIP(serverInfo.Address))
	if err != nil {
		return types.LobbyServerDetails{}, err
	}
	serverInfo.CountryCode = ipAddress.Country.IsoCode
	serverInfo.Country = maputil.GetFallBack("zh-CN", "en", ipAddress.Country.Names)
	serverInfo.City = maputil.GetFallBack("zh-CN", "en", ipAddress.City.Names)
	serverInfo.Continent = maputil.GetFallBack("zh-CN", "en", ipAddress.Continent.Names)
	if serverInfo.Platform == types.PlatformWeGame.String() {
		serverInfo.CountryCode = "CN"
		serverInfo.Continent = "亚洲"
		serverInfo.Country = "中国"
	}
	return types.LobbyServerDetails{
		LobbyServerInfo: serverInfo,
		MetaInfo:        serverDetails.Details,
	}, nil
}

// CreateServersBatch creates a list of servers in n batch
func (l *LobbyHandler) CreateServersBatch(ctx context.Context, servers []*ent.Server, batchSize int) (int64, error) {
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

// DeleteServerBatch delete a list of servers in n batch, filtered by query-version - duration.
func (l *LobbyHandler) DeleteServerBatch(ctx context.Context, duration time.Duration, batchSize int) (int64, error) {
	sum := int64(0)
	expiredTs := ts.Now().Add(-duration).UnixMicro()
	slog.Debug("delete server batch beginning", slog.Int("batch-size", batchSize), slog.Int64("expiredTs", expiredTs))
	r := &testutil.Round{}
	t := &testutil.Timer{}
	for {
		t.Start()
		ids, err := l.serverRepo.ExpiredRecords(ctx, expiredTs, batchSize)
		if err != nil {
			return 0, err
		}
		deleted, err := l.serverRepo.DeleteBulk(ctx, ids...)
		if err != nil {
			return 0, err
		}
		slog.Debug(fmt.Sprintf("delete round #%d", r.Round()), slog.Duration("cost", t.Stop()), slog.Int("deleted", deleted))
		// has been finished
		if len(ids) == 0 {
			break
		}
		sum += int64(deleted)
		t.Reset()
	}
	return sum, nil
}
