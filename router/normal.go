package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/utils"
)

// Normal apis are accessible by all users.

func GET[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.GET(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func POST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.POST(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func PUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.PUT(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func PATCH[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.PATCH(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func DELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.DELETE(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func OPTIONS[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.OPTIONS(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func HEAD[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.HEAD(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func Any[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.Any(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WS[T any](engine *micro.Engine, route string, handler micro.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.GET(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}
