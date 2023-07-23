package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/saas/constant"
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
			c.Err(constant.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if (!j.IsWorkspaceUser() && !j.IsAPIKey()) || !j.IsIPAllowed(c.ClientIP()) || !j.IsUserAgentAllowed(c.UserAgent()) {
			c.Err(constant.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if c.Workspace() == "" {
			c.Err(constant.ERR_CODE_WORKSPACE_NOT_FOUND)
			ctx.AbortWithStatusJSON(404, c.Response)
			return
		}

		if !j.IsWorkspaceAllowed(c.Workspace()) {
			c.Err(constant.ERR_CODE_FORBIDDEN)
			ctx.AbortWithStatusJSON(403, c.Response)
			return
		}

		if !micro.UsageManager.AskWorkspaceAllowed(c.Workspace(), j.UserUUID, micro.UsageManager.TagToApi[c.ApiTag()].UUID) {
			c.Err(constant.ERR_CODE_NOT_ENOUGH_CREDIT)
			ctx.AbortWithStatusJSON(403, c.Response)
			return
		}

		ctx.Next()
	}
}
