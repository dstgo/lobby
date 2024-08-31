package dst

type ServerType uint

const (
	// TypeAny 任意类型
	TypeAny ServerType = iota
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

// SortType represents the sort way of sorting server list
type SortType uint

const (
	SortByName SortType = iota
	SortByCountry
	SortByVersion
	SortByOnline
	SortByLevel
)
