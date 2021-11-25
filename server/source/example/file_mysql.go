package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var FileMysql = new(fileMysql)

type fileMysql struct{}

func (f *fileMysql) TableName() string {
	return "exa_file_upload_and_downloads"
}

func (f *fileMysql) Initialize() error {
	entities := []example.ExaFileUploadAndDownload{
		{Name: "10.png", Url: "https://qmplusimg.henrongyi.top/gvalogo.png", Tag: "png", Key: "158787308910.png"},
		{Name: "logo.png", Url: "https://qmplusimg.henrongyi.top/1576554439myAvatar.png", Tag: "png", Key: "1587973709logo.png"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, f.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (f *fileMysql) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("`name` = ? AND `key` = ?", "logo.png", "1587973709logo.png").First(&example.ExaFileUploadAndDownload{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
