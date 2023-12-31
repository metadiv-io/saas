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

// Workspace apis are accessible by all workspace users and api keys.

func WorkspaceGET[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("GET", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().GET(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceCachedGET[T any](engine *micro.Engine, route, uuid string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("GET", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().GET(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func WorkspacePOST[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("POST", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().POST(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceCachedPOST[T any](engine *micro.Engine, route, uuid string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("POST", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().POST(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func WorkspacePUT[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("PUT", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceCachedPUT[T any](engine *micro.Engine, route, uuid string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("PUT", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().PUT(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func WorkspacePATCH[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("PATCH", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().PATCH(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceDELETE[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("DELETE", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceCachedDELETE[T any](engine *micro.Engine, route, uuid string, duration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("DELETE", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().DELETE(route, append(middleware, ginmid.Cache(duration, handler.GinHandler(engine)))...)
}

func WorkspaceOPTIONS[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("OPTIONS", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().OPTIONS(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceHEAD[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("HEAD", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().HEAD(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceAny[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("Any", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().Any(route, append(middleware, handler.GinHandler(engine))...)
}

func WorkspaceWS[T any](engine *micro.Engine, route, uuid string, handler ginger.WsHandler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("WS", route, uuid)
	middleware = util.JoinHandlerAtStart(mid.WorkspaceUserOnly(engine), middleware...)
	middleware = util.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin().GET(route, append(middleware, handler.GinHandler(engine))...)
}
