package sysModel

import (
	"gin-vue-admin/init/initRedis"
	"gin-vue-admin/init/qmsql"
	"github.com/jinzhu/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text"`
}

func (j *JwtBlacklist) JsonInBlacklist() (err error) {
	err = qmsql.DEFAULTDB.Create(j).Error
	return
}

//判断JWT是否在黑名单内部
func (j *JwtBlacklist) IsBlacklist(Jwt string) bool {
	isNotFound := qmsql.DEFAULTDB.Where("jwt = ?", Jwt).First(j).RecordNotFound()
	return !isNotFound
}

//判断当前用户是否在线
func (j *JwtBlacklist) GetRedisJWT(userName string) (err error, RedisJWT string) {
	RedisJWT, err = initRedis.DEFAULTREDIS.Get(userName).Result()
	return err, RedisJWT
}

//设置当前用户在线
func (j *JwtBlacklist) SetRedisJWT(userName string) (err error) {
	err = initRedis.DEFAULTREDIS.Set(userName, j.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
