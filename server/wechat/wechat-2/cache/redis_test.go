package cache

import (
	"context"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	var (
		timeoutDuration = time.Second
		ctx             = context.Background()
		opts            = &RedisOpts{
			Host: "127.0.0.1:6379",
		}
		redis = NewRedis(ctx, opts)
		err   error
		val   = "silenceper"
		key   = "username"
	)
	redis.SetConn(redis.conn)
	redis.SetRedisCtx(ctx)

	if err = redis.Set(key, val, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist(key) {
		t.Error("IsExist Error")
	}

	name := redis.Get(key).(string)
	if name != val {
		t.Error("get Error")
	}

	if err = redis.Delete(key); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
