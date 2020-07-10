package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Plug interface {
	InitRouter([2]*gin.RouterGroup) error
	InitModel(*gorm.DB) error
}

func InstallPlug(db *gorm.DB, router [2]*gin.RouterGroup, p Plug) (err error) {
	err = p.InitModel(db)
	if err != nil {
		return err
	}
	err = p.InitRouter(router)
	if err != nil {
		return err
	}
	return nil
}
