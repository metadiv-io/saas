package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/micro"
)

// Admin only apis are only accessible by admin users.
// User, workspace users and api keys are not allowed to access these apis.

func AdminOnly(engine *micro.Engine) gin.HandlerFunc {
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
			log.Println("jwt is nil")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsAdmin() {
			log.Println("user is not admin role")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsIPAllowed(c.ClientIP()) {
			log.Panicln("ip is not allowed")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsUserAgentAllowed(c.UserAgent()) {
			log.Panicln("user agent is not allowed")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		ctx.Next()
	}
}
