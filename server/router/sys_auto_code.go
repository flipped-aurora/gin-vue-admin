package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
		AutoCodeRouter.GET("getTables", v1.GetTables)    // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", v1.GetDB)            // 获取数据库
		AutoCodeRouter.GET("getColume", v1.GetColume)    // 获取指定表所有字段信息
	}
}
