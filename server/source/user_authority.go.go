package source

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var UserAuthority = new(userAuthority)

type userAuthority struct{}

var userAuthorityModel = []system.SysUseAuthority{
	{1, "888"},
	{1, "8881"},
	{1, "9528"},
	{2, "888"},
}

//@description: user_authority 数据初始化
func (a *userAuthority) Init() error {
	return global.GVA_DB.Model(&system.SysUseAuthority{}).Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_user_id IN (1, 2)").Find(&[]system.SysUseAuthority{}).RowsAffected == 4 {
			color.Danger.Println("\n[Mysql] --> sys_user_authority 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userAuthorityModel).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_user_authority 表初始数据成功!")
		return nil
	})
}
