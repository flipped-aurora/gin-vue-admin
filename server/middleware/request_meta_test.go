package middleware

import (
	"net/http"
	"net/http/httptest"
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
