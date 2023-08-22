package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/micro"
)

// Workspace user only apis are only accessible by workspace users and api keys.
// Admin users and users are not allowed to access these apis.

func WorkspaceUserOnly(engine *micro.Engine) gin.HandlerFunc {
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

		if !j.IsWorkspaceUser() && !j.IsAPIKey() {
			log.Println("user is not workspace user or api key")
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

		if c.Workspace() == "" {
			log.Println("workspace is empty")
			c.Err(micro.ERR_CODE_WORKSPACE_NOT_FOUND)
			ctx.AbortWithStatusJSON(404, c.Response)
			return
		}

		if !micro.UsageManager.AskWorkspaceAllowed(c.Workspace(), j.UserUUID, micro.UsageManager.TagToApi[c.ApiTag()].UUID) {
			log.Println("workspace is not allowed (usage)")
			c.Err(micro.ERR_CODE_NOT_ENOUGH_CREDIT)
			ctx.AbortWithStatusJSON(403, c.Response)
			return
		}

		ctx.Next()
	}
}
