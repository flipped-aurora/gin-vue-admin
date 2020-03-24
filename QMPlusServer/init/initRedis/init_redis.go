package initRedis

import (
	"gin-vue-admin/config"
	"gin-vue-admin/init/initlog"
	"github.com/go-redis/redis"
)

var DEFAULTREDIS *redis.Client

func InitRedis(logger log.Logger) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     config.GinVueAdminconfig.RedisAdmin.Addr,
		Password: config.GinVueAdminconfig.RedisAdmin.Password, // no password set
		DB:       config.GinVueAdminconfig.RedisAdmin.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("redis connect ping response:", pong)
		DEFAULTREDIS = client
	}
	return client
}
