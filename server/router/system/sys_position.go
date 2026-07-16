package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysPositionRouter struct{}

func (s *SysPositionRouter) InitSysPositionRouter(Router *gin.RouterGroup) {
	posRouter := Router.Group("position").Use(middleware.OperationRecord())
	posRouterWithoutRecord := Router.Group("position")
	{
		posRouter.POST("createPosition", positionApi.CreateSysPosition)   // 创建岗位
		posRouter.PUT("updatePosition", positionApi.UpdateSysPosition)    // 更新岗位
		posRouter.DELETE("deletePosition", positionApi.DeleteSysPosition) // 删除岗位
		posRouter.POST("setPositionUsers", positionApi.SetPositionUsers)  // 设置岗位成员(反向分配)
	}
	{
		posRouterWithoutRecord.POST("getPositionList", positionApi.GetSysPositionList) // 分页获取岗位列表
		posRouterWithoutRecord.GET("findPosition", positionApi.FindSysPosition)        // 根据ID获取岗位
		posRouterWithoutRecord.GET("getPositionUsers", positionApi.GetPositionUsers)   // 获取岗位成员ID列表
	}
}
