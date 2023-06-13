package micro

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/types"
	"github.com/metadiv-io/saas/utils"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	Engine       *Engine
	GinCtx       *gin.Context
	Page         *sql.Pagination
	Sort         *sql.Sort
	Request      *T
	Response     *types.Response
	StartTime    time.Time
	Credit       float64
	ResponsePage *sql.Pagination
	IsResponded  bool
}

func NewContext[T any](engine *Engine, ginCtx *gin.Context, credit float64) *Context[T] {
	var page *sql.Pagination
	var sort *sql.Sort
	if ginCtx.Request.Method == "GET" {
		page = utils.GinRequest[sql.Pagination](ginCtx)
		sort = utils.GinRequest[sql.Sort](ginCtx)
	}
	ctx := &Context[T]{
		Engine:       engine,
		GinCtx:       ginCtx,
		Page:         page,
		Sort:         sort,
		Request:      utils.GinRequest[T](ginCtx),
		Response:     nil,
		StartTime:    time.Now(),
		Credit:       credit,
		ResponsePage: nil,
		IsResponded:  false,
	}
	ctx.TraceID() // generate trace id if not exist
	return ctx
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinCtx.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinCtx.Request.UserAgent()
}

func (ctx *Context[T]) TraceID() string {
	traceID := ctx.GinCtx.GetHeader(constant.MICRO_HEADER_TRACE_ID)
	if traceID == "" {
		traceID = uuid.NewString()
		ctx.SetTraceID(traceID)
	}
	return traceID
}

func (ctx *Context[T]) Traces() []types.Trace {
	traces := make([]types.Trace, 0)
	tracesHeader := ctx.GinCtx.GetHeader(constant.MICRO_HEADER_TRACES)
	if tracesHeader != "" {
		_ = json.Unmarshal([]byte(tracesHeader), &traces)
	}
	return traces
}

func (ctx *Context[T]) Workspace() string {
	return ctx.GinCtx.GetHeader(constant.MICRO_HEADER_WORKSPACE)
}

func (ctx *Context[T]) Locale() string {
	return ctx.GinCtx.GetHeader(constant.MICRO_HEADER_LOCALE)
}

func (ctx *Context[T]) SetTraceID(traceID string) {
	ctx.GinCtx.Request.Header.Set(constant.MICRO_HEADER_TRACE_ID, traceID)
}

func (ctx *Context[T]) SetTraces(traces []types.Trace) {
	bytes, _ := json.Marshal(traces)
	ctx.GinCtx.Request.Header.Set(constant.MICRO_HEADER_TRACES, string(bytes))
}

func (ctx *Context[T]) SetWorkspace(workspace string) {
	ctx.GinCtx.Request.Header.Set(constant.MICRO_HEADER_WORKSPACE, workspace)
}

func (ctx *Context[T]) SetLocale(locale string) {
	ctx.GinCtx.Request.Header.Set(constant.MICRO_HEADER_LOCALE, locale)
}

func (ctx *Context[T]) AuthToken() string {
	token := ctx.GinCtx.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}

func (ctx *Context[T]) AuthJwt() *types.Jwt {
	j := &types.Jwt{}
	j.ParseToken(ctx.Engine.PubPEM, ctx.AuthToken())
	if j.UUID == "" {
		return nil
	}
	return j
}

func (ctx *Context[T]) ApiTag() string {
	return ctx.GinCtx.Request.Method + " " + ctx.GinCtx.FullPath()
}

func (ctx *Context[T]) LogPrefix() string {
	return fmt.Sprintf("[system: %s] [trace: %s] [api: %s] [ip: %s] [agent: %s]", ctx.Engine.SystemUUID, ctx.TraceID(), ctx.ApiTag(), ctx.ClientIP(), ctx.UserAgent())
}

func (ctx *Context[T]) OK(data any) {
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Success:    true,
		SystemUUID: ctx.Engine.SystemUUID,
		SystemName: ctx.Engine.SystemName,
		TraceID:    ctx.TraceID(),
		Time:       time.Now().Unix(),
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
		Credit:     ctx.Credit,
	})
	ctx.Response = &types.Response{
		Success:    true,
		TraceID:    ctx.TraceID(),
		Locale:     ctx.Locale(),
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
		Credit:     ctx.Credit,
		Pagination: ctx.ResponsePage,
		Data:       data,
		Traces:     traces,
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
		Success:    false,
		SystemUUID: ctx.Engine.SystemUUID,
		SystemName: ctx.Engine.SystemName,
		TraceID:    ctx.TraceID(),
		Time:       time.Now().Unix(),
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
		Credit:     ctx.Credit,
		Error:      types.NewError(code, ErrMap.Get(code, ctx.Locale())),
	})
	ctx.Response = &types.Response{
		Success:  false,
		TraceID:  ctx.TraceID(),
		Locale:   ctx.Locale(),
		Duration: time.Since(ctx.StartTime).Milliseconds(),
		Credit:   ctx.Credit,
		Error:    types.NewError(code, ErrMap.Get(code, ctx.Locale())),
		Traces:   traces,
	}
	ctx.Response.Calculate()
	ctx.IsResponded = true
}
