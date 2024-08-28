package handler

import (
	"github.com/dstgo/lobby/server/handler/auth"
	"github.com/dstgo/lobby/server/handler/email"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	// auth handlers
	auth.NewAuthHandler,
	auth.NewTokenHandler,
	auth.NewVerifyCodeHandler,

	// email handlers
	email.NewEmailHandler,
)
