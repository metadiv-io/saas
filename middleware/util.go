package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/saas/micro"
)

func isMicro(ctx *gin.Context, engine *micro.Engine) bool {
	ip := ctx.ClientIP()
	for _, v := range engine.MicroIPs {
		if v == ip {
			return true
		}
	}
	return false
}
