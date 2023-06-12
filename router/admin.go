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
	engine.Gin.GET(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminPOST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.POST(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminPUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.PUT(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminPATCH[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.PATCH(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminDELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.DELETE(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminOPTIONS[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.OPTIONS(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminHEAD[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.HEAD(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminAny[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.Any(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func AdminWS[T any](engine *micro.Engine, route string, handler micro.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.AdminOnly[T](engine), middleware...)
	engine.Gin.GET(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}
