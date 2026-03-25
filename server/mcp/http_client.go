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
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type upstreamEnvelope[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
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

func requestTimeout() time.Duration {
	timeout := global.GVA_CONFIG.MCP.RequestTimeout
	if timeout <= 0 {
		timeout = 15
	}
	return time.Duration(timeout) * time.Second
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
	req.Header.Set(configuredAuthHeader(), token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

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
