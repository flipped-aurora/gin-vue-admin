package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

func setupRouterWithClaims(mustChange bool, path string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("claims", &systemReq.CustomClaims{MustChangePwd: mustChange})
		c.Next()
	})
	r.Use(MustChangePwdGuard())
	r.Any("/*any", func(c *gin.Context) { c.Status(http.StatusOK) })
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, path, nil)
	r.ServeHTTP(w, req)
	return w
}

func TestMustChangePwdGuard(t *testing.T) {
	// 未强制改密 任意路径放行
	if w := setupRouterWithClaims(false, "/user/getUserList"); w.Code != http.StatusOK {
		t.Fatalf("not-must-change should pass, got %d", w.Code)
	}
	// 强制改密 改密路径放行
	if w := setupRouterWithClaims(true, "/user/changePassword"); w.Code != http.StatusOK {
		t.Fatalf("changePassword should pass under must-change, got %d", w.Code)
	}
	// 强制改密 用户信息放行
	if w := setupRouterWithClaims(true, "/user/getUserInfo"); w.Code != http.StatusOK {
		t.Fatalf("getUserInfo should pass under must-change, got %d", w.Code)
	}
	// 强制改密 其它路径 403
	if w := setupRouterWithClaims(true, "/user/getUserList"); w.Code != http.StatusForbidden {
		t.Fatalf("other path should be 403 under must-change, got %d", w.Code)
	}
}
