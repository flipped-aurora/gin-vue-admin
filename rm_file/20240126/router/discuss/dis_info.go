package discuss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DisInfoRouter struct {
}

// InitDisInfoRouter 初始化 帖子信息 路由信息
func (s *DisInfoRouter) InitDisInfoRouter(Router *gin.RouterGroup) {
	disDataRouter := Router.Group("disData").Use(middleware.OperationRecord())
	disDataRouterWithoutRecord := Router.Group("disData")
	var disDataApi = v1.ApiGroupApp.DiscussApiGroup.DisInfoApi
	{
		disDataRouter.POST("createDisInfo", disDataApi.CreateDisInfo)   // 新建帖子信息
		disDataRouter.DELETE("deleteDisInfo", disDataApi.DeleteDisInfo) // 删除帖子信息
		disDataRouter.DELETE("deleteDisInfoByIds", disDataApi.DeleteDisInfoByIds) // 批量删除帖子信息
		disDataRouter.PUT("updateDisInfo", disDataApi.UpdateDisInfo)    // 更新帖子信息
	}
	{
		disDataRouterWithoutRecord.GET("findDisInfo", disDataApi.FindDisInfo)        // 根据ID获取帖子信息
		disDataRouterWithoutRecord.GET("getDisInfoList", disDataApi.GetDisInfoList)  // 获取帖子信息列表
	}
}
