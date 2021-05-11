package router

import (
	"github.com/eyotang/game-api-admin/server/api/v1"
	"github.com/eyotang/game-api-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system").Use(middleware.OperationRecord())
	{
		SystemRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("getServerInfo", v1.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("reloadSystem", v1.ReloadSystem)       // 重启服务
	}
}
