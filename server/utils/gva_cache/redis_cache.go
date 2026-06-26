package gva_cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisCache 包装 go-redis，使缓存语义在多实例间共享。
type redisCache struct {
	client redis.UniversalClient
}

// NewRedisCache 创建 Redis 后端。
func NewRedisCache(client redis.UniversalClient) Cache {
	return &redisCache{client: client}
}

func (r *redisCache) Get(key string) (any, bool) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		// redis.Nil（不存在）与其他错误统一按「未命中」处理
		return nil, false
	}
	return val, true
}

func (r *redisCache) Set(key string, value any, ttl time.Duration) {
	r.client.Set(context.Background(), key, value, ttl)
}

func (r *redisCache) SetDefault(key string, value any) {
	// 无过期
	r.client.Set(context.Background(), key, value, 0)
}

func (r *redisCache) Increment(key string, n int64) (int64, error) {
	return r.client.IncrBy(context.Background(), key, n).Result()
}

func (r *redisCache) IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error) {
	ctx := context.Background()
	v, err := r.client.IncrBy(ctx, key, n).Result()
	if err != nil {
		return 0, err
	}
	// 首次计数时设置过期（窗口起点）
	if v == n {
		if err := r.client.Expire(ctx, key, ttl).Err(); err != nil {
			return v, err
		}
	}
	return v, nil
}

func (r *redisCache) Exists(key string) bool {
	cnt, err := r.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return cnt > 0
}

func (r *redisCache) Delete(key string) {
	r.client.Del(context.Background(), key)
}
