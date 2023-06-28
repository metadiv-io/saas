package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/micro"
)

// Admin only apis are only accessible by admin users.
// User, workspace users and api keys are not allowed to access these apis.

func AdminOnly(engine *micro.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isMicro(ctx, engine) {
			ctx.Next()
			return
		}

		c := micro.NewContext[struct{}](engine, ctx, 0)
		j := c.AuthJwt()
		if j == nil {
			fmt.Println("j is nil")
			c.Err(constant.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		if !j.IsAdmin() || !j.IsIPAllowed(c.ClientIP()) || !j.IsUserAgentAllowed(c.UserAgent()) {
			fmt.Println("j is not admin")
			c.Err(constant.ERR_CODE_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(401, c.Response)
			return
		}

		ctx.Next()
	}
}
