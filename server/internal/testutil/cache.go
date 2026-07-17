package testutil

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
)

// DefaultCacheExpire 是 InitMemoryCache 的默认过期时间。
// 与生产 InitGvaCache 语义对齐（取自 gva_cache 的常用默认值）。
const DefaultCacheExpire = time.Hour

// InitMemoryCache 用纯内存后端初始化 global.GVA_CACHE，
// 并在 t.Cleanup 中还原旧值。返回新建的 Cache 供调用方使用。
//
// defaultExpire 控制 SetDefault 的默认过期时间，传 0 时使用 DefaultCacheExpire。
// 适用于依赖 GVA_CACHE 但无需真实 Redis 的测试场景（大多数单元测试）。
//
// 注意：本函数会修改全局 GVA_CACHE，不应用于 t.Parallel() 并行测试。
// 并行场景请直接 gva_cache.NewMemoryCache(...) 持有局部实例。
//
// 用法：
//
//	func TestFoo(t *testing.T) {
//	    c := testutil.InitMemoryCache(t, 0)
//	    _ = c // global.GVA_CACHE 已同步赋值
//	}
func InitMemoryCache(t testing.TB, defaultExpire time.Duration) gva_cache.Cache {
	t.Helper()
	if defaultExpire <= 0 {
		defaultExpire = DefaultCacheExpire
	}
	cache := gva_cache.NewMemoryCache(defaultExpire)
	old := global.GVA_CACHE
	global.GVA_CACHE = cache
	t.Cleanup(func() { global.GVA_CACHE = old })
	return cache
}
