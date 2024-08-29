package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizCmdToolApiRouter struct{}

// InitBizCmdToolApiRouter 初始化 后端工具指令api 路由信息
func (s *BizCmdToolApiRouter) InitBizCmdToolApiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	bizCmdToolApiRouter := Router.Group("bizCmdToolApi").Use(middleware.OperationRecord())
	bizCmdToolApiRouterWithoutRecord := Router.Group("bizCmdToolApi")
	bizCmdToolApiRouterWithoutAuth := PublicRouter.Group("bizCmdToolApi")
	{
		bizCmdToolApiRouter.POST("createBizCmdToolApi", bizCmdToolApiApi.CreateBizCmdToolApi)             // 新建后端工具指令api
		bizCmdToolApiRouter.DELETE("deleteBizCmdToolApi", bizCmdToolApiApi.DeleteBizCmdToolApi)           // 删除后端工具指令api
		bizCmdToolApiRouter.DELETE("deleteBizCmdToolApiByIds", bizCmdToolApiApi.DeleteBizCmdToolApiByIds) // 批量删除后端工具指令api
		bizCmdToolApiRouter.PUT("updateBizCmdToolApi", bizCmdToolApiApi.UpdateBizCmdToolApi)              // 更新后端工具指令api
	}
	{
		bizCmdToolApiRouterWithoutRecord.GET("findBizCmdToolApi", bizCmdToolApiApi.FindBizCmdToolApi)       // 根据ID获取后端工具指令api
		bizCmdToolApiRouterWithoutRecord.GET("getBizCmdToolApiList", bizCmdToolApiApi.GetBizCmdToolApiList) // 获取后端工具指令api列表
	}
	{
		bizCmdToolApiRouterWithoutAuth.GET("getBizCmdToolApiPublic", bizCmdToolApiApi.GetBizCmdToolApiPublic) // 获取后端工具指令api列表
	}
}
