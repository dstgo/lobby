package test

import (
	"context"
	"github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/types"
	"github.com/dstgo/lobby/test/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func newServerRepo() (*repo.ServerRepo, error) {
	ctx := context.Background()
	dbConf, err := testutil.ReadDBConf()
	if err != nil {
		return nil, err
	}
	db, err := server.NewDBClient(ctx, dbConf)
	if err != nil {
		return nil, err
	}
	return repo.NewServerRepo(db.Debug()), nil
}

func newServerEsRepo() (*repo.ServerEsRepo, error) {
	ctx := context.Background()
	cfg, err := testutil.ReadConf()
	if err != nil {
		return nil, err
	}
	elastic, err := server.NewElasticClient(ctx, cfg.Elastic)
	if err != nil {
		return nil, err
	}
	return repo.NewServerEsRepo(elastic), nil
}

func TestServerRepoCreate(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}
	_, err = serverRepo.CreateBulk(ctx, []*ent.Server{
		{Name: "server1", Edges: ent.ServerEdges{Tags: []*ent.Tag{{Value: "t1"}, {Value: "t2"}}, Secondaries: []*ent.Secondary{{Address: "a1"}, {Address: "a2"}}}},
		{Name: "server2", Edges: ent.ServerEdges{Tags: []*ent.Tag{{Value: "t3"}, {Value: "t4"}}, Secondaries: []*ent.Secondary{{Address: "a3"}, {Address: "a4"}}}},
	})
	assert.NoError(t, err)
}

func TestMaxQV(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}
	qv, err := serverRepo.MaxQV(ctx)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(qv)
}

func TestLatestVersion(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}
	version, err := serverRepo.QueryLatestVersion(ctx)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(version)
}

func TestPageQueryByOptions(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}
	list, total, err := serverRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{})
	if !assert.NoError(t, err) {
		return
	}
	t.Log(total)
	for _, e := range list {
		t.Log(e)
	}
}

func TestPageQueryByOptionsPage(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}

	type pagePair struct {
		page int
		size int

		expectedSize int
	}

	sample := []pagePair{
		{page: 1, size: 10, expectedSize: 10},
		{page: 2, size: 20, expectedSize: 20},
		{page: 1000, size: 10, expectedSize: 10},
		{page: 1, size: 1000, expectedSize: 1000},
	}
	for _, pair := range sample {
		list, total, err := serverRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{
			Page: pair.page,
			Size: pair.size,
		})
		if !assert.NoError(t, err) {
			return
		}
		t.Log(pair.page, pair.size, len(list), total)
		assert.Equal(t, pair.expectedSize, len(list))
	}
}

func TestPageQueryByOptionsWithMatch(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}

	samples := []string{
		"abc",
		"你好",
		"世界",
		"糖糕",
		"Aoksad",
	}

	for i, match := range samples {
		page := 1
		size := 20
		list, total, err := serverRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{
			Page:  page,
			Size:  size,
			Match: match,
		})
		if !assert.NoError(t, err) {
			return
		}
		t.Logf("#%d %s", i, match)
		t.Logf("total: %d page: %d size: %d list:%d", total, page, size, len(list))
		for _, e := range list {
			assert.Containsf(t, strings.ToLower(e.Name), strings.ToLower(match), "%s should containsFold %s", e.Name, match)
			t.Log(e)
		}
	}
}

func TestPageQueryByOptionsWithTags(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}

	samples := [][]string{
		{"洞穴", "简体中文"},
	}

	for i, tags := range samples {
		page := 1
		size := 20
		list, total, err := serverRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{
			Page: page,
			Size: size,
			Tags: tags,
		})
		if !assert.NoError(t, err) {
			return
		}
		t.Logf("#%d", i)
		t.Logf("total: %d page: %d size: %d list:%d", total, page, size, len(list))
		for _, e := range list {
			t.Log(e)
		}
	}
}

func TestPageQueryByOptionsWithServerType(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}

	samples := []types.DstServerType{
		types.TypeOfficial,
		types.TypeSteamClan,
		types.TypeSteamClanOnly,
		types.TypeFriendOnly,
		types.TypeClientHosted,
		types.TypeDedicated,
	}

	for i, serverType := range samples {
		page := 1
		size := 100
		list, total, err := serverRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{
			Page:       page,
			Size:       size,
			ServerType: serverType,
		})
		if !assert.NoError(t, err) {
			return
		}
		t.Logf("#%d - %d", i, serverType)
		t.Logf("total: %d page: %d size: %d list:%d", total, page, size, len(list))
		for _, e := range list {
			t.Log(e)
		}
	}
}

func TestPageQueryByOptionsWithPlatform(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}

	samples := []types.ServerPlatform{
		types.PlatformSteam,
		types.PlatformWeGame,
		types.PlatformSwitch,
		types.PlatformPSN,
		types.PlatformXBOne,
	}

	for i, platform := range samples {
		page := 1
		size := 100
		list, total, err := serverRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{
			Page:     page,
			Size:     size,
			Platform: platform,
		})
		if !assert.NoError(t, err) {
			return
		}
		t.Logf("#%d - %s", i, platform)
		t.Logf("total: %d page: %d size: %d list:%d", total, page, size, len(list))
		for _, e := range list {
			assert.Equal(t, e.Platform, platform.String())
			t.Log(e)
		}
	}
}

func TestPageQueryByOptionsWithManyCondition(t *testing.T) {
	ctx := context.Background()
	serverRepo, err := newServerRepo()
	if !assert.NoError(t, err) {
		return
	}

	samples := []types.LobbyServerSearchOptions{
		{Page: 1, Size: 100, Sort: types.DstSortByLevel, Desc: true},
		{Page: 1, Size: 100, Sort: types.DstSortByVersion, Desc: true},
		{Page: 1, Size: 100, Sort: types.DstSortByOnline, Desc: true},
		{Page: 1, Size: 100, Address: "45.74.14.148"},
		{Page: 1, Size: 100, Season: "spring"},
		{Page: 1, Size: 100, Platform: types.PlatformSteam, Season: "summer", CountryCode: "CN", ModEnabled: 1},
		{Page: 1, Size: 100, Level: 6},
	}

	for i, option := range samples {
		list, total, err := serverRepo.PageQueryByOption(ctx, option)
		if !assert.NoError(t, err) {
			return
		}
		t.Log()
		t.Logf("#%d------total: %d page: %d size: %d list:%d", i, total, option.Page, option.Size, len(list))
		for _, e := range list {
			t.Log(e)
		}
	}
}

func TestEsServerRepoMaxQV(t *testing.T) {
	ctx := context.Background()
	esRepo, err := newServerEsRepo()
	if !assert.NoError(t, err) {
		return
	}
	qv, err := esRepo.MaxQv(ctx)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(qv)
}

func TestEsServerRepoTotalCount(t *testing.T) {
	ctx := context.Background()
	esRepo, err := newServerEsRepo()
	if !assert.NoError(t, err) {
		return
	}
	qv, err := esRepo.TotalCount(ctx)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(qv)
}

func TestEsServerRepoPageQuery(t *testing.T) {
	ctx := context.Background()
	esRepo, err := newServerEsRepo()
	if !assert.NoError(t, err) {
		return
	}
	list, total, err := esRepo.PageQueryByOption(ctx, types.LobbyServerSearchOptions{
		Page:     10,
		Size:     100,
		Platform: types.PlatformSteam,
	})
	if !assert.NoError(t, err) {
		return
	}
	t.Log(total)
	for _, e := range list {
		t.Log(e)
	}
}
