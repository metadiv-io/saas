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
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}

func CachedGET[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.GET(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func POST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.POST(route, append(middleware, handler.GinHandler(engine))...)
}

func CachedPOST[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.POST(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func PUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func CachedPUT[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.PUT(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func PATCH[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.PATCH(route, append(middleware, handler.GinHandler(engine))...)
}

func DELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func CachedDELETE[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.DELETE(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func OPTIONS[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.OPTIONS(route, append(middleware, handler.GinHandler(engine))...)
}

func HEAD[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.HEAD(route, append(middleware, handler.GinHandler(engine))...)
}

func Any[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.Any(route, append(middleware, handler.GinHandler(engine))...)
}

func WS[T any](engine *micro.Engine, route string, handler micro.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Minute, 60), middleware...)
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}
