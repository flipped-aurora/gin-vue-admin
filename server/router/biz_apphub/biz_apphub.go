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
		bizAppHubRouter.GET("getDeployList", bizAppHubApi.GetDeployList)                  // 获取部署历史
		bizAppHubRouter.POST("rollbackVersion", bizAppHubApi.RollbackVersion)             // 回滚版本
		//bizAppHubRouter.POST("api/cmd/call", bizAppHubApi.Call)                           // 调用命令行工具
	}
	{
		bizAppHubRouterWithoutRecord.GET("getUserInfo", bizAppHubApi.GetUserInfo)           // 获取用户信息
		bizAppHubRouterWithoutRecord.GET("findBizAppHub", bizAppHubApi.FindBizAppHub)       // 根据ID获取biz_apphub
		bizAppHubRouterWithoutRecord.GET("bizAppHub/record", bizAppHubApi.BizAppHubRecord)  // 根据ID获取biz_apphub
		bizAppHubRouterWithoutRecord.GET("getBizAppHubList", bizAppHubApi.GetBizAppHubList) // 获取biz_apphub列表
	}
	{
		bizAppHubRouterWithoutAuth.GET("getBizAppHubPublic", bizAppHubApi.GetBizAppHubPublic)       // 获取biz_apphub列表
		bizAppHubRouterWithoutAuth.GET("getUploadToken", bizAppHubApi.GetUploadToken)               // 获取上传token
		bizAppHubRouterWithoutAuth.POST("api/cmd/call/:user/:soft/:command", bizAppHubApi.PostCall) // 调用命令行工具
		bizAppHubRouterWithoutAuth.GET("api/cmd/call/:user/:soft/:command", bizAppHubApi.GetCall)   // 调用命令行工具
	}
}
