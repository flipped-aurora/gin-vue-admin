package middleware

import (
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
)

func init() {
	// 测试环境注入 no-op 日志，避免 GVA_LOG nil panic
	global.GVA_LOG = zap.NewNop()
}

func TestDefaultCheckOrMark_LimitWithMemoryCache(t *testing.T) {
	// 注入内存后端，模拟无 Redis 也能限流
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Minute)
	key := "GVA_Limit_test_ip"
	global.GVA_CACHE.Delete(key)

	limit := 3
	// 前 3 次应放行
	for i := 1; i <= limit; i++ {
		if err := DefaultCheckOrMark(key, 60, limit); err != nil {
			t.Fatalf("第 %d 次不应被限流: %v", i, err)
		}
	}
	// 第 4 次应被限流
	if err := DefaultCheckOrMark(key, 60, limit); err == nil {
		t.Fatalf("第 4 次应被限流")
	}
	global.GVA_CACHE.Delete(key)
}

func TestDefaultCheckOrMark_FailOpenWhenCacheNil(t *testing.T) {
	global.GVA_CACHE = nil
	if err := DefaultCheckOrMark("any", 60, 1); err != nil {
		t.Fatalf("GVA_CACHE 为 nil 时应 fail-open 放行: %v", err)
	}
}
