package registTable

import (
	"github.com/jinzhu/gorm"
	"main/model/dbModel"
)

func RegistTable(db *gorm.DB) {
	db.AutoMigrate(dbModel.User{})
}
