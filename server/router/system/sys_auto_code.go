package system

import (
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	publicAutoCodeRouter := RouterPublic.Group("autoCode")
	{
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)         // 获取数据库
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables) // 获取对应数据库的表
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn) // 获取指定表所有字段信息
	}
	{
		autoCodeRouter.POST("preview", autoCodeTemplateApi.Preview)   // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", autoCodeTemplateApi.Create) // 创建自动化代码
		autoCodeRouter.POST("addFunc", autoCodeTemplateApi.AddFunc)   // 为代码插入方法
	}
	{
		autoCodeRouter.POST("getPackage", autoCodePackageApi.All)       // 获取package包
		autoCodeRouter.POST("delPackage", autoCodePackageApi.Delete)    // 删除package包
		autoCodeRouter.POST("createPackage", autoCodePackageApi.Create) // 创建package包
	}
	{
		autoCodeRouter.GET("getTemplates", autoCodePackageApi.Templates) // 创建package包
	}
	{
		autoCodeRouter.POST("pubPlug", autoCodePluginApi.Packaged)      // 打包插件
		autoCodeRouter.POST("installPlugin", autoCodePluginApi.Install) // 自动安装插件

	}
	{
		publicAutoCodeRouter.POST("llmAuto", autoCodeApi.LLMAuto)
		publicAutoCodeRouter.POST("initMenu", autoCodePluginApi.InitMenu) // 同步插件菜单
		publicAutoCodeRouter.POST("initAPI", autoCodePluginApi.InitAPI)   // 同步插件API
	}
}
