// 自动生成模板SysDictionary
package model

import (
	"gorm.io/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type SysDictionary struct {
	gorm.Model
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:'字典名（中）'"`
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:'字典名（英）'"`
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:'状态'"`
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:'描述'"`
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

func SysDictionaryData() []SysDictionary {
	status := new(bool)
	*status = true
	return []SysDictionary{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
}