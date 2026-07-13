package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestAccessLogEmitsAndCaches(t *testing.T) {
	gin.SetMode(gin.TestMode)
	core, logs := observer.New(zapcore.InfoLevel)
	global.GVA_LOG = zap.New(core)

	r := gin.New()
	r.Use(RequestMeta(), AccessLog())
	r.POST("/echo", func(c *gin.Context) {
		if c.GetString(ctxReqBodyKey) != `{"a":1}` {
			t.Errorf("req body cache wrong: %q", c.GetString(ctxReqBodyKey))
		}
		c.String(200, "pong")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(`{"a":1}`))
	r.ServeHTTP(w, req)

	if logs.Len() == 0 {
		t.Fatal("no access log emitted")
	}
	m := logs.All()[logs.Len()-1].ContextMap()
	if m["mod"] != "http" || m["http_status"].(int64) != 200 || m["http_path"] != "/echo" {
		t.Fatalf("access fields wrong: %+v", m)
	}
}

func TestAccessLogWideEventFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	core, logs := observer.New(zapcore.InfoLevel)
	global.GVA_LOG = zap.New(core)

	r := gin.New()
	r.Use(RequestMeta(), AccessLog())
	// 模拟 JWTAuth 已解析 claims 放入 gin context
	r.Use(func(c *gin.Context) {
		c.Set("claims", &systemReq.CustomClaims{
			BaseClaims: systemReq.BaseClaims{ID: 7, AuthorityId: 888},
		})
	})
	r.GET("/user/:id", func(c *gin.Context) { c.String(200, "pong") })

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/42?a=1", nil)
	r.ServeHTTP(w, req)

	m := logs.All()[logs.Len()-1].ContextMap()
	if m["http_route"] != "/user/:id" {
		t.Fatalf("http_route should be route template, got %v", m["http_route"])
	}
	if m["http_path"] != "/user/42" {
		t.Fatalf("http_path should keep raw path, got %v", m["http_path"])
	}
	if m["bytes_out"].(int64) != 4 || m["bytes_in"].(int64) != 0 {
		t.Fatalf("bytes wrong: in=%v out=%v", m["bytes_in"], m["bytes_out"])
	}
	if m["user_id"].(uint64) != 7 || m["authority_id"].(uint64) != 888 {
		t.Fatalf("identity wrong: user_id=%v authority_id=%v", m["user_id"], m["authority_id"])
	}
	if m["error"].(bool) != false {
		t.Fatalf("error flag wrong: %v", m["error"])
	}
	if _, ok := m["span_id"]; !ok {
		t.Fatal("span_id missing in access log")
	}
}

func TestAccessLogUnmatchedRouteAndError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	core, logs := observer.New(zapcore.InfoLevel)
	global.GVA_LOG = zap.New(core)

	r := gin.New()
	r.Use(RequestMeta(), AccessLog())

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/no/such/route", nil)
	r.ServeHTTP(w, req)

	m := logs.All()[logs.Len()-1].ContextMap()
	if m["http_route"] != "unmatched" {
		t.Fatalf("unmatched route should collapse to fixed value, got %v", m["http_route"])
	}
	if m["user_id"].(uint64) != 0 || m["authority_id"].(uint64) != 0 {
		t.Fatalf("anonymous identity should be 0: %+v", m)
	}

	// 5xx 置 error 标志
	r500 := gin.New()
	core2, logs2 := observer.New(zapcore.InfoLevel)
	global.GVA_LOG = zap.New(core2)
	r500.Use(RequestMeta(), AccessLog())
	r500.GET("/boom", func(c *gin.Context) { c.Status(500) })
	w2 := httptest.NewRecorder()
	r500.ServeHTTP(w2, httptest.NewRequest(http.MethodGet, "/boom", nil))
	m2 := logs2.All()[logs2.Len()-1].ContextMap()
	if m2["error"].(bool) != true {
		t.Fatalf("5xx should set error flag: %+v", m2)
	}
	// 仅 c.Status 无 body:gin Size() 为 -1,应归一化为 0
	if m2["bytes_out"].(int64) != 0 {
		t.Fatalf("no-body response should record bytes_out=0, got %v", m2["bytes_out"])
	}
}

func TestCaptureWriterCapsAtLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	buf := &bytes.Buffer{}
	w := captureWriter{ResponseWriter: c.Writer, buf: buf}

	// 填到距上限还剩 10 字节
	if _, err := w.Write(bytes.Repeat([]byte("a"), respCaptureLimit-10)); err != nil {
		t.Fatal(err)
	}
	// 跨界写:20 字节只捕获前 10,但必须完整透传给客户端
	n, err := w.Write(bytes.Repeat([]byte("b"), 20))
	if err != nil || n != 20 {
		t.Fatalf("straddling write must pass through fully: n=%d err=%v", n, err)
	}
	// 缓冲已满后继续写:不再捕获,依旧透传
	if _, err := w.Write([]byte("ccc")); err != nil {
		t.Fatal(err)
	}

	if buf.Len() != respCaptureLimit {
		t.Fatalf("capture buffer should cap at %d, got %d", respCaptureLimit, buf.Len())
	}
	tail := buf.Bytes()[respCaptureLimit-11:]
	if tail[0] != 'a' || tail[1] != 'b' || tail[10] != 'b' {
		t.Fatalf("boundary split wrong, tail=%q", tail)
	}
	total := respCaptureLimit - 10 + 20 + 3
	if rec.Body.Len() != total {
		t.Fatalf("client should receive all %d bytes, got %d", total, rec.Body.Len())
	}
}

func TestAccessLogLongStreamCapAndRealBytesOut(t *testing.T) {
	gin.SetMode(gin.TestMode)
	core, logs := observer.New(zapcore.InfoLevel)
	global.GVA_LOG = zap.New(core)

	const chunkSize, chunks = 64 * 1024, 20 // 1.25MB,超过 respCaptureLimit
	var buf *bytes.Buffer
	r := gin.New()
	r.Use(RequestMeta(), AccessLog())
	r.GET("/stream", func(c *gin.Context) {
		v, _ := c.Get(ctxRespBufferKey)
		buf = v.(*bytes.Buffer)
		chunk := bytes.Repeat([]byte("x"), chunkSize)
		for i := 0; i < chunks; i++ {
			if _, err := c.Writer.Write(chunk); err != nil {
				t.Errorf("stream write failed: %v", err)
			}
		}
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/stream", nil))

	if buf == nil {
		t.Fatal("resp buffer not cached in context")
	}
	if buf.Len() != respCaptureLimit {
		t.Fatalf("long stream capture should cap at %d, got %d", respCaptureLimit, buf.Len())
	}
	if w.Body.Len() != chunkSize*chunks {
		t.Fatalf("client should receive full stream %d, got %d", chunkSize*chunks, w.Body.Len())
	}
	// bytes_out 记真实发送量,不受捕获封顶影响
	m := logs.All()[logs.Len()-1].ContextMap()
	if m["bytes_out"].(int64) != int64(chunkSize*chunks) {
		t.Fatalf("bytes_out should be real sent size %d, got %v", chunkSize*chunks, m["bytes_out"])
	}
}

func TestAccessLogChunkedBytesIn(t *testing.T) {
	gin.SetMode(gin.TestMode)
	core, logs := observer.New(zapcore.InfoLevel)
	global.GVA_LOG = zap.New(core)

	r := gin.New()
	r.Use(RequestMeta(), AccessLog())
	r.POST("/echo", func(c *gin.Context) { c.String(200, "ok") })

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(`{"a":1}`))
	req.ContentLength = -1 // 模拟 chunked:无 Content-Length
	r.ServeHTTP(w, req)

	m := logs.All()[logs.Len()-1].ContextMap()
	if m["bytes_in"].(int64) != 7 {
		t.Fatalf("chunked request should record real body size 7, got %v", m["bytes_in"])
	}
}
