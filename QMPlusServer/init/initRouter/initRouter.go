package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "main/docs"
	"main/router"
)

//初始化总路由
func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Router.Use(middleware.Logger())
	router.InitUserRouter(Router)
	return Router
}
