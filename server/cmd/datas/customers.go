package datas

import (
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

var Customers = []model.ExaCustomer{
	{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CustomerName: "测试客户", CustomerPhoneData: "1761111111", SysUserID: 1, SysUserAuthorityID: "888"},
}

func InitExaCustomer(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&Customers).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}
