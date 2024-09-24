package types

import (
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
)

type DstServerType int

const (
	// TypeAny 任意类型
	TypeAny DstServerType = iota
	// TypeDedicated 专用服务器
	TypeDedicated
	// TypeClientHosted 客户端自建
	TypeClientHosted
	// TypeOfficial 官方服务器
	TypeOfficial
	// TypeSteamClan 群组服务器
	TypeSteamClan
	// TypeSteamClanOnly 仅群组服
	TypeSteamClanOnly
	// TypeFriendOnly 仅好友
	TypeFriendOnly
)

// DstSortType represents the sort way of sorting server list
type DstSortType int

const (
	DstSortByName DstSortType = iota
	DstSortByCountry
	DstSortByVersion
	DstSortByOnline
	DstSortByLevel
)

type ServerPlatform int

const (
	PlatformAny ServerPlatform = iota
	PlatformSteam
	// PlatformWeGame just a helper name, actually it is not exist
	PlatformWeGame
	PlatformPSN
	// PlatformPS4Official can not be use in api query params
	PlatformPS4Official
	PlatformXBOne
	PlatformSwitch
	// PlatformRail is alias of PlatformWeGame, only serve at ap-east-1
	PlatformRail
)

func (p ServerPlatform) String() string {
	switch p {
	case PlatformAny:
		return ""
	case PlatformSteam:
		return "Steam"
	case PlatformPSN:
		return "Rail"
	case PlatformXBOne:
		return "XBone"
	case PlatformPS4Official:
		return "PS4Official"
	case PlatformSwitch:
		return "Switch"
	case PlatformWeGame:
		return "WeGame"
	default:
		panic("unknown server platform")
	}
}

type LobbyServerSearchOptions struct {
	Qv int64 `form:"_qv"`
	// n page to search
	Page int `form:"page"`
	// size of page
	Size int `form:"size"`
	// search content
	Match string `form:"text"`
	// 0-name 1-country 2-version 3-online 4-level
	Sort DstSortType `form:"sort" binding:"gte=0,lte=4"`
	// descending order
	Desc bool `form:"desc"`

	// ip address
	Address string `form:"address"`
	// country name
	Country string `form:"country"`
	// iso country code, eg. CN
	CountryCode string `form:"countryCode"`
	// continent name
	Continent string `form:"continent"`
	// server platform, 0-Any, 1-Steam, 2-PlatformWeGame, 3-PSN, 4-PS4Official, 5-XBone, 6-Switch
	Platform ServerPlatform `form:"platform" binding:"gte=0,lte=7"`
	// 0-any 1-dedicated 2-client, hosted, 3-official, 4-steam group, 5-steam group only, 6-friend only
	ServerType DstServerType `form:"serverType" binding:"gte=0,lte=6"`

	// season spring | summer | autumn | winter | other mods season
	Season string `form:"season"`
	// tags of server
	Tags []string `form:"tags"`
	// mode for dst
	GameMode string `form:"gameMode"`
	// intent for dst
	Intent string `form:"intent"`
	// world level of server
	Level int `form:"level"`

	// -1-off, 0-any, 1-on
	PvpEnabled int `form:"pvp"`
	// -1-off, 0-any, 1-on
	ModEnabled int `form:"mod"`
	// -1-off, 0-any, 1-on
	HasPassword int `form:"password"`
}

type LobbyServerSearchResult struct {
	Total int               `json:"total"`
	List  []LobbyServerInfo `json:"list"`
}

type LobbyServerInfo struct {
	// network info
	GUID        string `json:"guid"`
	RowID       string `json:"rowID"`
	SteamID     string `json:"steamID"`
	SteamClanID string `json:"steamClanID"`
	OwnerID     string `json:"ownerID"`
	SteamRoom   string `json:"steamRoom"`
	Session     string `json:"session"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
	Host        string `json:"host"`
	Platform    string `json:"platform"`
	ClanOnly    bool   `json:"clanOnly"`
	LanOnly     bool   `json:"lanOnly"`
	// game info
	Name      string `json:"name"`
	GameMode  string `json:"gameMode"`
	Intent    string `json:"intent"`
	Season    string `json:"season"`
	Version   int    `json:"version"`
	MaxOnline int    `json:"maxOnline"`
	Online    int    `json:"online"`
	Level     int    `json:"level"`

	// switch options
	Mod             bool `json:"mod"`
	Pvp             bool `json:"pvp"`
	Password        bool `json:"password"`
	Dedicated       bool `json:"dedicated"`
	ClientHosted    bool `json:"clientHosted"`
	AllowNewPlayers bool `json:"allowNewPlayers"`
	ServerPaused    bool `json:"serverPaused"`
	FriendOnly      bool `json:"friendOnly"`
	// geo info
	Country     string `json:"country"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Continent   string `json:"continent"`
	Region      string `json:"region"`

	Tags        []string                        `json:"tags,omitempty"`
	Secondaries map[string]lobbyapi.Secondaries `json:"secondaries,omitempty"`
}

