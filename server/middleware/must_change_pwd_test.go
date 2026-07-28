package middleware

import (
	"encoding/json"
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
	for _, path := range []string{
		"/user/changePassword",
		"/user/getUserInfo",
		"/jwt/jsonInBlacklist",
	} {
		t.Run("allow "+path, func(t *testing.T) {
			if w := setupRouterWithClaims(true, path); w.Code != http.StatusOK {
				t.Fatalf("%s should pass under must-change, got %d", path, w.Code)
			}
		})
	}
	// 强制改密 其它路径返回账号状态冲突和专用业务码
	if w := setupRouterWithClaims(true, "/user/getUserList"); w.Code != http.StatusConflict {
		t.Fatalf("other path should be 409 under must-change, got %d", w.Code)
	} else {
		var body struct {
			Code int `json:"code"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if body.Code != 10001 {
			t.Fatalf("must-change response code = %d, want 10001", body.Code)
		}
	}
}
