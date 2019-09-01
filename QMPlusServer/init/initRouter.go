package init

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router = gin.Default()

func InitRouter() {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Router.Use(middleware.Logger())
}
