package user

import (
	"context"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/types/user"
	"github.com/ginx-contribs/ginx/pkg/resp/statuserr"
)

func NewUserHandler(userRepo *repo.UserRepo) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

type UserHandler struct {
	userRepo *repo.UserRepo
}

func (u *UserHandler) FindByUID(ctx context.Context, uid string) (user.UserInfo, error) {
	record, err := u.userRepo.FindByUID(ctx, uid)
	if err != nil {
		return user.UserInfo{}, statuserr.InternalError(err)
	}
	return user.RecordToUser(record), nil
}

func (u *UserHandler) ListUserByPage(ctx context.Context, page, size int, pattern string) (user.UserListResult, error) {
	users, err := u.userRepo.ListByPage(ctx, page, size, pattern)
	if err != nil {
		return user.UserListResult{}, statuserr.InternalError(err)
	}
	toUsers := user.RecordsToUsers(users)
	return user.UserListResult{Total: int64(len(users)), List: toUsers}, nil
}
