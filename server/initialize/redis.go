package initialize

import (
	"gin-vue-admin/global"
	"github.com/go-redis/redis"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error(err)
	} else {
		global.GVA_LOG.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
