package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "main/docs"
	"main/middleware"
	"main/router"
)

//初始化总路由
func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Logger()) // 如果不需要日志 请关闭这里
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Router.Use(middleware.Logger())
	router.InitUserRouter(Router)                  // 注册用户路由
	router.InitBaseRouter(Router)                  // 注册基础功能路由
	router.InitMenuRouter(Router)                  // 注册menu路由
	router.InitAuthorityRouter(Router)             // 注册角色路由
	router.InitApiRouter(Router)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(Router) // 文件上传下载功能路由
	router.InitWorkflowRouter(Router)              // 工作流相关路由
	return Router
}
