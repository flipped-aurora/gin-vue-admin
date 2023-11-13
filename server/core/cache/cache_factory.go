package cache

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/redis/go-redis/v9"
)

type CreateCacheError error

func Create() (ICache, error) {
	mode := global.GVA_CONFIG.Cache.Model
	if len(mode) == 0 {
		return nil, CreateCacheError(errors.New("mode not found"))
	}

	switch mode {
	case "rds":
		redisCfg := global.GVA_CONFIG.Cache.Redis
		return createRedis(redisCfg.Addr, redisCfg.Password, redisCfg.DB)
	default:
		return nil, CreateCacheError(errors.New("mode not found"))
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
