package NestInfo

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NestInfoRouter struct {
}

// InitNestInfoRouter 初始化 NestInfo 路由信息
func (s *NestInfoRouter) InitNestInfoRouter(Router *gin.RouterGroup) {
	nestinfoRouter := Router.Group("nestinfo").Use(middleware.OperationRecord())
	nestinfoRouterWithoutRecord := Router.Group("nestinfo")
	var nestinfoApi = v1.ApiGroupApp.NestInfoApiGroup.NestInfoApi
	{
		nestinfoRouter.POST("createNestInfo", nestinfoApi.CreateNestInfo)             // 新建NestInfo
		nestinfoRouter.DELETE("deleteNestInfo", nestinfoApi.DeleteNestInfo)           // 删除NestInfo
		nestinfoRouter.DELETE("deleteNestInfoByIds", nestinfoApi.DeleteNestInfoByIds) // 批量删除NestInfo
		nestinfoRouter.PUT("updateNestInfo", nestinfoApi.UpdateNestInfo)              // 更新NestInfo
	}
	{
		nestinfoRouterWithoutRecord.GET("findNestInfo", nestinfoApi.FindNestInfo)       // 根据ID获取NestInfo
		nestinfoRouterWithoutRecord.GET("getNestInfoList", nestinfoApi.GetNestInfoList) // 获取NestInfo列表
		nestinfoRouterWithoutRecord.GET("getAllUserList", nestinfoApi.GetAllUserList)   // 获取NestInfo列表
	}
}
