package micro

import (
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/types"
	"github.com/metadiv-io/saas/utils"
	"github.com/metadiv-io/sql"
)

type MockContextParams[T any] struct {
	Request   *T
	Page      *sql.Pagination
	Sort      *sql.Sort
	Method    string
	Path      string
	ClientIP  string
	UserAgent string
	Headers   map[string]string
}

func MockContext[T any](params MockContextParams[T]) *Context[T] {
	w := httptest.NewRecorder()
	ctx, e := gin.CreateTestContext(w)

	if params.Method == "" {
		params.Method = "GET"
	}
	if params.Path == "" {
		params.Path = "/"
	}

	ctx.Request = httptest.NewRequest(params.Method, params.Path, nil)

	if params.ClientIP != "" {
		ctx.Request.RemoteAddr = params.ClientIP
	} else {
		ctx.Request.RemoteAddr = "127.0.0.1"
	}

	if params.UserAgent != "" {
		ctx.Request.Header.Set("User-Agent", params.UserAgent)
	} else {
		ctx.Request.Header.Set("User-Agent", "Mock Context")
	}

	if params.Headers != nil {
		for k, v := range params.Headers {
			ctx.Request.Header.Set(k, v)
		}
	}

	return &Context[T]{
		Context: ginger.Context[T]{
			GinCtx:    ctx,
			Request:   params.Request,
			Page:      params.Page,
			Sort:      params.Sort,
			StartTime: time.Now(),
		},
		Engine: &Engine{
			GingerEngine: &ginger.Engine{
				Gin: e,
			},
		},
	}
}

func MockContextWithAdminAuth[T any](params MockContextParams[T],
	adminUUID, adminUsername string) (*Context[T], *types.Jwt, string) {
	ctx := MockContext[T](params)

	// key pairs
	privPEM, pubPEM, err := utils.CreateRSAKeyPair()
	if err != nil {
		panic(err)
	}

	// create jwt
	j := &types.Jwt{}
	token, err := j.IssueAdminToken(1*time.Hour, privPEM,
		adminUUID, adminUsername, ctx.ClientIP(), ctx.UserAgent())
	if err != nil {
		panic(err)
	}

	// set public key
	ctx.Engine.PubPEM = pubPEM

	// set headers
	ctx.GinCtx.Request.Header.Set("Authorization", "Bearer "+token)

	return ctx, j, token
}

func MockContextWithUserAuth[T any](params MockContextParams[T],
	userUUID, userUsername string, workspaces []string) (*Context[T], *types.Jwt, string) {
	ctx := MockContext[T](params)

	// key pairs
	privPEM, pubPEM, err := utils.CreateRSAKeyPair()
	if err != nil {
		panic(err)
	}

	// create jwt
	j := &types.Jwt{}
	token, err := j.IssueUserToken(1*time.Hour, privPEM,
		userUUID, userUsername, ctx.ClientIP(), ctx.UserAgent(), workspaces)
	if err != nil {
		panic(err)
	}

	// set public key
	ctx.Engine.PubPEM = pubPEM

	// set headers
	ctx.GinCtx.Request.Header.Set("Authorization", "Bearer "+token)

	return ctx, j, token
}

func MockContextWithWorkspaceUserAuth[T any](params MockContextParams[T],
	workspaceUserUUID, workspaceUserUsername, workspaceUUID string) (*Context[T], *types.Jwt, string) {
	ctx := MockContext[T](params)

	// key pairs
	privPEM, pubPEM, err := utils.CreateRSAKeyPair()
	if err != nil {
		panic(err)
	}

	// create jwt
	j := &types.Jwt{}
	token, err := j.IssueWorkspaceUserToken(1*time.Hour, privPEM,
		workspaceUserUUID, workspaceUserUsername, ctx.ClientIP(), ctx.UserAgent(), workspaceUUID)
	if err != nil {
		panic(err)
	}

	// set public key
	ctx.Engine.PubPEM = pubPEM

	// set headers
	ctx.GinCtx.Request.Header.Set("Authorization", "Bearer "+token)

	return ctx, j, token
}
