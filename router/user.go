package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/saas/internal/util"
	"github.com/metadiv-io/saas/micro"
	mid "github.com/metadiv-io/saas/middleware"
)

// User apis are accessible by all users.

func UserGET[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().GET(route, append(middleware, handler.GinHandler(engine))...)
}

func UserCachedGET[T any](engine micro.IEngine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().GET(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func UserPOST[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().POST(route, append(middleware, handler.GinHandler(engine))...)
}

func UserCachedPOST[T any](engine micro.IEngine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().POST(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func UserPUT[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func UserCachedPUT[T any](engine micro.IEngine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().PUT(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func UserPATCH[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().PATCH(route, append(middleware, handler.GinHandler(engine))...)
}

func UserDELETE[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func UserCachedDELETE[T any](engine micro.IEngine, route string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	engine.Gin().DELETE(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func UserOPTIONS[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().OPTIONS(route, append(middleware, handler.GinHandler(engine))...)
}

func UserHEAD[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().HEAD(route, append(middleware, handler.GinHandler(engine))...)
}

func UserAny[T any](engine micro.IEngine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().Any(route, append(middleware, handler.GinHandler(engine))...)
}

func UserWs[T any](engine micro.IEngine, route string, handler ginger.WsHandler[T], middleware ...gin.HandlerFunc) {
	middleware = util.JoinHandlerAtStart(mid.UserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*2), middleware...)
	engine.Gin().GET(route, append(middleware, handler.GinHandler(engine))...)
}
