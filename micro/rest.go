package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/saas/constant"
)

type Service[T any] func(ctx *Context[T])

type Handler[T any] func() Service[T]

func (h Handler[T]) GinHandler(engine *Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service := h()
		c := NewContext[T](engine, ctx, h.QueryApiCredit(ctx))
		service(c)

		// unexpected error
		if !c.IsResponded || c.Response != nil {
			ctx.JSON(500, gin.H{
				"message": "Service did not define response",
			})
		}

		if c.Response.Success {
			ctx.JSON(200, c.Response)
			return
		}

		// unexpected error
		if c.Response.Error == nil {
			ctx.JSON(500, gin.H{
				"message": "Service did not define error",
			})
			return
		}

		switch c.Response.Error.Code {
		case constant.ERR_CODE_UNAUTHORIZED:
			ctx.JSON(401, c.Response)
		case constant.ERR_CODE_FORBIDDEN, constant.ERR_CODE_NOT_ENOUGH_CREDIT:
			ctx.JSON(403, c.Response)
		case constant.ERR_CODE_WORKSPACE_NOT_FOUND:
			ctx.JSON(404, c.Response)
		case constant.ERR_CODE_INTERNAL_SERVER_ERROR:
			ctx.JSON(500, c.Response)
		default:
			ctx.JSON(200, c.Response)
		}
	}
}

func (h Handler[T]) QueryApiCredit(ctx *gin.Context) float64 {
	return UsageManager.GetByTag(ctx.Request.Method + ":" + ctx.FullPath()).Credit
}
