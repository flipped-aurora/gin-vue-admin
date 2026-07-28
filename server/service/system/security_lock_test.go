package system

import (
	"context"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
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

func TestShouldForcePasswordChange(t *testing.T) {
	now := time.Date(2026, time.July, 28, 0, 0, 0, 0, time.UTC)
	expired := now.AddDate(0, 0, -40)

	tests := []struct {
		name string
		user system.SysUser
		cfg  system.SysSecurityConfig
		want bool
	}{
		{
			name: "new user obligation",
			user: system.SysUser{MustChangePassword: true},
			want: true,
		},
		{
			name: "expired password",
			user: system.SysUser{PasswordUpdatedAt: &expired},
			cfg:  system.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 30},
			want: true,
		},
		{
			name: "no obligation",
			user: system.SysUser{PasswordUpdatedAt: &now},
			cfg:  system.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 30},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldForcePasswordChange(context.Background(), tt.user, tt.cfg, now)
			if got != tt.want {
				t.Fatalf("ShouldForcePasswordChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLockCounting(t *testing.T) {
	testutil.InitMemoryCache(t, 0)
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
