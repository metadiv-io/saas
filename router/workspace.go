package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/saas/micro"
	mid "github.com/metadiv-io/saas/middleware"
	"github.com/metadiv-io/saas/utils"
)

// Workspace apis are accessible by all workspace users and api keys.

func WorkspaceGET[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("GET", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.GET(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspacePOST[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("POST", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.POST(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspacePUT[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("PUT", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.PUT(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspacePATCH[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("PATCH", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.PATCH(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspaceDELETE[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("DELETE", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.DELETE(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspaceOPTIONS[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("OPTIONS", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.OPTIONS(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspaceHEAD[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("HEAD", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.HEAD(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspaceAny[T any](engine *micro.Engine, route, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("Any", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.Any(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}

func WorkspaceWS[T any](engine *micro.Engine, route, uuid string, handler micro.WsHandler[T], middleware ...gin.HandlerFunc) {
	micro.UsageManager.Register("WS", route, uuid)
	middleware = utils.JoinHandlerAtStart(mid.WorkspaceUserOnly[T](engine), middleware...)
	middleware = utils.JoinHandlerAtStart(ginmid.RateLimited(time.Hour, 60*60*10), middleware...)
	engine.Gin.GET(route, utils.JoinHandlerAtStart(handler.GinHandler(engine), middleware...)...)
}
