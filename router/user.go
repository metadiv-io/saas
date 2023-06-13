package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/saas/micro"
	mid "github.com/metadiv-io/saas/middleware"
	"github.com/metadiv-io/saas/utils"
)

// User apis are accessible by all users.

func UserGET[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}

func UserPOST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.POST(route, append(middleware, handler.GinHandler(engine))...)
}

func UserPUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func UserPATCH[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.PATCH(route, append(middleware, handler.GinHandler(engine))...)
}

func UserDELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func UserOPTIONS[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.OPTIONS(route, append(middleware, handler.GinHandler(engine))...)
}

func UserHEAD[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.HEAD(route, append(middleware, handler.GinHandler(engine))...)
}

func UserAny[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.Any(route, append(middleware, handler.GinHandler(engine))...)
}

func UserWs[T any](engine *micro.Engine, route string, handler micro.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = utils.JoinHandlerAtStart(mid.UserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}
