package datas

import (
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

var Files = []model.ExaFileUploadAndDownload{
	{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
	{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
}

func InitExaFileUploadAndDownload(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&Files).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}
