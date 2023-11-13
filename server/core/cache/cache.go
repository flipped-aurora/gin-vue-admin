package cache

import (
	"context"
	"errors"
	"time"
)

var ErrorKeyNotFound = errors.New("key not found")

type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
	SetEx(ctx context.Context, key, value string, expiration time.Duration) error
	Exist(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Ttl(ctx context.Context, key string) (int, error) // -1 永不过期， -2 没有找到key
}
