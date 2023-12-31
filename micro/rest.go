package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger"
)

type Service[T any] func(ctx IContext[T])

type Handler[T any] func() Service[T]

func (h Handler[T]) GinHandler(engine IEngine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service := h()
		c := NewContext[T](engine, ctx, h.QueryApiCredit(ctx))
		service(c)

		// if file is served, no need to respond
		if c.IsResponded() && c.IsFile() {
			return
		}

		// unexpected error
		if !c.IsResponded() || c.Response() == nil {
			ctx.JSON(500, gin.H{
				"message": "Service did not define response",
			})
			return
		}

		if c.Response().Success {
			ctx.JSON(200, c.Response)
			return
		}

		// unexpected error
		if c.Response().Error == nil {
			ctx.JSON(500, gin.H{
				"message": "Service did not define error",
			})
			return
		}

		switch c.Response().Error.Code {
		case ginger.ERR_CODE_UNAUTHORIZED:
			ctx.JSON(401, c.Response)
		case ginger.ERR_CODE_FORBIDDEN, ERR_CODE_NOT_ENOUGH_CREDIT:
			ctx.JSON(403, c.Response)
		case ERR_CODE_WORKSPACE_NOT_FOUND:
			ctx.JSON(404, c.Response)
		case ginger.ERR_CODE_INTERNAL_SERVER_ERROR:
			ctx.JSON(500, c.Response)
		default:
			ctx.JSON(200, c.Response)
		}
	}
}

func (h Handler[T]) QueryApiCredit(ctx *gin.Context) float64 {
	api := UsageManager.GetByTag(ctx.Request.Method + ":" + ctx.FullPath())
	if api == nil {
		return 0
	}
	return api.Credit
}
