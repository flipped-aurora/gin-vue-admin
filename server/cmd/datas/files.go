package datas

import (
	"gin-vue-admin/global"
	"github.com/gookit/color"
	"os"
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

var Files = []model.ExaFileUploadAndDownload{
	{global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
	{global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
}

func InitExaFileUploadAndDownload(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.ExaFileUploadAndDownload{}).RowsAffected == 2 {
			color.Danger.Println("exa_file_upload_and_downloads表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&Files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("[Mysql]-->初始化数据成功")
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> exa_file_upload_and_downloads 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
