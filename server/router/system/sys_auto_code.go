package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

var AutoCodeRouterApp = new(AutoCodeRouter)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	publicAutoCodeRouter := RouterPublic.Group("autoCode")
	{
		api := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
		autoCodeRouter.GET("getDB", api.GetDB)         // 获取数据库
		autoCodeRouter.GET("getTables", api.GetTables) // 获取对应数据库的表
		autoCodeRouter.GET("getColumn", api.GetColumn) // 获取指定表所有字段信息
	}
	{
		api := v1.ApiGroupApp.SystemApiGroup.AutoCodeTemplate
		autoCodeRouter.POST("preview", api.Preview)   // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", api.Create) // 创建自动化代码
	}
	{
		api := v1.ApiGroupApp.SystemApiGroup.AutoCodePackage
		autoCodeRouter.POST("getPackage", api.All)       // 获取package包
		autoCodeRouter.POST("delPackage", api.Delete)    // 删除package包
		autoCodeRouter.POST("createPackage", api.Create) // 创建package包
	}
	{
		api := v1.ApiGroupApp.SystemApiGroup.AutoCodePlugin
		autoCodeRouter.POST("pubPlug", api.Packaged)      // 打包插件
		autoCodeRouter.POST("installPlugin", api.Install) // 自动安装插件
	}
	{
		api := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
		publicAutoCodeRouter.POST("llmAuto", api.LLMAuto)
	}
}
