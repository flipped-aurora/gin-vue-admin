package system

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestAccountLockFlow(t *testing.T) {
	global.GVA_LOG = zap.NewNop()
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Hour)
	cfg := sysModel.SysSecurityConfig{LockEnable: true, LockThreshold: 2, LockDuration: 30}
	user := "integration_lock_user"
	ctx := context.Background()

	systemService.ClearLoginFail(ctx, user)
	if systemService.IsAccountLocked(ctx, user) {
		t.Fatalf("should not be locked after clear")
	}
	systemService.RecordLoginFail(ctx, user, cfg)
	if systemService.IsAccountLocked(ctx, user) {
		t.Fatalf("should not be locked after 1 fail")
	}
	systemService.RecordLoginFail(ctx, user, cfg)
	if !systemService.IsAccountLocked(ctx, user) {
		t.Fatalf("should be locked after threshold reached")
	}
	systemService.ClearLoginFail(ctx, user)
	if systemService.IsAccountLocked(ctx, user) {
		t.Fatalf("should be unlocked after clear")
	}
}

func TestMustChangePwdMiddlewareBlocks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("claims", &systemReq.CustomClaims{MustChangePwd: true})
		c.Next()
	})
	r.Use(middleware.MustChangePwdGuard())
	r.POST("/user/getUserList", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.POST("/user/changePassword", func(c *gin.Context) { c.Status(http.StatusOK) })

	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, httptest.NewRequest(http.MethodPost, "/user/getUserList", nil))
	if w1.Code != http.StatusForbidden {
		t.Fatalf("getUserList under must-change should be 403, got %d", w1.Code)
	}

	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, httptest.NewRequest(http.MethodPost, "/user/changePassword", nil))
	if w2.Code != http.StatusOK {
		t.Fatalf("changePassword under must-change should pass, got %d", w2.Code)
	}
}

func TestPasswordExpiryFlow(t *testing.T) {
	now := time.Now()
	cfg := sysModel.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 30}
	old := now.AddDate(0, 0, -40)
	if !systemService.IsPasswordExpired(context.Background(), &old, cfg, now) {
		t.Fatalf("40 days old with 30d expiry should be expired")
	}
	fresh := now.AddDate(0, 0, -5)
	if systemService.IsPasswordExpired(context.Background(), &fresh, cfg, now) {
		t.Fatalf("5 days old should not be expired")
	}
}
