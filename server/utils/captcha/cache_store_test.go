package captcha

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
)

// TestCacheStoreRoundTrip 用内存后端验证 Set/Get/Verify/clear 的完整往返语义。
func TestCacheStoreRoundTrip(t *testing.T) {
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Minute)
	defer func() { global.GVA_CACHE = nil }()

	s := NewCacheStore()
	if err := s.Set("id1", "1234"); err != nil {
		t.Fatalf("Set 返回错误: %v", err)
	}
	if got := s.Get("id1", false); got != "1234" {
		t.Fatalf("Get = %q, 期望 1234", got)
	}
	if s.Verify("id1", "0000", false) {
		t.Fatal("错误答案不应通过 Verify")
	}
	if !s.Verify("id1", "1234", true) {
		t.Fatal("正确答案应通过 Verify")
	}
	// clear=true 后应已删除
	if got := s.Get("id1", false); got != "" {
		t.Fatalf("clear 后 Get = %q, 期望空串", got)
	}
}

// TestCacheStoreNilCacheSafe GVA_CACHE 未就绪时不应 panic，且语义安全(Set 报错/Get 空/Verify false)。
func TestCacheStoreNilCacheSafe(t *testing.T) {
	global.GVA_CACHE = nil

	s := NewCacheStore()
	if err := s.Set("x", "y"); err == nil {
		t.Fatal("GVA_CACHE 为 nil 时 Set 应返回错误")
	}
	if got := s.Get("x", true); got != "" {
		t.Fatalf("GVA_CACHE 为 nil 时 Get = %q, 期望空串", got)
	}
	if s.Verify("x", "y", true) {
		t.Fatal("GVA_CACHE 为 nil 时 Verify 应为 false")
	}
}
