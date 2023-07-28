package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/saas/micro"
	mid "github.com/metadiv-io/saas/middleware"
	"github.com/metadiv-io/saas/utils"
)

// Admin apis are accessible by admin users only.

func AdminGET[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminCachedGET[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.GET(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func AdminPOST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.POST(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminCachedPOST[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.POST(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func AdminPUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminCachedPUT[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.PUT(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func AdminPATCH[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.PATCH(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminDELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminCachedDELETE[T any](engine *micro.Engine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.DELETE(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func AdminOPTIONS[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.OPTIONS(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminHEAD[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.HEAD(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminAny[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.Any(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminWS[T any](engine *micro.Engine, route string, handler ginger.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly(engine), middleware...)
	engine.GingerEngine.Gin.GET(route, append(middleware, handler.GinHandler(engine.GingerEngine))...)
}
