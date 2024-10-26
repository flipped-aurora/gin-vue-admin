package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliSetvipRouter struct{}

// InitCliSetvipRouter 初始化 团队设置 路由信息
func (s *CliSetvipRouter) InitCliSetvipRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	clisetvipRouter := Router.Group("clisetvip").Use(middleware.OperationRecord())
	clisetvipRouterWithoutRecord := Router.Group("clisetvip")
	clisetvipRouterWithoutAuth := PublicRouter.Group("clisetvip")
	{
		clisetvipRouter.POST("createCliSetvip", clisetvipApi.CreateCliSetvip)             // 新建团队设置
		clisetvipRouter.DELETE("deleteCliSetvip", clisetvipApi.DeleteCliSetvip)           // 删除团队设置
		clisetvipRouter.DELETE("deleteCliSetvipByIds", clisetvipApi.DeleteCliSetvipByIds) // 批量删除团队设置
		clisetvipRouter.PUT("updateCliSetvip", clisetvipApi.UpdateCliSetvip)              // 更新团队设置
	}
	{
		clisetvipRouterWithoutRecord.GET("findCliSetvip", clisetvipApi.FindCliSetvip)       // 根据ID获取团队设置
		clisetvipRouterWithoutRecord.GET("getCliSetvipList", clisetvipApi.GetCliSetvipList) // 获取团队设置列表
	}
	{
		clisetvipRouterWithoutAuth.GET("getCliSetvipPublic", clisetvipApi.GetCliSetvipPublic) // 团队设置开放接口
	}
}
