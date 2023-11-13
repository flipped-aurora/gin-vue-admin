package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ ICache = (*CacheRedis)(nil)

type CacheRedis struct {
	rds *redis.Client
}

func NewCacheRedis(client *redis.Client) *CacheRedis {
	return &CacheRedis{
		rds: client,
	}
}

func (c *CacheRedis) Get(ctx context.Context, key string) (string, error) {
	ret := c.rds.Get(ctx, key)
	r, err := ret.Result()
	switch {
	case err == redis.Nil:
		return "", ErrorKeyNotFound
	case err != nil:
		return "", err
	}
	return r, nil
}

func (c *CacheRedis) Set(ctx context.Context, key, value string) error {
	return c.rds.Set(ctx, key, value, redis.KeepTTL).Err()
}

func (c *CacheRedis) SetEx(ctx context.Context, key, value string, expiration time.Duration) error {
	return c.rds.SetEx(ctx, key, value, expiration).Err()
}

func (c *CacheRedis) Exist(ctx context.Context, key string) (bool, error) {
	result, err := c.rds.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (c *CacheRedis) Delete(ctx context.Context, key string) error {
	return c.rds.Del(ctx, key).Err()
}

func (c *CacheRedis) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.rds.Expire(ctx, key, expiration).Result()
}

func (c *CacheRedis) Ttl(ctx context.Context, key string) (int, error) {
	result, err := c.rds.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if result <= 0 {
		return int(result), nil
	}

	return int(result.Seconds()), nil
}

func (c *CacheRedis) Incr(ctx context.Context, key string) (int64, error) {
	return c.rds.Incr(ctx, key).Result()
}

func (c *CacheRedis) Decr(ctx context.Context, key string) (int64, error) {
	return c.rds.Decr(ctx, key).Result()
}
