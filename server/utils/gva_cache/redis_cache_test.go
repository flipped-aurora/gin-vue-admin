package gva_cache

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// 编译期保证 redisCache 实现 Cache（接口断言）
var _ Cache = (*redisCache)(nil)

// newTestRedis 尝试连本地 redis；连不上则跳过（CI/本地无 redis 时不阻塞）。
func newTestRedis(t *testing.T) redis.UniversalClient {
	t.Helper()
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		t.Skipf("跳过 redis 后端测试：无可用 redis (%v)", err)
	}
	return client
}

func TestRedisCache_RoundTrip(t *testing.T) {
	client := newTestRedis(t)
	c := NewRedisCache(client)
	key := "gva_cache_test:rt"
	c.Delete(key)

	c.Set(key, "v", time.Minute)
	v, ok := c.Get(key)
	if !ok || v.(string) != "v" {
		t.Fatalf("expected v, got %v ok=%v", v, ok)
	}
	if !c.Exists(key) {
		t.Fatalf("expected key to exist")
	}
	c.Delete(key)
	if c.Exists(key) {
		t.Fatalf("expected key deleted")
	}
}

func TestRedisCache_Increment(t *testing.T) {
	client := newTestRedis(t)
	c := NewRedisCache(client)
	key := "gva_cache_test:cnt"
	c.Delete(key)

	n, err := c.IncrementWithExpire(key, 1, time.Minute)
	if err != nil || n != 1 {
		t.Fatalf("expected 1, got %d err=%v", n, err)
	}
	n, err = c.Increment(key, 2)
	if err != nil || n != 3 {
		t.Fatalf("expected 3, got %d err=%v", n, err)
	}
	c.Delete(key)
}
