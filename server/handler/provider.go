package handler

import (
	"github.com/dstgo/lobby/server/handler/auth"
	"github.com/dstgo/lobby/server/handler/dst"
	"github.com/dstgo/lobby/server/handler/email"
	"github.com/dstgo/lobby/server/handler/user"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	// auth handlers
	auth.NewAuthHandler,
	auth.NewTokenHandler,
	auth.NewVerifyCodeHandler,

	// email handlers
	email.NewEmailHandler,

	// user handlers
	user.NewUserHandler,

	// dst handlers
	dst.NewLobbyHandler,
)
