package micro

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/ginger/context"
	gingerTypes "github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/types"
	"github.com/metadiv-io/sql"
)

type IContext[T any] interface {
	Engine() IEngine
	GinCtx() *gin.Context
	Page() *sql.Pagination
	Sort() *sql.Sort
	Request() *T
	Response() *types.Response
	SetResponse(resp *types.Response)
	StartTime() time.Time
	IsFile() bool
	IsResponded() bool
	SetIsResponded(isResp bool)
	ClientIP() string
	UserAgent() string
	SetTraceID(traceID string)
	TraceID() string
	Locale() string
	SetLocale(locale string)
	Traces() []types.Trace
	SetTraces(traces []types.Trace)
	BearerToken() string
	OK(data any, page ...*sql.Pagination)
	Err(code string)
	Credit() float64
	AuthJwt() *types.Jwt
	Workspace() string
	ApiTag() string
	LogPrefix() string
}

type Context[T any] struct {
	context.IContext[T]
	engine   IEngine
	response *types.Response
	credit   float64
}

func NewContext[T any](e IEngine, ginCtx *gin.Context, credit float64) IContext[T] {
	ctx := &Context[T]{
		IContext: ginger.NewContext[T](e, ginCtx),
		engine:   e,
		credit:   credit,
	}
	return ctx
}

func (ctx *Context[T]) Engine() IEngine {
	return ctx.engine
}

func (ctx *Context[T]) Credit() float64 {
	return ctx.credit
}

func (ctx *Context[T]) SetResponse(resp *types.Response) {
	ctx.response = resp
}

func (ctx *Context[T]) Response() *types.Response {
	return ctx.response
}

func (ctx *Context[T]) SetTraces(traces []types.Trace) {
	ctx.response.Traces = traces
}

func (ctx *Context[T]) AuthJwt() *types.Jwt {
	j := &types.Jwt{}
	err := j.ParseToken(ctx.Engine().PubPEM(), ctx.BearerToken())
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
	return ctx.GinCtx().Request.Method + ":" + ctx.GinCtx().FullPath()
}

func (ctx *Context[T]) LogPrefix() string {
	return fmt.Sprintf("[system: %s] [trace: %s] [api: %s] [ip: %s] [agent: %s]", ctx.Engine().SystemUUID(), ctx.TraceID(), ctx.ApiTag(), ctx.ClientIP(), ctx.UserAgent())
}

func (ctx *Context[T]) Traces() []types.Trace {
	traces := make([]types.Trace, 0)
	tracesHeader := ctx.GinCtx().GetHeader(ginger.HEADER_TRACE)
	if tracesHeader != "" {
		_ = json.Unmarshal([]byte(tracesHeader), &traces)
	}
	return traces
}

func (ctx *Context[T]) OK(data any, page ...*sql.Pagination) {
	if ctx.IsResponded() {
		log.Println("Warning: context already responded")
		return
	}
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Trace: gingerTypes.Trace{
			Success:    true,
			SystemUUID: ctx.Engine().SystemUUID(),
			SystemName: ctx.Engine().SystemName(),
			TraceID:    ctx.TraceID(),
			Time:       time.Now().Format("2006-01-02 15:04:05.000"),
			Duration:   time.Since(ctx.StartTime()).Milliseconds(),
		},
		Credit: ctx.Credit(),
	})
	var pageResponse *sql.Pagination
	if len(page) > 0 {
		pageResponse = page[0]
	}
	ctx.SetResponse(&types.Response{
		Response: gingerTypes.Response{
			Success:    true,
			TraceID:    ctx.TraceID(),
			Locale:     ctx.Locale(),
			Duration:   time.Since(ctx.StartTime()).Milliseconds(),
			Pagination: pageResponse,
			Data:       data,
		},
		Credit: ctx.Credit(),
		Traces: traces,
	})
	ctx.Response().Calculate()
	ctx.SetIsResponded(true)
}

func (ctx *Context[T]) Err(code string) {
	if ctx.IsResponded() {
		log.Println("Warning: context already responded")
		return
	}
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Trace: gingerTypes.Trace{
			Success:    false,
			SystemUUID: ctx.Engine().SystemUUID(),
			SystemName: ctx.Engine().SystemName(),
			TraceID:    ctx.TraceID(),
			Time:       time.Now().Format("2006-01-02 15:04:05.000"),
			Duration:   time.Since(ctx.StartTime()).Milliseconds(),
			Error:      gingerTypes.NewError(code, ginger.GetError(code, ctx.Locale())),
		},
		Credit: ctx.Credit(),
	})
	ctx.SetResponse(&types.Response{
		Response: gingerTypes.Response{
			Success:  false,
			TraceID:  ctx.TraceID(),
			Locale:   ctx.Locale(),
			Duration: time.Since(ctx.StartTime()).Milliseconds(),
			Error:    gingerTypes.NewError(code, ginger.GetError(code, ctx.Locale())),
		},
		Credit: ctx.Credit(),
		Traces: traces,
	})
	ctx.Response().Calculate()
	ctx.SetIsResponded(true)
}
