package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct {
}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	var authorityApi = v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.POST("delSysHistory", authorityApi.DelSysHistory) // 删除回滚记录
		autoCodeRouter.POST("getMeta", authorityApi.GetMeta)             // 根据id获取meta信息
		autoCodeRouter.POST("getSysHistory", authorityApi.GetSysHistory) // 获取回滚记录分页
		autoCodeRouter.POST("rollback", authorityApi.RollBack)           // 回滚
		autoCodeRouter.POST("preview", authorityApi.PreviewTemp)         // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", authorityApi.CreateTemp)       // 创建自动化代码
		autoCodeRouter.GET("getTables", authorityApi.GetTables)          // 获取对应数据库的表
		autoCodeRouter.GET("getDB", authorityApi.GetDB)                  // 获取数据库
		autoCodeRouter.GET("getColumn", authorityApi.GetColumn)          // 获取指定表所有字段信息
	}
}
