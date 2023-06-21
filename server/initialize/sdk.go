package initialize

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/ali"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/cache"
)

func InitSDK() {
	//wechat.NewWechat()
	ali.NewAli()
	//NewDouyin()
}

func NewDouyin() {
	redisCfg := global.GVA_CONFIG.Redis
	redisOpts := &cache.RedisOpts{
		Host:        redisCfg.Addr,     // redis host
		Password:    redisCfg.Password, //redis password
		Database:    redisCfg.DB,       // redis db
		MaxActive:   10,                // 连接池最大活跃连接数
		MaxIdle:     10,                //连接池最大空闲连接数
		IdleTimeout: 60,                //空闲连接超时时间，单位：second
	}
	redisCache := cache.NewRedis(context.Background(), redisOpts)
	dy := douyin.New()
	cfg := &config.Config{
		ClientKey:    global.GVA_CONFIG.Douyin.Key,
		ClientSecret: global.GVA_CONFIG.Douyin.Secret,
		Cache:        redisCache,
	}
	openAPI := dy.GetOpenAPI(cfg)
	global.GVA_DOUYIN_OPEN = openAPI
}
