package router

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/saas/micro"
	mid "github.com/metadiv-io/saas/middleware"
	"github.com/metadiv-io/saas/utils"
)

// Admin apis are accessible by admin users only.

func AdminGET[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminPOST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.POST(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminPUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminPATCH[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.PATCH(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminDELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminOPTIONS[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.OPTIONS(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminHEAD[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.HEAD(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminAny[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.Any(route, append(middleware, handler.GinHandler(engine))...)
}

func AdminWS[T any](engine *micro.Engine, route string, handler micro.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}
