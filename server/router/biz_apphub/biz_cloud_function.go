package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizCloudFunctionRouter struct {}

// InitBizCloudFunctionRouter 初始化 云函数 路由信息
func (s *BizCloudFunctionRouter) InitBizCloudFunctionRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	bizCloudFunctionRouter := Router.Group("bizCloudFunction").Use(middleware.OperationRecord())
	bizCloudFunctionRouterWithoutRecord := Router.Group("bizCloudFunction")
	bizCloudFunctionRouterWithoutAuth := PublicRouter.Group("bizCloudFunction")
	{
		bizCloudFunctionRouter.POST("createBizCloudFunction", bizCloudFunctionApi.CreateBizCloudFunction)   // 新建云函数
		bizCloudFunctionRouter.DELETE("deleteBizCloudFunction", bizCloudFunctionApi.DeleteBizCloudFunction) // 删除云函数
		bizCloudFunctionRouter.DELETE("deleteBizCloudFunctionByIds", bizCloudFunctionApi.DeleteBizCloudFunctionByIds) // 批量删除云函数
		bizCloudFunctionRouter.PUT("updateBizCloudFunction", bizCloudFunctionApi.UpdateBizCloudFunction)    // 更新云函数
	}
	{
		bizCloudFunctionRouterWithoutRecord.GET("findBizCloudFunction", bizCloudFunctionApi.FindBizCloudFunction)        // 根据ID获取云函数
		bizCloudFunctionRouterWithoutRecord.GET("getBizCloudFunctionList", bizCloudFunctionApi.GetBizCloudFunctionList)  // 获取云函数列表
	}
	{
	    bizCloudFunctionRouterWithoutAuth.GET("getBizCloudFunctionPublic", bizCloudFunctionApi.GetBizCloudFunctionPublic)  // 获取云函数列表
	}
}
