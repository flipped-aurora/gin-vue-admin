package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
