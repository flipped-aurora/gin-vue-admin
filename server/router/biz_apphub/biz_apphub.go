package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizAppHubRouter struct{}

// InitBizAppHubRouter 初始化 biz_apphub 路由信息
func (s *BizAppHubRouter) InitBizAppHubRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	bizAppHubRouter := Router.Group("bizAppHub").Use(middleware.OperationRecord())
	bizAppHubRouterWithoutRecord := Router.Group("bizAppHub")
	bizAppHubRouterWithoutAuth := PublicRouter.Group("bizAppHub")
	{
		bizAppHubRouter.POST("createBizAppHub", bizAppHubApi.CreateBizAppHub)             // 新建biz_apphub
		bizAppHubRouter.DELETE("deleteBizAppHub", bizAppHubApi.DeleteBizAppHub)           // 删除biz_apphub
		bizAppHubRouter.DELETE("deleteBizAppHubByIds", bizAppHubApi.DeleteBizAppHubByIds) // 批量删除biz_apphub
		bizAppHubRouter.PUT("updateBizAppHub", bizAppHubApi.UpdateBizAppHub)              // 更新biz_apphub
	}
	{
		bizAppHubRouterWithoutRecord.GET("findBizAppHub", bizAppHubApi.FindBizAppHub)       // 根据ID获取biz_apphub
		bizAppHubRouterWithoutRecord.GET("getBizAppHubList", bizAppHubApi.GetBizAppHubList) // 获取biz_apphub列表
	}
	{
		bizAppHubRouterWithoutAuth.GET("getBizAppHubPublic", bizAppHubApi.GetBizAppHubPublic) // 获取biz_apphub列表
		bizAppHubRouterWithoutAuth.GET("getUploadToken", bizAppHubApi.GetUploadToken)         // 获取biz_apphub列表
	}
}
