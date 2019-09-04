package registTable

import (
	"github.com/jinzhu/gorm"
	"main/model/dbModel"
)

//注册数据库表专用
func RegistTable(db *gorm.DB) {
	db.AutoMigrate(dbModel.User{})
}
