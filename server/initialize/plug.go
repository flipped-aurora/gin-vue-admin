package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Plug interface {
	InitRouter(*gin.Engine) error
	InitModel(*gorm.DB) error
}

func InstallPlug(db *gorm.DB, router *gin.Engine, p ...Plug) (err error) {
	for _, v := range p {
		err = v.InitModel(db)
		if err != nil {
			return err
		}
		err = v.InitRouter(router)
		if err != nil {
			return err
		}
	}
	return nil
}
