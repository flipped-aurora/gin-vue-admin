package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

type JwtService struct{}

var JwtServiceApp = new(JwtService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(ctx context.Context, jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(&jwtList).Error
	if err != nil {
		return
	}
	global.GVA_CACHE.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(ctx context.Context, userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(ctx, userName).Result()
	return redisJWT, err
}

func LoadAll(ctx context.Context) {
	var data []string
	err := global.GVA_DB.WithContext(ctx).Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		logger.WithCtx(ctx).Mod("biz").Err(err).Error("加载数据库jwt黑名单失败!")
		return
	}
	for i := 0; i < len(data); i++ {
		global.GVA_CACHE.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 GVA_CACHE 中
}
