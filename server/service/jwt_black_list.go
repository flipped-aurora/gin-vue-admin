package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// @title    JsonInBlacklist
// @description   create jwt blacklist
// @param     jwtList         model.JwtBlacklist
// @auth                     （2020/04/05  20:22）
// @return    err             error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

// @title    IsBlacklist
// @description   check if the Jwt is in the blacklist or not, 判断JWT是否在黑名单内部
// @auth                     （2020/04/05  20:22）
// @param     jwt             string
// @param     jwtList         model.JwtBlacklist
// @return    err             error

func IsBlacklist(jwt string, jwtList model.JwtBlacklist) bool {
	isNotFound := global.GVA_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}

// @title    GetRedisJWT
// @description   Get user info in redis
// @auth                     （2020/04/05  20:22）
// @param     userName        string
// @return    err             error
// @return    redisJWT        string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

// @title    SetRedisJWT
// @description   set jwt into the Redis
// @auth                     （2020/04/05  20:22）
// @param     jwtList         model.JwtBlacklist
// @param     userName        string
// @return    err             error

func SetRedisJWT(jwtList model.JwtBlacklist, userName string) (err error) {
	err = global.GVA_REDIS.Set(userName, jwtList.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
