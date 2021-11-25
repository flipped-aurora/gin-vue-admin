package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

func (d *dictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}

func (d *dictionaryDetail) Initialize() error {
	status := new(bool)
	*status = true
	entities := []system.SysDictionaryDetail{
		{Label: "男", Value: 1, Status: status, Sort: 1, SysDictionaryID: 1},
		{Label: "女", Value: 2, Status: status, Sort: 2, SysDictionaryID: 1},
		{Label: "smallint", Value: 1, Status: status, Sort: 1, SysDictionaryID: 2},
		{Label: "mediumint", Value: 2, Status: status, Sort: 2, SysDictionaryID: 2},
		{Label: "int", Value: 3, Status: status, Sort: 3, SysDictionaryID: 2},
		{Label: "bigint", Value: 4, Status: status, Sort: 4, SysDictionaryID: 2},
		{Label: "date", Status: status, SysDictionaryID: 3},
		{Label: "time", Value: 1, Status: status, Sort: 1, SysDictionaryID: 3},
		{Label: "year", Value: 2, Status: status, Sort: 2, SysDictionaryID: 3},
		{Label: "datetime", Value: 3, Status: status, Sort: 3, SysDictionaryID: 3},
		{Label: "timestamp", Value: 5, Status: status, Sort: 5, SysDictionaryID: 3},
		{Label: "float", Status: status, SysDictionaryID: 4},
		{Label: "double", Value: 1, Status: status, Sort: 1, SysDictionaryID: 4},
		{Label: "decimal", Value: 2, Status: status, Sort: 2, SysDictionaryID: 4},
		{Label: "char", Status: status, SysDictionaryID: 5},
		{Label: "varchar", Value: 1, Status: status, Sort: 1, SysDictionaryID: 5},
		{Label: "tinyblob", Value: 2, Status: status, Sort: 2, SysDictionaryID: 5},
		{Label: "tinytext", Value: 3, Status: status, Sort: 3, SysDictionaryID: 5},
		{Label: "text", Value: 4, Status: status, Sort: 4, SysDictionaryID: 5},
		{Label: "blob", Value: 5, Status: status, Sort: 5, SysDictionaryID: 5},
		{Label: "mediumblob", Value: 6, Status: status, Sort: 6, SysDictionaryID: 5},
		{Label: "mediumtext", Value: 7, Status: status, Sort: 7, SysDictionaryID: 5},
		{Label: "longblob", Value: 8, Status: status, Sort: 8, SysDictionaryID: 5},
		{Label: "longtext", Value: 9, Status: status, Sort: 9, SysDictionaryID: 5},
		{Label: "tinyint", Status: status, SysDictionaryID: 6},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, d.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (d *dictionaryDetail) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("id = ?", 23).First(&system.SysDictionaryDetail{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
