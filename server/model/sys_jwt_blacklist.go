package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text"`
}

func (j *JwtBlacklist) JsonInBlacklist() (err error) {
	err = global.GVA_DB.Create(j).Error
	return
}

//判断JWT是否在黑名单内部
func (j *JwtBlacklist) IsBlacklist(Jwt string) bool {
	isNotFound := global.GVA_DB.Where("jwt = ?", Jwt).First(j).RecordNotFound()
	return !isNotFound
}

//判断当前用户是否在线
func (j *JwtBlacklist) GetRedisJWT(userName string) (err error, RedisJWT string) {
	RedisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, RedisJWT
}

//设置当前用户在线
func (j *JwtBlacklist) SetRedisJWT(userName string) (err error) {
	err = global.GVA_REDIS.Set(userName, j.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
