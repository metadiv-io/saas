package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/micro"
)

// User apis are only accessible by users.
// Admin users, workspace users and api keys are not allowed to access these apis.

func UserOnly(engine *micro.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if env.String("GIN_MODE") == gin.DebugMode {
			ctx.Next()
			return
		}

		if isMicro(ctx, engine) {
			ctx.Next()
			return
		}

		c := micro.NewContext[struct{}](engine, ctx, 0)
		j := c.AuthJwt()
		if j == nil {
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsUser() || !j.IsIPAllowed(c.ClientIP()) || !j.IsUserAgentAllowed(c.UserAgent()) {
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		ctx.Next()
	}
}
