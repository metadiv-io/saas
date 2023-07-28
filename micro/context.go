package micro

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger"
	gingerTypes "github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/types"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	ginger.Context[T]
	Engine   *Engine
	Response *types.Response
	Credit   float64
}

func NewContext[T any](engine *Engine, ginCtx *gin.Context, credit float64) *Context[T] {
	gingerCtx := ginger.NewContext[T](engine.GingerEngine, ginCtx)
	ctx := &Context[T]{
		Context: *gingerCtx,
		Engine:  engine,
		Credit:  credit,
	}
	return ctx
}

func (ctx *Context[T]) AuthJwt() *types.Jwt {
	j := &types.Jwt{}
	err := j.ParseToken(ctx.Engine.PubPEM, ctx.BearerToken())
	if err != nil {
		logger.Prefix(ctx.LogPrefix()).Error(err)
		return nil
	}
	if j.UUID == "" {
		return nil
	}
	return j
}

func (ctx *Context[T]) Workspace() string {
	j := ctx.AuthJwt()
	if j == nil {
		return ""
	}
	if j.Workspaces != nil && len(j.Workspaces) > 0 {
		return j.Workspaces[0]
	}
	return ""
}

func (ctx *Context[T]) ApiTag() string {
	return ctx.GinCtx.Request.Method + ":" + ctx.GinCtx.FullPath()
}

func (ctx *Context[T]) LogPrefix() string {
	return fmt.Sprintf("[system: %s] [trace: %s] [api: %s] [ip: %s] [agent: %s]", ctx.Engine.GingerEngine.SystemUUID, ctx.TraceID(), ctx.ApiTag(), ctx.ClientIP(), ctx.UserAgent())
}

func (ctx *Context[T]) Traces() []types.Trace {
	traces := make([]types.Trace, 0)
	tracesHeader := ctx.GinCtx.GetHeader(ginger.HEADER_TRACE)
	if tracesHeader != "" {
		_ = json.Unmarshal([]byte(tracesHeader), &traces)
	}
	return traces
}

func (ctx *Context[T]) OK(data any, page ...*sql.Pagination) {
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Trace: gingerTypes.Trace{
			Success:    true,
			SystemUUID: ctx.Engine.GingerEngine.SystemUUID,
			SystemName: ctx.Engine.GingerEngine.SystemName,
			TraceID:    ctx.TraceID(),
			Time:       time.Now().Format("2006-01-02 15:04:05.000"),
			Duration:   time.Since(ctx.StartTime).Milliseconds(),
		},
		Credit: ctx.Credit,
	})
	var pageResponse *sql.Pagination
	if len(page) > 0 {
		pageResponse = page[0]
	}
	ctx.Response = &types.Response{
		Response: gingerTypes.Response{
			Success:    true,
			TraceID:    ctx.TraceID(),
			Locale:     ctx.Locale(),
			Duration:   time.Since(ctx.StartTime).Milliseconds(),
			Pagination: pageResponse,
			Data:       data,
		},
		Credit: ctx.Credit,
		Traces: traces,
	}
	ctx.Response.Calculate()
	ctx.IsResponded = true
}

func (ctx *Context[T]) Err(code string) {
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Trace: gingerTypes.Trace{
			Success:    false,
			SystemUUID: ctx.Engine.GingerEngine.SystemUUID,
			SystemName: ctx.Engine.GingerEngine.SystemName,
			TraceID:    ctx.TraceID(),
			Time:       time.Now().Format("2006-01-02 15:04:05.000"),
			Duration:   time.Since(ctx.StartTime).Milliseconds(),
			Error:      gingerTypes.NewError(code, ginger.ErrMap.Get(code, ctx.Locale())),
		},
		Credit: ctx.Credit,
	})
	ctx.Response = &types.Response{
		Response: gingerTypes.Response{
			Success:  false,
			TraceID:  ctx.TraceID(),
			Locale:   ctx.Locale(),
			Duration: time.Since(ctx.StartTime).Milliseconds(),
			Error:    gingerTypes.NewError(code, ginger.ErrMap.Get(code, ctx.Locale())),
		},
		Credit: ctx.Credit,
		Traces: traces,
	}
	ctx.Response.Calculate()
	ctx.IsResponded = true
}
