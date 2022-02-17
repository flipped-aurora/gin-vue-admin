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
		{GVA_MODEL: global.GVA_MODEL{ID: 1}, Name: global.Translate("system.dictionary.gender"), Type: "gender", Status: status, Desc: global.Translate("system.dictionary.genderDict")},
		{GVA_MODEL: global.GVA_MODEL{ID: 2}, Name: global.Translate("system.dictionary.intType"), Type: "int", Status: status, Desc: global.Translate("system.dictionary.intTypeDict")},
		{GVA_MODEL: global.GVA_MODEL{ID: 3}, Name: global.Translate("system.dictionary.timeDateType"), Type: "time.Time", Status: status, Desc: global.Translate("system.dictionary.timeDateTypeDict")},
		{GVA_MODEL: global.GVA_MODEL{ID: 4}, Name: global.Translate("system.dictionary.floatType"), Type: "float64", Status: status, Desc: global.Translate("system.dictionary.floatType")},
		{GVA_MODEL: global.GVA_MODEL{ID: 5}, Name: global.Translate("system.dictionary.stringType"), Type: "string", Status: status, Desc: global.Translate("system.dictionary.stringType")},
		{GVA_MODEL: global.GVA_MODEL{ID: 6}, Name: global.Translate("system.dictionary.boolType"), Type: "bool", Status: status, Desc: global.Translate("system.dictionary.boolType")},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, d.TableName()+" "+"general.tabelDataInitFail")
	}
	return nil
}

func (d *dictionary) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("type = ?", "bool").First(&system.SysDictionary{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
