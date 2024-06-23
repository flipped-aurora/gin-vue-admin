package initialize

import (
	"gorm.io/gorm"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate()
}
