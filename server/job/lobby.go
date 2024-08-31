package job

import (
	"context"
	"fmt"
	"github.com/dstgo/lobby/pkg/geo"
	"github.com/dstgo/lobby/pkg/lobbyapi"
	"github.com/dstgo/lobby/pkg/ts"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/handler/dst"
	dstype "github.com/dstgo/lobby/server/types/dst"
	"github.com/dstgo/lobby/server/utils/maputil"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func NewLobbyCollectJob(handler *dst.LobbyHandler, client *lobbyapi.Client) *LobbyCollectJob {
	return &LobbyCollectJob{handler: handler, client: client}
}

type LobbyCollectJob struct {
	handler *dst.LobbyHandler
	client  *lobbyapi.Client

	count atomic.Int64
}

func (l *LobbyCollectJob) Name() string {
	return "lobby-Collect"
}

func (l *LobbyCollectJob) Cron() string {
	// collect data every 2 minutes
	return "*/2 * * * *"
}

func (l *LobbyCollectJob) Cmd() func() {
	return func() {
		l.CollectBatch(10, 1000)
	}
}

func (l *LobbyCollectJob) CollectBatch(limit, batch int) {
	qv := ts.UnixMicro()
	servers, cost, err := l.Collect(qv, limit)
	if err != nil {
		slog.Error(fmt.Sprintf("%s-%d failed", l.Name(), l.count.Load()), slog.Any("error", err))
		return
	}
	collected, err := l.handler.CreateServersBatch(context.Background(), servers, batch)
	if err != nil {
		slog.Error(fmt.Sprintf("%s-%d batch failed", l.Name(), l.count.Load()), slog.Any("error", err))
	} else {
		slog.Info(fmt.Sprintf("%s-%d ok", l.Name(), l.count.Load()),
			slog.Int64("servers", collected), slog.Duration("cost", cost))
	}
	l.count.Add(1)
}

// Collect collects servers data from klei lobby server
func (l *LobbyCollectJob) Collect(v int64, limit int) (collected []*ent.Server, cost time.Duration, err error) {
	start := ts.Now()
	regions, err := l.client.GetCapableRegions()
	if err != nil {
		return nil, 0, err
	}
	slog.Debug("getting regions ok", slog.Duration("cost", ts.Now().Sub(start)))

	var collectServers []*ent.Server
	var mu sync.Mutex

	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(limit)

	for _, region := range regions.Regions {
		for _, platform := range lobbyapi.ExplicitPlatforms {
			// Collect servers concurrently
			group.Go(func() error {
				gstart := ts.Now()
				servers, err := l.getLobbyServers(region.Region, platform, v)
				gcost := ts.Now().Sub(gstart)
				if err != nil {
					slog.Debug(fmt.Sprintf("%s-%s failed", region.Region, platform), slog.Duration("cost", gcost))
					return err
				}
				mu.Lock()
				collectServers = append(collectServers, servers...)
				mu.Unlock()
				slog.Debug(fmt.Sprintf("%s-%s ok", region.Region, platform),
					slog.Int64("collected", int64(len(servers))), slog.Duration("cost", gcost))
				return nil
			})
		}
	}

	err = group.Wait()
	if err != nil {
		return nil, 0, err
	}
	cost = ts.Now().Sub(start)
	return collectServers, cost, nil
}

func (l *LobbyCollectJob) getLobbyServers(region string, platform string, qv int64) ([]*ent.Server, error) {
	servers, err := l.client.GetLobbyServers(region, platform)
	if err != nil {
		return nil, err
	}
	return l.ProcessServers(qv, servers.List)
}

// ProcessServers converts lobbyapi.Server to *ent.Server
func (l *LobbyCollectJob) ProcessServers(qv int64, servers []lobbyapi.Server) ([]*ent.Server, error) {
	var entServers []*ent.Server
	for _, server := range servers {
		createdServer := dstype.LobbyServerToEntServer(server)
		createdServer.QueryVersion = qv
		createdServer.Level = len(server.Secondaries) + 1
		// process tag str
		var tags []*ent.Tag
		for _, t := range server.Tags {
			tags = append(tags, &ent.Tag{Value: t})
		}
		createdServer.Edges.Tags = tags
		// process secondary
		var secondaries []*ent.Secondary
		for _, secondary := range server.Secondaries {
			secondaries = append(secondaries, &ent.Secondary{
				Sid:     secondary.Id,
				SteamID: secondary.SteamId,
				Address: secondary.Address,
				Port:    secondary.Port,
			})
		}
		createdServer.Edges.Secondaries = secondaries
		// process geo info
		ipAddress, err := geo.City(net.ParseIP(server.Address))
		if err != nil {
			return nil, err
		}

		// geo process
		createdServer.CountryCode = ipAddress.Country.IsoCode
		createdServer.Country = maputil.GetFbMap("zh-CN", "en", ipAddress.Country.Names)
		createdServer.City = maputil.GetFbMap("zh-CN", "en", ipAddress.City.Names)
		createdServer.Continent = maputil.GetFbMap("zh-CN", "en", ipAddress.Continent.Names)
		if createdServer.Platform == "WeGame" {
			createdServer.CountryCode = "CN"
			createdServer.Continent = "亚洲"
			createdServer.Country = "中国"
		}

		entServers = append(entServers, createdServer)
	}
	return entServers, nil
}
