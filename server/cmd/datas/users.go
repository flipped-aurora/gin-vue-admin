package datas

import (
	"time"

	"gin-vue-admin/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Users = []model.SysUser{
	{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/1571627762timg.jpg", AuthorityId: "888"},
	{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", NickName: "QMPlusUser", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
}

func InitSysUser(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&Users).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}
