package sysModel

import (
	"gin-vue-admin/init/qmsql"
	"github.com/jinzhu/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text"`
}

func(j *JwtBlacklist)JsonInBlacklist()(err error){
	err = qmsql.DEFAULTDB.Create(j).Error
	return
}


//判断JWT是否在黑名单内部
func (j *JwtBlacklist)IsBlacklist(Jwt string)(bool){
	isNotFound := qmsql.DEFAULTDB.Where("jwt = ?",Jwt).First(j).RecordNotFound()
	return !isNotFound
}