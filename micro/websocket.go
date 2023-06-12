package micro

import (
	"net/http"

	"github.com/metadiv-io/saas/types"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsService[T any] func(ctx *Context[T]) *types.Error

type WsHandler[T any] func() WsService[T]

func (h WsHandler[T]) GinHandler(engine *Engine) gin.HandlerFunc {
	wsUpGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return func(ctx *gin.Context) {
		ws, err := wsUpGrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer ws.Close()

		service := h()
		c := NewContext[T](engine, ctx, h.QueryApiCredit(ctx))
		err1 := service(c)
		if err1 != nil {
			ctx.JSON(500, gin.H{
				"code":    err1.Code,
				"message": err1.Message,
			})
		}
	}
}

func (h WsHandler[T]) QueryApiCredit(ctx *gin.Context) float64 {
	return UsageManager.GetByTag(ctx.Request.Method + ":" + ctx.FullPath()).Credit
}
