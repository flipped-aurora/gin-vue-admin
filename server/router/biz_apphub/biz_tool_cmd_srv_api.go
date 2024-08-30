package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizToolCmdSrvApiRouter struct{}

// InitBizToolCmdSrvApiRouter 初始化 后端工具指令api 路由信息
func (s *BizToolCmdSrvApiRouter) InitBizToolCmdSrvApiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	bizToolCmdSrvApiRouter := Router.Group("bizToolCmdSrvApi").Use(middleware.OperationRecord())
	bizToolCmdSrvApiRouterWithoutRecord := Router.Group("bizToolCmdSrvApi")
	bizToolCmdSrvApiRouterWithoutAuth := PublicRouter.Group("bizToolCmdSrvApi")
	{
		bizToolCmdSrvApiRouter.POST("createBizToolCmdSrvApi", bizToolCmdSrvApiApi.CreateBizToolCmdSrvApi)             // 新建后端工具指令api
		bizToolCmdSrvApiRouter.DELETE("deleteBizToolCmdSrvApi", bizToolCmdSrvApiApi.DeleteBizToolCmdSrvApi)           // 删除后端工具指令api
		bizToolCmdSrvApiRouter.DELETE("deleteBizToolCmdSrvApiByIds", bizToolCmdSrvApiApi.DeleteBizToolCmdSrvApiByIds) // 批量删除后端工具指令api
		bizToolCmdSrvApiRouter.PUT("updateBizToolCmdSrvApi", bizToolCmdSrvApiApi.UpdateBizToolCmdSrvApi)              // 更新后端工具指令api
	}
	{
		bizToolCmdSrvApiRouterWithoutRecord.GET("findBizToolCmdSrvApi", bizToolCmdSrvApiApi.FindBizToolCmdSrvApi)       // 根据ID获取后端工具指令api
		bizToolCmdSrvApiRouterWithoutRecord.GET("getBizToolCmdSrvApiList", bizToolCmdSrvApiApi.GetBizToolCmdSrvApiList) // 获取后端工具指令api列表
	}
	{
		bizToolCmdSrvApiRouterWithoutAuth.GET("getBizToolCmdSrvApiPublic", bizToolCmdSrvApiApi.GetBizToolCmdSrvApiPublic) // 获取后端工具指令api列表
	}
}
