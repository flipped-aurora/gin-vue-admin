package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Dictionary = new(dictionary)

type dictionary struct{}

func (d *dictionary) TableName() string {
	return "sys_dictionaries"
}

func (d *dictionary) Initialize() error {
	status := new(bool)
	*status = true
	entities := []system.SysDictionary{
		{Name: "性别", Type: "gender", Status: status, Desc: "性别字典"},
		{Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, d.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (d *dictionary) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("type = ?", "bool").First(&system.SysDictionary{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
