package mids

import (
	"errors"
	authhandler "github.com/dstgo/lobby/server/handler/auth"
	"github.com/dstgo/lobby/server/types"
	"github.com/dstgo/lobby/server/types/route"
	"github.com/gin-gonic/gin"
	"github.com/ginx-contribs/ginx"
	"github.com/ginx-contribs/ginx/constant/headers"
	"github.com/ginx-contribs/ginx/constant/status"
	"github.com/ginx-contribs/ginx/pkg/resp"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// TokenAuthenticator authenticates each request
func TokenAuthenticator(tokenHandler *authhandler.TokenHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// check if is public api
		metadata := ginx.MetaFromCtx(ctx)
		if !metadata.Contains(route.Private) {
			ctx.Next()
			return
		}

		// parse token string from header
		now := time.Now()
		header := ctx.Request.Header.Get(headers.Authorization)
		pair := strings.Split(header, " ")
		if len(pair) != 2 || pair[0] != "Bearer" {
			resp.Fail(ctx).Status(status.Unauthorized).Error(types.ErrCredentialInvalid).JSON()
			ctx.Abort()
			return
		}
		tokenString := pair[1]

		// verify token if is valid
		tokenInfo, err := tokenHandler.VerifyAccess(ctx, tokenString, now)
		if err == nil {
			// stores token info into context
			types.SetTokenInfo(ctx, &tokenInfo)
			ctx.Next()
		} else {
			// check if is needed to refresh
			if errors.Is(err, types.ErrTokenNeedsRefresh) {
				resp.Fail(ctx).Error(types.ErrTokenNeedsRefresh).JSON()
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				resp.Fail(ctx).Error(types.ErrCredentialExpired).JSON()
			} else { // invalid token
				ctx.Error(err)
				resp.Fail(ctx).Error(types.ErrCredentialInvalid).JSON()
			}
		}
	}
}
