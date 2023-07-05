package call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/types"
)

func get[T any](ctx *gin.Context, host string, path string,
	params map[string]string, headers map[string]string) (*Response[T], error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers[constant.MICRO_HEADER_TRACE_ID] = getTraceID(ctx)
	headers[constant.MICRO_HEADER_TRACES] = getTraces(ctx)
	headers[constant.MICRO_HEADER_WORKSPACE] = getWorkspaceUUID(ctx)
	headers[constant.MICRO_HEADER_LOCALE] = getLocale(ctx)
	headers["Authorization"] = getAuthToken(ctx)

	path += "?"
	for k, v := range params {
		path += k + "=" + v + "&"
	}
	path = path[:len(path)-1]

	req, err := http.NewRequest("GET", host+path, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response[T]
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("failed to unmarshal response body: ", string(bodyBytes))
		return nil, err
	}

	if ctx != nil {
		setTraceID(ctx, response.TraceID)
		setTraces(ctx, response.Traces)
	}
	return &response, nil
}

func nonGet[T any](ctx *gin.Context, host, path, method string,
	body interface{}, headers map[string]string) (*Response[T], error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers[constant.MICRO_HEADER_TRACE_ID] = getTraceID(ctx)
	headers[constant.MICRO_HEADER_TRACES] = getTraces(ctx)
	headers[constant.MICRO_HEADER_WORKSPACE] = getWorkspaceUUID(ctx)
	headers[constant.MICRO_HEADER_LOCALE] = getLocale(ctx)
	headers["Authorization"] = getAuthToken(ctx)

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, host+path, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response[T]
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("failed to unmarshal response body: ", string(bodyBytes))
		return nil, err
	}

	if ctx != nil {
		setTraceID(ctx, response.TraceID)
		setTraces(ctx, response.Traces)
	}
	return &response, nil
}

func getTraceID(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetHeader(constant.MICRO_HEADER_TRACE_ID)
}

func setTraceID(ctx *gin.Context, traceID string) {
	if ctx == nil {
		return
	}
	ctx.Request.Header.Set(constant.MICRO_HEADER_TRACE_ID, traceID)
}

func getTraces(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetHeader(constant.MICRO_HEADER_TRACES)
}

func setTraces(ctx *gin.Context, traces []types.Trace) {
	if ctx == nil {
		return
	}
	bytes, _ := json.Marshal(traces)
	ctx.Request.Header.Set(constant.MICRO_HEADER_TRACES, string(bytes))
}

func getWorkspaceUUID(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetHeader(constant.MICRO_HEADER_WORKSPACE)
}

func getLocale(ctx *gin.Context) string {
	if ctx == nil {
		return constant.LOCALE_EN
	}
	return ctx.GetHeader(constant.MICRO_HEADER_LOCALE)
}

func getAuthToken(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	token := ctx.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}
