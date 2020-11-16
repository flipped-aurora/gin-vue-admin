package datas

import (
	"gin-vue-admin/global"
	"github.com/gookit/color"
	"os"
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

var Customers = []model.ExaCustomer{
	{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CustomerName: "测试客户", CustomerPhoneData: "1761111111", SysUserID: 1, SysUserAuthorityID: "888"},
}

func InitExaCustomer(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ? ", []int{1}).Find(&[]model.ExaCustomer{}).RowsAffected == 1 {
			color.Danger.Println("exa_customers表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&Customers).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> exa_customers 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
