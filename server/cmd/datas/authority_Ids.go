package datas

import (
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string
	DataAuthorityIdAuthorityId string
}

var DataAuthorityId = []SysDataAuthorityId{
	{"888", "888"},
	{"888", "8881"},
	{"888", "9528"},
	{"9528", "8881"},
	{"9528", "9528"},
}

func InitSysDataAuthorityId(db *gorm.DB) {
	if err := db.Table("sys_data_authority_id").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ?", []string{"888", "9528"}).Find(&[]SysDataAuthorityId{}).RowsAffected == 5 {
			color.Danger.Println("sys_data_authority_id表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&DataAuthorityId).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_data_authority_id 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
