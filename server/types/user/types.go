package user

import "github.com/dstgo/lobby/server/data/ent"

func RecordToUser(user *ent.User) UserInfo {
	if user == nil {
		return UserInfo{}
	}

	return UserInfo{
		Uid:       user.UID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func RecordsToUsers(users []*ent.User) []UserInfo {
	if users == nil {
		return []UserInfo{}
	}
	var us []UserInfo
	for _, u := range users {
		us = append(us, RecordToUser(u))
	}
	return us
}
