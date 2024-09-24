package repo

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/ent/secondary"
	"github.com/dstgo/lobby/server/data/ent/server"
	"github.com/dstgo/lobby/server/data/ent/tag"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/server/types"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
	"strings"
)

// NewServerRepo returns a new instance of ServerRepo
func NewServerRepo(client *ent.Client) *ServerRepo {
	return &ServerRepo{Ent: client}
}

type ServerRepo struct {
	Ent *ent.Client
}

// CreateBulk creates servers in batches and its associations relationships
func (s *ServerRepo) CreateBulk(ctx context.Context, servers []*ent.Server) (int64, error) {
	tx, err := s.Ent.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// create servers
	var serverBulk []*ent.ServerCreate
	for _, server := range servers {
		serverBulk = append(serverBulk, tx.Server.Create().
			SetServer(server))
	}
	// save servers
	savedServers, err := tx.Server.CreateBulk(serverBulk...).Save(ctx)
	if err != nil {
		return 0, err
	}

	var secondaryBulk []*ent.SecondaryCreate
	var tagBulk []*ent.TagCreate
	for i, server := range servers {
		// create tags and secondaries
		for _, tag := range server.Edges.Tags {
			tagBulk = append(tagBulk, tx.Tag.Create().SetTag(tag).SetOwnerID(savedServers[i].ID))
		}
		for _, secondary := range server.Edges.Secondaries {
			secondaryBulk = append(secondaryBulk, tx.Secondary.Create().SetSecondary(secondary).SetOwnerID(savedServers[i].ID))
		}
	}

	// save tags and secondaries
	_, err = tx.Tag.CreateBulk(tagBulk...).Save(ctx)
	if err != nil {
		return 0, err
	}
	_, err = tx.Secondary.CreateBulk(secondaryBulk...).Save(ctx)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int64(len(savedServers)), nil
}

func (s *ServerRepo) MaxQV(ctx context.Context) (int64, error) {
	maxQV, err := s.Ent.Server.Query().
		Aggregate(ent.Max(server.FieldQueryVersion)).
		Int(ctx)
	if err != nil {
		return 0, err
	}
	return int64(maxQV), nil
}

// PageQueryByOption returns a list of servers that match the given query options
func (s *ServerRepo) PageQueryByOption(ctx context.Context, options types.LobbyServerSearchOptions) (list []*ent.Server, total int, err error) {
	// select the latest query version
	query := s.Ent.Server.Query()
	qv, err := s.MaxQV(ctx)
	if err != nil {
		return
	}
	if options.Qv != 0 {
		qv = options.Qv
	}

	// count total size
	query = query.Where(server.QueryVersionEQ(qv))
	count, err := query.Count(ctx)
	if err != nil {
		return
	}

	// revise the pagination options
	if options.Size <= 0 {
		options.Size = 10
	}

	if options.Page <= 0 {
		options.Page = 1
	} else if options.Page > count/options.Size && count%options.Page != 0 { // the last page
		options.Page = count/options.Size + 1
	} else if options.Page > count/options.Size && count%options.Page == 0 {
		options.Page = count / options.Size
	}

	// main pattern match
	if options.Match != "" {
		query = query.Where(server.NameContainsFold(options.Match))
	}

	if options.Address != "" {
		query = query.Where(server.AddressContainsFold(options.Address))
	}

	// geo es_search
	if options.Country != "" {
		query = query.Where(server.CountryCodeContainsFold(options.Country))
	}
	if options.CountryCode != "" {
		query = query.Where(server.CountryCode(options.CountryCode))
	}
	if options.Continent != "" {
		query = query.Where(server.ContinentContainsFold(options.Continent))
	}
	if options.Platform != types.PlatformAny {
		query = query.Where(server.Platform(options.Platform.String()))
	}

	// server type
	switch options.ServerType {
	case types.TypeDedicated:
		query = query.Where(server.Dedicated(true))
	case types.TypeClientHosted:
		query = query.Where(server.ClientHosted(true))
	case types.TypeOfficial:
		query = query.Where(server.Or(
			server.NameContains("Public EsServer"),
			server.NameContains("Klei Official"),
			server.HostIn("KU_KleiServ", "KU_KleiServ", "KU_zHxWQDSW"),
		))
	case types.TypeSteamClan:
		query = query.Where(server.SteamClanIDNEQ(""))
	case types.TypeSteamClanOnly:
		query = query.Where(server.SteamClanIDNEQ(""), server.ClanOnly(true))
	case types.TypeFriendOnly:
		query = query.Where(server.FriendOnly(true))
	}

	// game options
	if options.Season != "" {
		query = query.Where(server.Season(options.Season))
	}
	if options.GameMode != "" {
		query = query.Where(server.GameMode(options.GameMode))
	}
	if options.Intent != "" {
		query = query.Where(server.Intent(options.Intent))
	}
	if options.Level > 0 {
		query = query.Where(server.LevelGTE(options.Level))
	}
	if options.PvpEnabled != 0 {
		query = query.Where(server.PvpEQ(options.PvpEnabled == 1))
	}
	if options.ModEnabled != 0 {
		query = query.Where(server.ModEQ(options.ModEnabled == 1))
	}
	if options.HasPassword != 0 {
		query = query.Where(server.PasswordEQ(options.HasPassword == 1))
	}

	// tags filter
	if len(options.Tags) > 0 {
		query = query.Where(server.HasTagsWith(tag.ValueIn(options.Tags...)))
	}

	// pagination
	query = query.Offset((options.Page - 1) * options.Size).Limit(options.Size)

	// sort
	orderOpt := func(opt *sql.OrderTermOptions) {
		opt.Desc = options.Desc
	}
	switch options.Sort {
	case types.DstSortByName:
		query = query.Order(server.ByName(orderOpt))
	case types.DstSortByCountry:
		query = query.Order(server.ByCountry(orderOpt))
	case types.DstSortByVersion:
		query = query.Order(server.ByVersion(orderOpt))
	case types.DstSortByOnline:
		query = query.Order(server.ByOnline(orderOpt))
	case types.DstSortByLevel:
		query = query.Order(server.ByLevel(orderOpt))
	default:
		query = query.Order(server.ByID(orderOpt))
	}

	// execute the query
	total, err = query.Count(ctx)
	if err != nil {
		return
	}
	list, err = query.All(ctx)
	return
}

