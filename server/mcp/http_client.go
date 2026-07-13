package mcpTool

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

type upstreamEnvelope[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

// upstreamAuthHeader 是回打主服务时携带 token 的请求头名。主服务 JWT 中间件固定读取 x-token
// (见 utils/claims.go),与 MCP 入站可配置的 auth_header 解耦:入站用 configuredAuthHeader()
// 读外部客户端 token,出站一律用 x-token,否则运维一旦把 auth_header 配成非默认值,所有工具全部 401。
const upstreamAuthHeader = "x-token"

// parseOptionalBool 解析可选的布尔参数,兼容原生 bool 与字符串("true"/"false"/"1"/"0" 等);
// 缺省或非法返回 def。用于 MCP 客户端把布尔发成字符串、原生 bool 断言失败会静默回落的场景
func parseOptionalBool(v any, def bool) bool {
	switch value := v.(type) {
	case bool:
		return value
	case string:
		if b, err := strconv.ParseBool(strings.TrimSpace(value)); err == nil {
			return b
		}
	}
	return def
}

// truncateUpstreamBody 截断上游非 JSON 响应片段,便于在错误里暴露真实内容而不刷屏
func truncateUpstreamBody(b []byte) string {
	const max = 200
	s := strings.TrimSpace(string(b))
	if len(s) > max {
		return s[:max] + "…"
	}
	return s
}

func ResolveMCPServiceURL() string {
	baseURL := strings.TrimSpace(global.GVA_CONFIG.MCP.BaseURL)
	if baseURL != "" {
		return strings.TrimRight(baseURL, "/")
	}

	addr := global.GVA_CONFIG.MCP.Addr
	if addr <= 0 {
		addr = 8889
	}

	path := strings.TrimSpace(global.GVA_CONFIG.MCP.Path)
	if path == "" {
		path = "/mcp"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return fmt.Sprintf("http://127.0.0.1:%d%s", addr, path)
}

func upstreamBaseURL() string {
	baseURL := strings.TrimSpace(global.GVA_CONFIG.MCP.UpstreamBaseURL)
	if baseURL != "" {
		return strings.TrimRight(baseURL, "/")
	}

	return "http://127.0.0.1:8888"
}

// upstreamURL 基于上游 base url 拼接完整地址。upstreamBaseURL() 已 TrimRight,
// 此处只补齐 path 前导斜杠,不再重复裁剪。
func upstreamURL(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return upstreamBaseURL() + path
}

func requestTimeout() time.Duration {
	timeout := global.GVA_CONFIG.MCP.RequestTimeout
	if timeout <= 0 {
		timeout = 15
	}
	return time.Duration(timeout) * time.Second
}

// defaultUpstreamTimeout 是回打主服务公共(免鉴权)接口的默认超时。
const defaultUpstreamTimeout = 10 * time.Second

// fetchPublicUpstream 回打主服务的公共(免鉴权)接口:带超时的 GET + 标准信封解析。
// 动态 tool 与编排 prompt 的注册共用同一模式,T 为信封 Data 字段的具体负载类型。
func fetchPublicUpstream[T any](path string) (*upstreamEnvelope[T], error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), defaultUpstreamTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, upstreamURL(path), nil)
	if err != nil {
		return nil, fmt.Errorf("构建请求失败: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求上游接口失败: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("上游接口返回状态码 %d: %s", resp.StatusCode, string(rawBody))
	}

	var result upstreamEnvelope[T]
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("上游接口业务错误: %s", result.Msg)
	}
	return &result, nil
}

func getUpstream[T any](ctx context.Context, endpoint string, query url.Values) (*upstreamEnvelope[T], error) {
	return doUpstream[T](ctx, http.MethodGet, endpoint, query, nil)
}

func postUpstream[T any](ctx context.Context, endpoint string, body any) (*upstreamEnvelope[T], error) {
	return doUpstream[T](ctx, http.MethodPost, endpoint, nil, body)
}

