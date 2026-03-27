package request

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

func HttpRequest(
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any) (*http.Response, error) {
	return doJSONRequest(context.Background(), 0, urlStr, method, headers, params, data)
}

// HttpRequestWithTimeout 发送HTTP请求，支持自定义超时时间
// timeout 参数可选，单位为时间.Duration，默认值为 10 分钟
func HttpRequestWithTimeout(
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any,
	timeout ...time.Duration) (*http.Response, error) {
	t := 10 * time.Minute
	if len(timeout) > 0 && timeout[0] > 0 {
		t = timeout[0]
	}
	return doJSONRequest(context.Background(), t, urlStr, method, headers, params, data)
}

// HttpRequestWithContextAndTimeout 发送HTTP请求，支持自定义超时时间和上下文
func HttpRequestWithContextAndTimeout(
	ctx context.Context,
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any,
	timeout ...time.Duration) (*http.Response, error) {
	t := 10 * time.Minute
	if len(timeout) > 0 {
		if timeout[0] < 0 {
			t = 0 // 负值表示不设置超时（用于 SSE 等流式场景）
		} else if timeout[0] > 0 {
			t = timeout[0]
		}
	}
	return doJSONRequest(ctx, t, urlStr, method, headers, params, data)
}

func doJSONRequest(
	ctx context.Context,
	timeout time.Duration,
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	// 将数据编码为JSON
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = timeout
	}

	// 当请求 SSE 流时，禁用 Transport 层的自动 gzip 压缩
	// 避免 gzip 解压缓冲导致 SSE 事件无法实时到达
	if req.Header.Get("Accept") == "text/event-stream" {
		client.Transport = &http.Transport{
			DisableCompression: true,
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 返回响应，调用方负责关闭
	return resp, nil
}
