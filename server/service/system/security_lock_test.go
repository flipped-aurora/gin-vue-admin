package system

import (
	"context"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
)

func TestIsPasswordExpired(t *testing.T) {
	now := time.Date(2026, 6, 25, 0, 0, 0, 0, time.UTC)
	cfg := system.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 90}

	old := now.AddDate(0, 0, -100)
	recent := now.AddDate(0, 0, -10)

	if !IsPasswordExpired(context.Background(), &old, cfg, now) {
		t.Fatalf("expected expired for 100 days old")
	}
	if IsPasswordExpired(context.Background(), &recent, cfg, now) {
		t.Fatalf("expected not expired for 10 days old")
	}
	if IsPasswordExpired(context.Background(), nil, cfg, now) {
		t.Fatalf("nil time should be not expired")
	}
	off := system.SysSecurityConfig{PwdExpireEnable: false, PwdExpireDays: 90}
	if IsPasswordExpired(context.Background(), &old, off, now) {
		t.Fatalf("disabled config should be not expired")
	}
}

func TestLockCounting(t *testing.T) {
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Hour)
	cfg := system.SysSecurityConfig{LockEnable: true, LockThreshold: 3, LockDuration: 30}
	user := "locktester"
	ctx := context.Background()

	ClearLoginFail(ctx, user)
	if IsAccountLocked(ctx, user) {
		t.Fatalf("should not be locked initially")
	}
	RecordLoginFail(ctx, user, cfg)
	RecordLoginFail(ctx, user, cfg)
	if IsAccountLocked(ctx, user) {
		t.Fatalf("should not be locked at 2 fails (threshold 3)")
	}
	RecordLoginFail(ctx, user, cfg)
	if !IsAccountLocked(ctx, user) {
		t.Fatalf("should be locked at 3 fails")
	}
	ClearLoginFail(ctx, user)
	if IsAccountLocked(ctx, user) {
		t.Fatalf("clear should remove lock")
	}
}