// QueryLatestVersion query the latest version of lobby
func (s *ServerRepo) QueryLatestVersion(ctx context.Context) (int, error) {
	maxVersion, err := s.Ent.Server.Query().
		Where(server.PlatformEQ(lobbyapi.Steam.String())).
		Aggregate(ent.Max(server.FieldVersion)).Int(ctx)
	if err != nil {
		return 0, err
	}
	return maxVersion, nil
}

// DeleteByQV deletes servers by specified query version
func (s *ServerRepo) DeleteByQV(ctx context.Context, qv int64) error {
	tx, err := s.Ent.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dlServers, err := tx.Server.Query().Where(server.QueryVersionEQ(qv)).All(ctx)
	if err != nil {
		return err
	}
	var ownids []int
	for _, dlServer := range dlServers {
		ownids = append(ownids, dlServer.ID)
	}

	// delete all associated tags
	_, err = tx.Tag.Delete().Where(tag.OwnerIDIn(ownids...)).Exec(ctx)
	if err != nil {
		return err
	}

	// delete all associated secondaries
	_, err = tx.Secondary.Delete().Where(secondary.OwnerIDIn(ownids...)).Exec(ctx)
	if err != nil {
		return err
	}

	// delete servers
	_, err = tx.Server.Delete().Where(server.IDIn(ownids...)).Exec(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// ExpiredRecords retrurns a list id of expired records, use limit to return
func (s *ServerRepo) ExpiredRecords(ctx context.Context, ts int64, limit int) ([]int, error) {
	ids, err := s.Ent.Server.Query().
		Where(server.QueryVersionLTE(ts)).
		Order(ent.Desc(server.FieldQueryVersion)).
		Limit(limit).IDs(ctx)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// DeleteBulk deletes all records in given query versions.
// not recommended delete a large mount of records in one time, it will block for long time.
func (s *ServerRepo) DeleteBulk(ctx context.Context, ids ...int) (int, error) {
	tx, err := s.Ent.Tx(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	// delete all associated tags
	tags, err := tx.Tag.Query().
		Where(
			tag.HasServersWith(server.IDIn(ids...)),
		).IDs(ctx)
	if err != nil {
		return 0, err
	}
	_, err = tx.Tag.Delete().Where(tag.IDIn(tags...)).Exec(ctx)
	if err != nil {
		return 0, err
	}

	// delete all associated secondaries
	secondaries, err := tx.Secondary.Query().
		Where(secondary.HasServersWith(
			server.IDIn(ids...),
		)).IDs(ctx)
	if err != nil {
		return 0, err
	}
	_, err = tx.Secondary.Delete().Where(secondary.IDIn(secondaries...)).Exec(ctx)
	if err != nil {
		return 0, err
	}

	deleted, err := tx.Server.Delete().Where(server.IDIn(ids...)).Exec(ctx)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func NewServerEsRepo(elastic *elasticsearch.Client) *ServerEsRepo {
	return &ServerEsRepo{Elastic: elastic}
}

type ServerEsRepo struct {
	Elastic *elasticsearch.Client
}

func (s *ServerEsRepo) MaxQv(ctx context.Context) (int64, error) {
	dsl := types.H{
		"aggs": types.H{
			"max_qv": types.H{"max": types.H{"field": "data.query_version"}}},
	}
	// get the latest query version
	_, body, err := esSearch(s.Elastic, ctx, []string{"lobby-servers"}, dsl)
	if err != nil {
		return 0, err
	}
	path := "aggregations.max_qv.value"
	maxQv := gjson.GetBytes(body, path)
	if !maxQv.Exists() {
		return 0, fmt.Errorf("missing field in elastic es_search resp: %s", path)
	}
	return maxQv.Int(), nil
}

// TotalCount return total count of document without any condition
func (s *ServerEsRepo) TotalCount(ctx context.Context) (int64, error) {
	dsl := types.H{
		"_source": types.S{"data.id"},
		"aggs": types.H{
			"total_count": types.H{"value_count": types.H{"field": "data.id"}}},
	}
	_, body, err := esSearch(s.Elastic, ctx, []string{"lobby-servers"}, dsl)
	if err != nil {
		return 0, err
	}
	path := "aggregations.total_count.value"
	maxQv := gjson.GetBytes(body, path)
	if !maxQv.Exists() {
		return 0, fmt.Errorf("missing field in elastic es_search resp: %s", path)
	}
	return maxQv.Int(), nil
}

func (s *ServerEsRepo) PageQueryByOption(ctx context.Context, options types.LobbyServerSearchOptions) (list []*ent.Server, total int, err error) {
	qv, err := s.MaxQv(ctx)
	if err != nil {
		return nil, 0, err
	}
	if options.Qv != 0 {
		qv = options.Qv
	}

	count, err := s.TotalCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	// revise the pagination options
	if options.Size <= 0 {
		options.Size = 10
	}

	if options.Page <= 0 {
		options.Page = 1
	} else if int64(options.Page) > count/int64(options.Size) && count%int64(options.Page) != 0 { // the last page
		options.Page = int(count/int64(options.Size) + 1)
	} else if int64(options.Page) > count/int64(options.Size) && count%int64(options.Page) == 0 {
		options.Page = int(count / int64(options.Size))
	}

	// limit for elasticsearch
	if options.Page*options.Size >= 10000 {
		options.Page = (10000-options.Size)/options.Size + 1
	}

	// query dsl
	dsl := types.H{
		"from": (options.Page - 1) * options.Size,
		"size": options.Size,
		"query": types.H{
			"bool": types.H{
				"must": types.S{
					types.H{
						"match": types.H{
							"data.query_version": qv,
						},
					},
				},
			},
		},
	}

	// match section
	condition := dsl["query"].(types.H)["bool"].(types.H)["must"].(types.S)

	// server name match
	if options.Match != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.name": options.Match,
			},
		})
	}

	// ip address
	if options.Address != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.address": options.Address,
			},
		})
	}

	// geo info
	if options.Country != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.country": options.Country,
			},
		})
	}
	if options.CountryCode != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.country_code": options.CountryCode,
			},
		})
	}
	if options.Continent != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.continent": options.Continent,
			},
		})
	}
	if options.Platform != types.PlatformAny {
		condition = append(condition, types.H{
			"match": types.H{
				"data.platform": options.Platform.String(),
			},
		})
	}

	// game options
	if options.Season != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.season": options.Season,
			},
		})
	}
	if options.GameMode != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.game_mode": options.GameMode,
			},
		})
	}
	if options.Intent != "" {
		condition = append(condition, types.H{
			"match": types.H{
				"data.intent": options.Intent,
			},
		})
	}
	if options.Level > 0 {
		condition = append(condition, types.H{
			"range": types.H{
				"data.level": types.H{
					"gte": options.Level,
				},
			},
		})
	}
	if options.PvpEnabled != 0 {
		condition = append(condition, types.H{
			"match": types.H{
				"data.pvp": max(options.PvpEnabled, 0),
			},
		})
	}
	if options.ModEnabled != 0 {
		condition = append(condition, types.H{
			"match": types.H{
				"data.mod": max(options.ModEnabled, 0),
			},
		})
	}
	if options.HasPassword != 0 {
		condition = append(condition, types.H{
			"match": types.H{
				"data.password": max(options.HasPassword, 0),
			},
		})
	}

	// server type
	switch options.ServerType {
	case types.TypeDedicated:
		// dedicated = 1
		condition = append(condition, types.H{
			"match": types.H{
				"data.dedicated": 1,
			},
		})
	case types.TypeClientHosted:
		// client_hosted = 1
		condition = append(condition, types.H{
			"match": types.H{
				"data.client_hosted": 1,
			},
		})
	case types.TypeOfficial:
		// name LIKE "%Public EsServer%" OR ame LIKE "%Klei Official%" OR host IN ("KU_KleiServ", "KU_KleiServ", "KU_zHxWQDSW")
		condition = append(condition, types.H{
			"bool": types.H{
				"should": types.S{
					types.H{
						"match": types.H{
							"data.name": "Public EsServer",
						},
					},
					types.H{
						"match": types.H{
							"data.name": "Klei Official",
						},
					},
					types.H{
						"terms": types.H{
							"data.host": types.S{"KU_KleiServ", "KU_KleiServ", "KU_zHxWQDSW"},
						},
					},
				},
			},
		})
	case types.TypeSteamClan:
		// steam_clan_id != ""
		condition = append(condition, types.H{
			"bool": types.H{
				"must_not": types.S{
					types.H{
						"term": types.H{
							"data.steam_clan_id": "",
						},
					},
				},
			},
		})
	case types.TypeSteamClanOnly:
		// steam_clan_id != "" AND clan_only = 1
		condition = append(condition, types.H{
			"bool": types.H{
				"must_not": types.S{
					types.H{
						"term": types.H{
							"data.steam_clan_id": "",
						},
					},
				},
				"must": types.S{
					types.H{
						"term": types.H{
							"data.clan_only": 1,
						},
					},
				},
			},
		})
	case types.TypeFriendOnly:
		// friend_only = 1
		condition = append(condition, types.H{
			"match": types.H{
				"data.friend_only": 1,
			},
		})
	}

	order := func(desc bool) string {
		if desc {
			return "desc"
		}
		return "asc"
	}

	// sort order
	switch options.Sort {
	case types.DstSortByName:
		dsl["sort"] = types.H{
			"data.name.keyword": types.H{
				"order": order(options.Desc),
			},
		}
	case types.DstSortByCountry:
		dsl["sort"] = types.H{
			"data.country.keyword": types.H{
				"order": order(options.Desc),
			},
		}
	case types.DstSortByVersion:
		dsl["sort"] = types.H{
			"data.version": types.H{
				"order": order(options.Desc),
			},
		}
	case types.DstSortByOnline:
		dsl["sort"] = types.H{
			"data.online": types.H{
				"order": order(options.Desc),
			},
		}
	case types.DstSortByLevel:
		dsl["sort"] = types.H{
			"data.level": types.H{
				"order": order(options.Desc),
			},
		}
	default:
		dsl["sort"] = types.H{
			"data.id": types.H{
				"order": order(options.Desc),
			},
		}
	}

	// apply condition
	dsl["query"].(types.H)["bool"].(types.H)["must"] = condition

	// search from es
	_, body, err := esSearch(s.Elastic, ctx, []string{"lobby-servers"}, dsl)
	if err != nil {
		return nil, 0, err
	}
	// total of result
	totalResult := gjson.GetBytes(body, "hits.total.value").Int()

	// collect data
	var servers []types.EsServer
	hits := gjson.GetBytes(body, "hits.hits").Array()
	buffer := strings.NewReader("")
	for _, result := range hits {
		dataServer := types.EsServer{}
		dataJson := gjson.Get(result.Raw, "_source.data").Raw
		buffer.Reset(dataJson)
		err = json.NewDecoder(buffer).DecodeWithOption(&dataServer)
		if err != nil {
			return nil, 0, err
		}
		servers = append(servers, dataServer)
	}

	// type convert
	var result []*ent.Server
	for _, ess := range servers {
		result = append(result, types.EsServerToEntServer(ess))
	}
	return result, int(totalResult), err
}