type LobbyServerDetailsOptions struct {
	Region string `form:"region" binding:"required"`
	RowID  string `form:"rowId" binding:"required"`
}

type LobbyServerDetails struct {
	LobbyServerInfo
	lobbyapi.MetaInfo
}

// EntServerToServerInfo converts *ent.Server => LobbyServerInfo
func EntServerToServerInfo(server *ent.Server) LobbyServerInfo {
	return LobbyServerInfo{
		GUID:            server.GUID,
		RowID:           server.RowID,
		SteamID:         server.SteamID,
		SteamClanID:     server.SteamClanID,
		OwnerID:         server.OwnerID,
		SteamRoom:       server.SteamRoom,
		Session:         server.Session,
		Address:         server.Address,
		Port:            server.Port,
		Host:            server.Host,
		Platform:        server.Platform,
		ClanOnly:        server.ClanOnly,
		LanOnly:         server.LanOnly,
		Name:            server.Name,
		GameMode:        server.GameMode,
		Intent:          server.Intent,
		Season:          server.Season,
		Version:         server.Version,
		MaxOnline:       server.MaxOnline,
		Online:          server.Online,
		Mod:             server.Mod,
		Pvp:             server.Pvp,
		Password:        server.Password,
		Dedicated:       server.Dedicated,
		ClientHosted:    server.ClientHosted,
		AllowNewPlayers: server.AllowNewPlayers,
		ServerPaused:    server.ServerPaused,
		FriendOnly:      server.FriendOnly,
		Country:         server.Country,
		CountryCode:     server.CountryCode,
		Continent:       server.Continent,
		City:            server.City,
		Region:          server.Region,
		Level:           server.Level,
	}
}

// LobbyServerToServerInfo converts lobbyapi.Server => lobbyapi.Server
func LobbyServerToServerInfo(server lobbyapi.Server) LobbyServerInfo {
	return LobbyServerInfo{
		GUID:            server.Guid,
		RowID:           server.RowId,
		SteamID:         server.SteamId,
		SteamClanID:     server.SteamClanId,
		OwnerID:         server.OwnerNetId,
		SteamRoom:       server.SteamRoom,
		Session:         server.Session,
		Address:         server.Address,
		Port:            server.Port,
		Host:            server.Host,
		Platform:        lobbyapi.PlatformDisplayName(server.Region, server.Platform),
		ClanOnly:        server.ClanOnly,
		LanOnly:         server.LanOnly,
		Name:            server.Name,
		GameMode:        server.GameMode,
		Intent:          server.Intent,
		Season:          server.Season,
		Version:         server.Version,
		MaxOnline:       server.MaxConnections,
		Online:          server.Connected,
		Mod:             server.Mod,
		Pvp:             server.Pvp,
		Password:        server.HasPassword,
		Dedicated:       server.IsDedicated,
		ClientHosted:    server.ClientHosted,
		AllowNewPlayers: server.AllowNewPlayers,
		ServerPaused:    server.ServerPaused,
		FriendOnly:      server.FriendOnly,
		Region:          server.Region,
		Tags:            server.Tags,
		Secondaries:     server.Secondaries,
	}
}

// LobbyServerToEntServer lobbyapi.Server => *ent.Server
func LobbyServerToEntServer(server lobbyapi.Server) *ent.Server {
	return &ent.Server{
		GUID:            server.Guid,
		RowID:           server.RowId,
		SteamID:         server.SteamId,
		SteamClanID:     server.SteamClanId,
		OwnerID:         server.OwnerNetId,
		SteamRoom:       server.SteamRoom,
		Session:         server.Session,
		Address:         server.Address,
		Port:            server.Port,
		Host:            server.Host,
		Platform:        lobbyapi.PlatformDisplayName(server.Region, server.Platform),
		ClanOnly:        server.ClanOnly,
		LanOnly:         server.LanOnly,
		Name:            server.Name,
		GameMode:        server.GameMode,
		Intent:          server.Intent,
		Season:          server.Season,
		Version:         server.Version,
		MaxOnline:       server.MaxConnections,
		Online:          server.Connected,
		Mod:             server.Mod,
		Pvp:             server.Pvp,
		Password:        server.HasPassword,
		Dedicated:       server.IsDedicated,
		ClientHosted:    server.ClientHosted,
		AllowNewPlayers: server.AllowNewPlayers,
		ServerPaused:    server.ServerPaused,
		FriendOnly:      server.FriendOnly,
		Region:          server.Region,
	}
}

