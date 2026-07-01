package system

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

func loginFailKey(username string) string {
	return "login_fail:" + username
}

func loginLockKey(username string) string {
	return "login_lock:" + username
}

// IsAccountLocked 查询账号是否处于锁定状态
func IsAccountLocked(ctx context.Context, username string) bool {
	return global.GVA_CACHE.Exists(loginLockKey(username))
}

// RecordLoginFail 记录一次登录失败 计数滚动窗口 TTL=LockDuration 达阈值置锁
func RecordLoginFail(ctx context.Context, username string, cfg system.SysSecurityConfig) {
	if !cfg.LockEnable {
		return
	}
	n, err := global.GVA_CACHE.IncrementWithExpire(loginFailKey(username), 1, cfg.LockDurationTimeout())
	if err != nil {
		logger.WithCtx(ctx).Mod("biz").Err(err).Error("登录失败计数失败")
		return
	}
	if int(n) >= cfg.LockThreshold {
		global.GVA_CACHE.Set(loginLockKey(username), 1, cfg.LockDurationTimeout())
	}
}

// ClearLoginFail 清除失败计数与锁 登录成功调用
func ClearLoginFail(ctx context.Context, username string) {
	global.GVA_CACHE.Delete(loginFailKey(username))
	global.GVA_CACHE.Delete(loginLockKey(username))
}

// IsPasswordExpired 纯函数 判定密码是否过期
func IsPasswordExpired(ctx context.Context, passwordUpdatedAt *time.Time, cfg system.SysSecurityConfig, now time.Time) bool {
	if !cfg.PwdExpireEnable || cfg.PwdExpireDays <= 0 {
		return false
	}
	if passwordUpdatedAt == nil {
		return false
	}
	deadline := passwordUpdatedAt.AddDate(0, 0, cfg.PwdExpireDays)
	return now.After(deadline)
}