func deleteUpstream[T any](ctx context.Context, endpoint string, body any) (*upstreamEnvelope[T], error) {
	return doUpstream[T](ctx, http.MethodDelete, endpoint, nil, body)
}

func doUpstream[T any](ctx context.Context, method, endpoint string, query url.Values, body any) (*upstreamEnvelope[T], error) {
	token := authTokenFromContext(ctx)
	if token == "" {
		return nil, fmt.Errorf("缺少MCP鉴权请求头: %s", configuredAuthHeader())
	}

	endpoint = strings.TrimSpace(endpoint)
	if endpoint == "" {
		return nil, fmt.Errorf("上游接口路径不能为空")
	}
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	baseURL := upstreamBaseURL()
	requestURL, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("构建上游请求地址失败: %w", err)
	}
	if len(query) > 0 {
		requestURL.RawQuery = query.Encode()
	}

	var reader io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("序列化上游请求失败: %w", err)
		}
		reader = bytes.NewReader(payload)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, requestTimeout())
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, method, requestURL.String(), reader)
	if err != nil {
		return nil, fmt.Errorf("创建上游请求失败: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set(upstreamAuthHeader, token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// 链路传播:外部 AI → MCP 进程 → GVA 主服务串成同一条 trace
	logger.InjectTraceHeaders(timeoutCtx, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求上游服务失败: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取上游响应失败: %w", err)
	}

	var result upstreamEnvelope[T]
	if len(rawBody) > 0 {
		if err := json.Unmarshal(rawBody, &result); err != nil {
			// 上游返回非 JSON(如 404/502 的 HTML 或网关错误页):先暴露真实状态码,
			// 不让解析错误掩盖真实的 HTTP 失败
			if resp.StatusCode >= http.StatusBadRequest {
				return nil, fmt.Errorf("上游请求失败，状态码: %d，响应: %s", resp.StatusCode, truncateUpstreamBody(rawBody))
			}
			return nil, fmt.Errorf("解析上游响应失败: %w", err)
		}
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if result.Msg != "" {
			return nil, errors.New(result.Msg)
		}
		return nil, fmt.Errorf("上游请求失败，状态码: %d", resp.StatusCode)
	}

	if result.Code != 0 {
		if result.Msg != "" {
			return nil, errors.New(result.Msg)
		}
		return nil, fmt.Errorf("上游请求失败，业务码: %d", result.Code)
	}

	return &result, nil
}

// doUpstreamRaw 是动态 tool 专用的上游调用：接受完整 path（已替换路径参数）、method、query、body，
// 返回原始响应字节（动态 tool 把响应原样包成 MCP text content 返回给外部 AI）。
// 认证头从 ctx 取（与 doUpstream 一致，由外部 AI 通过 MCP 请求头透传 token）。
func doUpstreamRaw(ctx context.Context, method, path string, query url.Values, body any) (int, []byte, error) {
	token := authTokenFromContext(ctx)
	if token == "" {
		return 0, nil, fmt.Errorf("缺少MCP鉴权请求头: %s", configuredAuthHeader())
	}

	path = strings.TrimSpace(path)
	if path == "" {
		return 0, nil, fmt.Errorf("上游接口路径不能为空")
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	requestURL, err := url.Parse(upstreamBaseURL() + path)
	if err != nil {
		return 0, nil, fmt.Errorf("构建上游请求地址失败: %w", err)
	}
	if len(query) > 0 {
		requestURL.RawQuery = query.Encode()
	}

	var reader io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return 0, nil, fmt.Errorf("序列化上游请求失败: %w", err)
		}
		reader = bytes.NewReader(payload)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, requestTimeout())
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, method, requestURL.String(), reader)
	if err != nil {
		return 0, nil, fmt.Errorf("创建上游请求失败: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set(upstreamAuthHeader, token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// 链路传播:外部 AI → MCP 进程 → GVA 主服务串成同一条 trace
	logger.InjectTraceHeaders(timeoutCtx, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("请求上游服务失败: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("读取上游响应失败: %w", err)
	}
	return resp.StatusCode, rawBody, nil
}
