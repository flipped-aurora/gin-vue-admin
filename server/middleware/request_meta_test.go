package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

func TestRequestMetaInjectsAndResponds(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestMeta())
	var seen *logger.Fields
	r.GET("/ping", func(c *gin.Context) {
		seen = logger.FromCtx(c.Request.Context())
		c.Status(200)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	r.ServeHTTP(w, req)

	if w.Header().Get("X-Request-Id") == "" {
		t.Fatal("response missing X-Request-Id")
	}
	if seen == nil || seen.RequestID == "" || seen.HTTPPath != "/ping" || seen.HTTPMethod != http.MethodGet {
		t.Fatalf("ctx fields not injected: %+v", seen)
	}
}

func TestRequestMetaPassthrough(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestMeta())
	var seen *logger.Fields
	r.GET("/ping", func(c *gin.Context) {
		seen = logger.FromCtx(c.Request.Context())
		c.Status(200)
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	req.Header.Set("X-Request-Id", "upstream-id")
	r.ServeHTTP(w, req)
	if seen.RequestID != "upstream-id" || w.Header().Get("X-Request-Id") != "upstream-id" {
		t.Fatalf("upstream X-Request-Id not honored: %+v", seen)
	}
}

// serveWithHeaders 挂 RequestMeta 发一个带 headers 的请求,返回 ctx Fields 与响应
func serveWithHeaders(t *testing.T, headers map[string]string) (*logger.Fields, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestMeta())
	var seen *logger.Fields
	r.GET("/ping", func(c *gin.Context) {
		seen = logger.FromCtx(c.Request.Context())
		c.Status(200)
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	r.ServeHTTP(w, req)
	return seen, w
}

func TestRequestMetaGeneratesTrace(t *testing.T) {
	seen, w := serveWithHeaders(t, nil)
	if !logger.IsValidTraceID(seen.TraceID) {
		t.Fatalf("trace_id should be locally generated 32hex, got %q", seen.TraceID)
	}
	if seen.SpanID == "" || seen.ParentSpanID != "" {
		t.Fatalf("span fields wrong: %+v", seen)
	}
	if w.Header().Get("X-Trace-Id") != seen.TraceID {
		t.Fatal("response missing X-Trace-Id")
	}
	wantTP := logger.BuildTraceparent(seen.TraceID, seen.SpanID)
	if w.Header().Get("traceparent") != wantTP {
		t.Fatalf("traceparent header = %q, want %q", w.Header().Get("traceparent"), wantTP)
	}
}

func TestRequestMetaHonorsUpstreamTraceparent(t *testing.T) {
	seen, w := serveWithHeaders(t, map[string]string{
		"traceparent": "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01",
		"X-Trace-Id":  "should-be-ignored",
	})
	if seen.TraceID != "4bf92f3577b34da6a3ce929d0e0e4736" {
		t.Fatalf("upstream traceparent not honored: %q", seen.TraceID)
	}
	if seen.ParentSpanID != "00f067aa0ba902b7" {
		t.Fatalf("parent span not extracted: %q", seen.ParentSpanID)
	}
	if seen.SpanID == "" || seen.SpanID == seen.ParentSpanID {
		t.Fatalf("server span should be newly generated: %+v", seen)
	}
	if w.Header().Get("X-Trace-Id") != seen.TraceID {
		t.Fatal("response missing X-Trace-Id")
	}
}

func TestRequestMetaMalformedTraceparentFallsBack(t *testing.T) {
	// 畸形 traceparent 拒绝后落到 X-Trace-Id 透传
	seen, w := serveWithHeaders(t, map[string]string{
		"traceparent": "00-ZZZZ-bad-01",
		"X-Trace-Id":  "legacy-gateway-id",
	})
	if seen.TraceID != "legacy-gateway-id" {
		t.Fatalf("X-Trace-Id fallback not honored: %q", seen.TraceID)
	}
	// 非 W3C 格式的 trace_id 不得组装 traceparent 响应头
	if w.Header().Get("traceparent") != "" {
		t.Fatalf("traceparent should be omitted for non-W3C trace id, got %q", w.Header().Get("traceparent"))
	}
	if w.Header().Get("X-Trace-Id") != "legacy-gateway-id" {
		t.Fatal("response missing X-Trace-Id")
	}
}

func TestRequestMetaMalformedTraceparentGenerates(t *testing.T) {
	seen, _ := serveWithHeaders(t, map[string]string{
		"traceparent": "not-a-traceparent",
	})
	if !logger.IsValidTraceID(seen.TraceID) {
		t.Fatalf("malformed traceparent should trigger local generation, got %q", seen.TraceID)
	}
}

func TestRequestMetaRejectsInsaneHeaderIDs(t *testing.T) {
	long := strings.Repeat("x", 65) // 超过落库列宽 varchar(64)
	seen, _ := serveWithHeaders(t, map[string]string{
		"X-Trace-Id":   long,
		"X-Request-Id": "bad\nvalue",
	})
	if seen.TraceID == long || !logger.IsValidTraceID(seen.TraceID) {
		t.Fatalf("oversized X-Trace-Id should be replaced by local generation, got %q", seen.TraceID)
	}
	if seen.RequestID == "bad\nvalue" || len(seen.RequestID) != 36 {
		t.Fatalf("dirty X-Request-Id should be regenerated as uuid, got %q", seen.RequestID)
	}

	// 恰好 64 字符的干净值应放行
	ok64 := strings.Repeat("a", 64)
	seen2, _ := serveWithHeaders(t, map[string]string{"X-Trace-Id": ok64})
	if seen2.TraceID != ok64 {
		t.Fatalf("64-char clean X-Trace-Id should pass through, got %q", seen2.TraceID)
	}
}
