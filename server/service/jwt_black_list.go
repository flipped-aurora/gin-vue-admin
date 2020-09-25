package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gorm.io/gorm"
	"time"
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

func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GVA_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
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

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := 60 * 60 * 24 * 7 * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
