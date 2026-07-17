package testutil

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func TestInitMemoryCache_AssignsGlobalAndRoundTrip(t *testing.T) {
	old := global.GVA_CACHE
	t.Cleanup(func() { global.GVA_CACHE = old })

	cache := InitMemoryCache(t, 0) // 0 走 DefaultCacheExpire
	if cache == nil {
		t.Fatal("InitMemoryCache 返回 nil")
	}
	if global.GVA_CACHE == nil {
		t.Fatal("global.GVA_CACHE 未被赋值")
	}
	if global.GVA_CACHE != cache {
		t.Fatal("global.GVA_CACHE 与返回值不一致")
	}

	cache.Set("testutil:rt", "v", time.Minute)
	got, ok := cache.Get("testutil:rt")
	if !ok {
		t.Fatal("Set 后 Get 失败")
	}
	if got != "v" {
		t.Fatalf("Get 值不匹配: %v", got)
	}
}

func TestInitMemoryCache_RestoresOldValueOnCleanup(t *testing.T) {
	// 先放一个非 nil 的 sentinel（用 NewMemoryCache 也是 Cache 接口实例）
	sentinel := InitMemoryCache(&fakeT{TB: t}, 0)
	// 上面的 fakeT 会赋值全局，但我们想手动控制，先拿回 sentinel 引用后还原
	old := global.GVA_CACHE
	global.GVA_CACHE = sentinel
	t.Cleanup(func() { global.GVA_CACHE = old })

	sub := &fakeT{TB: t}
	InitMemoryCache(sub, time.Hour)
	if global.GVA_CACHE == sentinel {
		t.Fatal("InitMemoryCache 未覆盖旧值")
	}
	sub.runCleanup()
	if global.GVA_CACHE != sentinel {
		t.Fatal("cleanup 未还原 global.GVA_CACHE")
	}
}
