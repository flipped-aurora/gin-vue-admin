package initRedis

import (
	"fmt"
	"gin-vue-admin/config"
	"github.com/go-redis/redis"
)

var DEFAULTREDIS *redis.Client

func InitRedis() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     config.GinVueAdminconfig.RedisAdmin.Addr,
		Password: config.GinVueAdminconfig.RedisAdmin.Password, // no password set
		DB:       config.GinVueAdminconfig.RedisAdmin.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	} else {
		fmt.Println(pong, err)
		DEFAULTREDIS = client
	}
	return client
}
