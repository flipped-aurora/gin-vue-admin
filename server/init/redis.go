package init

import (
	"gin-vue-admin/global"
	"github.com/go-redis/redis"
)

func RegisterRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     GinVueAdminconfig.RedisAdmin.Addr,
		Password: GinVueAdminconfig.RedisAdmin.Password, // no password set
		DB:       GinVueAdminconfig.RedisAdmin.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		L.Error(err)
	} else {
		L.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
