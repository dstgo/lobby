package repo

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/ent/secondary"
	"github.com/dstgo/lobby/server/data/ent/server"
	"github.com/dstgo/lobby/server/data/ent/tag"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
	"github.com/dstgo/lobby/server/types"
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

	// geo search
	if options.Country != "" {
		query = query.Where(server.CountryCodeContainsFold(options.Country))
	}
	if options.CountryCode != "" {
		query = query.Where(server.CountryCode(options.CountryCode))
	}
	if options.Continent != "" {
		query = query.Where(server.ContinentContainsFold(options.Continent))
	}
	if options.Platform != "" {
		query = query.Where(server.Platform(options.Platform))
	}

	// server type
	switch options.ServerType {
	case types.TypeDedicated:
		query = query.Where(server.Dedicated(true))
	case types.TypeClientHosted:
		query = query.Where(server.ClientHosted(true))
	case types.TypeOfficial:
		query = query.Where(server.Or(
			server.NameContains("Public Server"),
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
