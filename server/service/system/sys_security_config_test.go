package system

import (
	"context"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 仅测试纯内存缓存读写 不触碰 DB
func TestSecurityConfigCurrentReadsCache(t *testing.T) {
	s := &SecurityConfigService{}
	want := system.SysSecurityConfig{PwdMinLength: 12, LockThreshold: 3}
	setSecurityConfigCache(want)
	got := s.Current(context.Background())
	if got.PwdMinLength != 12 || got.LockThreshold != 3 {
		t.Fatalf("Current() = %+v, want PwdMinLength=12 LockThreshold=3", got)
	}
}

func TestSetSecurityConfigCacheRoundTrip(t *testing.T) {
	c := system.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 45}
	setSecurityConfigCache(c)
	got := getSecurityConfigCache()
	if !got.PwdExpireEnable || got.PwdExpireDays != 45 {
		t.Fatalf("cache round trip failed: %+v", got)
	}
}
