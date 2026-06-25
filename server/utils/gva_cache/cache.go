package gva_cache

import "time"

// Cache 通用缓存抽象：有 Redis 用 Redis 后端，无 Redis 用内存后端。
// 该接口为跨计划钉死契约，不得改名 / 改签名。
type Cache interface {
	Get(key string) (any, bool)
	Set(key string, value any, ttl time.Duration)
	SetDefault(key string, value any)
	Increment(key string, n int64) (int64, error)
	IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error)
	Exists(key string) bool
	Delete(key string)
}