func EsServerToEntServer(ess EsServer) *ent.Server {
	return &ent.Server{
		ID:              ess.ID,
		GUID:            ess.GUID,
		RowID:           ess.RowID,
		SteamID:         ess.SteamID,
		SteamClanID:     ess.SteamClanID,
		OwnerID:         ess.OwnerID,
		SteamRoom:       ess.SteamRoom,
		Session:         ess.Session,
		Address:         ess.Address,
		Port:            ess.Port,
		Host:            ess.Host,
		Platform:        ess.Platform,
		ClanOnly:        ess.ClanOnly > 0,
		LanOnly:         ess.LanOnly > 0,
		Name:            ess.Name,
		GameMode:        ess.GameMode,
		Intent:          ess.Intent,
		Season:          ess.Season,
		Version:         ess.Version,
		MaxOnline:       ess.MaxOnline,
		Online:          ess.Online,
		Level:           ess.Level,
		Mod:             ess.Mod > 0,
		Pvp:             ess.Pvp > 0,
		Password:        ess.Password > 0,
		Dedicated:       ess.Dedicated > 0,
		ClientHosted:    ess.ClientHosted > 0,
		AllowNewPlayers: ess.AllowNewPlayers > 0,
		ServerPaused:    ess.ServerPaused > 0,
		FriendOnly:      ess.FriendOnly > 0,
		QueryVersion:    ess.QueryVersion,
		Country:         ess.Country,
		Continent:       ess.Continent,
		CountryCode:     ess.CountryCode,
		City:            ess.City,
		Region:          ess.Region,
	}
}

type EsServer struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GUID holds the value of the "guid" field.
	GUID string `json:"guid,omitempty"`
	// RowID holds the value of the "row_id" field.
	RowID string `json:"row_id,omitempty"`
	// SteamID holds the value of the "steam_id" field.
	SteamID string `json:"steam_id,omitempty"`
	// SteamClanID holds the value of the "steam_clan_id" field.
	SteamClanID string `json:"steam_clan_id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// SteamRoom holds the value of the "steam_room" field.
	SteamRoom string `json:"steam_room,omitempty"`
	// Session holds the value of the "session" field.
	Session string `json:"session,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Port holds the value of the "port" field.
	Port int `json:"port,omitempty"`
	// Host holds the value of the "host" field.
	Host string `json:"host,omitempty"`
	// Platform holds the value of the "platform" field.
	Platform string `json:"platform,omitempty"`
	// ClanOnly holds the value of the "clan_only" field.
	ClanOnly int `json:"clan_only,omitempty"`
	// LanOnly holds the value of the "lan_only" field.
	LanOnly int `json:"lan_only,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GameMode holds the value of the "game_mode" field.
	GameMode string `json:"game_mode,omitempty"`
	// Intent holds the value of the "intent" field.
	Intent string `json:"intent,omitempty"`
	// Season holds the value of the "season" field.
	Season string `json:"season,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// MaxOnline holds the value of the "max_online" field.
	MaxOnline int `json:"max_online,omitempty"`
	// Online holds the value of the "online" field.
	Online int `json:"online,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// Mod holds the value of the "mod" field.
	Mod int `json:"mod,omitempty"`
	// Pvp holds the value of the "pvp" field.
	Pvp int `json:"pvp,omitempty"`
	// Password holds the value of the "password" field.
	Password int `json:"password,omitempty"`
	// Dedicated holds the value of the "dedicated" field.
	Dedicated int `json:"dedicated,omitempty"`
	// ClientHosted holds the value of the "client_hosted" field.
	ClientHosted int `json:"client_hosted,omitempty"`
	// AllowNewPlayers holds the value of the "allow_new_players" field.
	AllowNewPlayers int `json:"allow_new_players,omitempty"`
	// ServerPaused holds the value of the "server_paused" field.
	ServerPaused int `json:"server_paused,omitempty"`
	// FriendOnly holds the value of the "friend_only" field.
	FriendOnly int `json:"friend_only,omitempty"`
	// QueryVersion holds the value of the "query_version" field.
	QueryVersion int64 `json:"query_version,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Continent holds the value of the "continent" field.
	Continent string `json:"continent,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
}
