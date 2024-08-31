package types

import (
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/pkg/lobbyapi"
)

type DstServerType uint

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
type DstSortType uint

const (
	DstSortByName DstSortType = iota
	DstSortByCountry
	DstSortByVersion
	DstSortByOnline
	DstSortByLevel
)

type LobbyServerSearchOptions struct {
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
	// platform name, Steam | WeGame | PSN | XBone | PS4Official | Rail | Switch
	Platform string `form:"platform"`
	// 0-any 1-dedicated 2-client hosted 3-official 4-steam group 5-steam group only 6-friend only
	ServerType DstServerType `form:"serverType" binding:"gte=0,lte=6"`

	// season spring | summer | autumn | winter or other mods season
	Season string `form:"season"`
	// tags of server
	Tags []string `form:"tags"`
	// mode for dst
	GameMode string `form:"gameMode"`
	// intent for dst
	Intent string `form:"intent"`
	// world level of server
	Level int `form:"level"`

	// -1-off 0-all 1-on
	PvpEnabled int `form:"pvp"`
	// -1-off 0-all 1-on
	ModEnabled int `form:"mod"`
	// -1-off 0-all 1-on
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
