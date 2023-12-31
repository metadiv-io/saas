package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/micro"
)

// User apis are only accessible by users.
// Admin users, workspace users and api keys are not allowed to access these apis.

func UserOnly(engine micro.IEngine) gin.HandlerFunc {
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

		if !j.IsUser() {
			log.Println("user is not user role")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsIPAllowed(c.ClientIP()) {
			log.Println("ip is not allowed")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsUserAgentAllowed(c.UserAgent()) {
			log.Println("user agent is not allowed")
			c.Err(ginger.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		ctx.Next()
	}
}
