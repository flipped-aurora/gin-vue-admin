package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	autoCodeApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)                  // 获取数据库
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables)          // 获取对应数据库的表
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn)          // 获取指定表所有字段信息
		autoCodeRouter.POST("preview", autoCodeApi.PreviewTemp)         // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", autoCodeApi.CreateTemp)       // 创建自动化代码
		autoCodeRouter.POST("createPackage", autoCodeApi.CreatePackage) // 创建package包
		autoCodeRouter.POST("getPackage", autoCodeApi.GetPackage)       // 获取package包
		autoCodeRouter.POST("delPackage", autoCodeApi.DelPackage)       // 删除package包
		autoCodeRouter.POST("createPlug", autoCodeApi.AutoPlug)         // 自动插件包模板
		autoCodeRouter.POST("installPlugin", autoCodeApi.InstallPlugin) // 自动安装插件
	}
}
