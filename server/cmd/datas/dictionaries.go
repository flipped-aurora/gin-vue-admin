package datas

import (
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

type SysDictionaryToPostgresql struct {
	gorm.Model
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:状态"`
	Description          string                `json:"description" form:"description" gorm:"column:description;comment:'描述'"`
	SysDictionaryDetails []model.SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

func InitSysDictionary(db *gorm.DB) (err error) {
	var status = new(bool)
	*status = true
	Dictionaries := []model.SysDictionary{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&Dictionaries).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

func InitSysDictionaryToPostgresql(db *gorm.DB) (err error) {
	status := new(bool)
	*status = true
	tx := db.Begin() // 开始事务
	insert := []SysDictionaryToPostgresql{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Description: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Description: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Description: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Description: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Description: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Description: "数据库bool类型"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}
