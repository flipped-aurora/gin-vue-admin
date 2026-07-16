package gva_cache

import (
	"testing"
	"time"
)

// 编译期保证 memoryCache 实现 Cache（接口断言）
var _ Cache = (*memoryCache)(nil)

func TestMemoryCache_SetGet(t *testing.T) {
	c := NewMemoryCache(time.Minute)
	c.Set("k1", 1, time.Minute)
	v, ok := c.Get("k1")
	if !ok {
		t.Fatalf("expected key k1 to exist")
	}
	if v.(int) != 1 {
		t.Fatalf("expected 1, got %v", v)
	}
	if _, ok := c.Get("missing"); ok {
		t.Fatalf("expected missing key to be absent")
	}
}

func TestMemoryCache_SetDefaultAndExists(t *testing.T) {
	c := NewMemoryCache(time.Minute)
	c.SetDefault("jwt-token", struct{}{})
	if !c.Exists("jwt-token") {
		t.Fatalf("expected jwt-token to exist after SetDefault")
	}
	c.Delete("jwt-token")
	if c.Exists("jwt-token") {
		t.Fatalf("expected jwt-token to be deleted")
	}
}

func TestMemoryCache_IncrementFromMissing(t *testing.T) {
	c := NewMemoryCache(time.Minute)
	// key 不存在时第一次 Increment 应建立计数并返回新值
	n, err := c.Increment("cnt", 1)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if n != 1 {
		t.Fatalf("expected 1, got %d", n)
	}
	n, err = c.Increment("cnt", 2)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if n != 3 {
		t.Fatalf("expected 3, got %d", n)
	}
}

func TestMemoryCache_IncrementWithExpire(t *testing.T) {
	c := NewMemoryCache(time.Minute)
	n, err := c.IncrementWithExpire("limit", 1, 50*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if n != 1 {
		t.Fatalf("expected 1, got %d", n)
	}
	n, err = c.IncrementWithExpire("limit", 1, 50*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if n != 2 {
		t.Fatalf("expected 2, got %d", n)
	}
	time.Sleep(120 * time.Millisecond)
	if c.Exists("limit") {
		t.Fatalf("expected limit key to expire")
	}
}
