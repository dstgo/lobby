package user

type PageOption struct {
	Page   int    `form:"page" binding:"required,gt=0"`
	Size   int    `form:"size" binding:"required,gt=0"`
	Search string `form:"search"`
}

type UidOption struct {
	Uid string `form:"uid" binding:"required"`
}

type UserInfo struct {
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

type UserListResult struct {
	Total int64      `json:"total"`
	List  []UserInfo `json:"list"`
}
