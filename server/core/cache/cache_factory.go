package cache

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/redis/go-redis/v9"
)

func Create(c config.Cache) (ICache, error) {
	if len(c.Mode) == 0 {
		return nil, errors.New("mode not found")
	}

	switch c.Mode {
	case "redis":
		redisCfg := c.Redis
		return createRedis(redisCfg.Addr, redisCfg.Password, redisCfg.DB)
	default:
		return nil, errors.New("mode not found")
	}
}

func createRedis(addr, password string, db int) (ICache, error) {
	opt := &redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	}

	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return NewCacheRedis(client), nil
}
