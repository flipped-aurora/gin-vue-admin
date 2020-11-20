package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode")
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
		AutoCodeRouter.GET("getTables", v1.GetTables)    // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", v1.GetDB)            // 获取数据库
		AutoCodeRouter.GET("getColumn", v1.GetColumn)    // 获取指定表所有字段信息
	}
}
