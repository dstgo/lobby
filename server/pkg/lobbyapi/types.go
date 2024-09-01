package lobbyapi

const GameId = "DontStarveTogether"

// OriginalPlatforms returns the original platforms name got from klei api
var OriginalPlatforms = []string{Steam.String(), PSN.String(), Rail.String(), XBOne.String(), Switch.String()}

const (
	Any Platform = iota
	Steam
	// WeGame just a helper name, actually it is not exist
	WeGame
	PSN
	// PS4Official can not be use in api query params
	PS4Official
	XBOne
	Switch
	// Rail is alias of WeGame, only serve at ap-east-1
	Rail
)

// Platform represents dst server platform, it may be updated by klei in the future
type Platform int

func (p Platform) String() string {
	switch p {
	case Any:
		return "Any"
	case Steam:
		return "Steam"
	case PSN:
		return "PSN"
	case Rail:
		return "Rail"
	case XBOne:
		return "XBone"
	case PS4Official:
		return "PS4Official"
	case Switch:
		return "Switch"
	case WeGame:
		return "WeGame"
	default:
		panic("unhandled default case")
	}
}

// Region represents dst lobby server region, it may be updated by klei in the future
const (
	UsEast1     = "us-east-1"
	EuCentral   = "eu-central-1"
	ApSoutheast = "ap-southeast-1"
	ApEast      = "ap-east-1"
)

// Lobby server urls, may be updated by klei in the future
const (
	LobbyRegionURL  = `https://lobby-v2-cdn.klei.com/regioncapabilities-v2.json`
	LobbyServersURL = `https://lobby-v2-cdn.klei.com/{{.region}}-{{.platform}}.json.gz`
	LobbyDetailsURL = "https://lobby-v2-{{.region}}.klei.com/lobby/read"
)

type Regions struct {
	Regions []struct {
		Region string `json:"Region"`
	} `json:"LobbyRegions"`
}

// Server includes all the information about single dst server
type Server struct {
	// network options
	Guid  string `json:"guid"`
	RowId string `json:"__rowId"`
	// only at steam platform
	SteamId string `json:"steamid"`
	// only for clan server
	SteamClanId string `json:"steamclanid"`
	// only for no password server
	OwnerNetId string   `json:"ownernetid"`
	SteamRoom  string   `json:"steamroom"`
	Session    string   `json:"session"`
	Address    string   `json:"__addr"`
	Port       int      `json:"port"`
	Host       string   `json:"host"`
	Platform   Platform `json:"platform"`

	ClanOnly bool `json:"clanonly"`
	LanOnly  bool `json:"lanonly"`

	// second shard
	Secondaries map[string]Secondaries

	// game options
	Name     string `json:"name"`
	GameMode string `json:"mode"`
	Intent   string `json:"intent"`
	Season   string `json:"season"`
	TagStr   string `json:"tags"`
	Version  int    `json:"v"`
	// max players allowed
	MaxConnections int `json:"maxconnections"`
	// online players number
	Connected int `json:"connected"`

	Mod             bool `json:"mods"`
	Pvp             bool `json:"pvp"`
	HasPassword     bool `json:"password"`
	IsDedicated     bool `json:"dedicated"`
	ClientHosted    bool `json:"clienthosted"`
	AllowNewPlayers bool `json:"allownewplayers"`
	ServerPaused    bool `json:"serverpaused"`
	FriendOnly      bool `json:"fo"`

	// processed info
	Region string   `json:"region"`
	Tags   []string `json:"tag_arr"`
}

// Secondaries represents the secondaries shard among dst servers
type Secondaries struct {
	Id      string `json:"id"`
	SteamId string `json:"steamid"`
	Address string `json:"__addr"`
	Port    int    `json:"port"`
}

type Servers struct {
	List []Server `json:"GET"`
}

type Player struct {
	Name    string `json:"name"`
	Prefab  string `json:"prefab"`
	SteamId string `json:"steamId"`
	// hex color code
	Colour string `json:"colour"`
	// shard level
	Level int `json:"level"`
}

type Mod struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Version1 string `json:"version1"`
	Version2 string `json:"version2"`
	Enabled  bool   `json:"enabled"`
}

// ServerDetails includes some details information
type ServerDetails struct {
	// repeat options
	Server

	Tick          int  `json:"tick"`
	ClientModsOff bool `json:"clientmodsoff"`
	Nat           int  `json:"nat"`

	// raw lua script data
	Data          string `json:"data"`
	WorldGen      string `json:"worldgen"`
	OnlinePlayers string `json:"players"`
	Mods          []any  `json:"mods_info"`

	// parsed lua data
	Details MetaInfo `json:"details"`
}

type MetaInfo struct {
	Day                int      `json:"day"`
	DayElapsedInSeason int      `json:"dayElapsedInSeason"`
	DaysLeftInSeason   int      `json:"daysLeftInSeason"`
	Players            []Player `json:"players"`
	ModsInfo           []Mod    `json:"mods"`
}
