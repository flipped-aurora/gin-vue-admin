package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode")
	{
		AutoCodeRouter.POST("delSysHistory", v1.DelSysHistory) // 删除回滚记录
		AutoCodeRouter.POST("getMeta", v1.GetMeta)             // 根据id获取meta信息
		AutoCodeRouter.POST("getSysHistory", v1.GetSysHistory) // 获取回滚记录分页
		AutoCodeRouter.POST("rollback", v1.RollBack)           // 回滚
		AutoCodeRouter.POST("preview", v1.PreviewTemp)         // 获取自动创建代码预览
		AutoCodeRouter.POST("createTemp", v1.CreateTemp)       // 创建自动化代码
		AutoCodeRouter.GET("getTables", v1.GetTables)          // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", v1.GetDB)                  // 获取数据库
		AutoCodeRouter.GET("getColumn", v1.GetColumn)          // 获取指定表所有字段信息
	}
}
