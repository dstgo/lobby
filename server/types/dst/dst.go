package dst

import "github.com/dstgo/lobby/pkg/lobbyapi"

type SearchOptions struct {
	// n page to search
	Page int `form:"page"`
	// size of page
	Size int `form:"size"`
	// search content
	Match string `form:"text"`
	// 0-name 1-country 2-version 3-online 4-level
	Sort SortType `form:"sort" binding:"gte=0,lte=4"`
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
	ServerType ServerType `form:"server_type" binding:"gte=0,lte=6"`

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

type QueryListResult struct {
	Total int          `json:"total"`
	List  []ServerInfo `json:"list"`
}

type ServerInfo struct {
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

type QueryDetailsOptions struct {
	Region string `form:"region" binding:"required"`
	RowID  string `form:"rowId" binding:"required"`
}

type QueryDetailsResult struct {
	ServerInfo
	lobbyapi.MetaInfo
}
