package competition

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComInfoRouter struct {
}

// InitComInfoRouter 初始化 比赛信息 路由信息
func (s *ComInfoRouter) InitComInfoRouter(Router *gin.RouterGroup) {
	comDataRouter := Router.Group("comData").Use(middleware.OperationRecord())
	comDataRouterWithoutRecord := Router.Group("comData")
	var comDataApi = v1.ApiGroupApp.CompetitionApiGroup.ComInfoApi
	{
		comDataRouter.POST("createComInfo", comDataApi.CreateComInfo)   // 新建比赛信息
		comDataRouter.DELETE("deleteComInfo", comDataApi.DeleteComInfo) // 删除比赛信息
		comDataRouter.DELETE("deleteComInfoByIds", comDataApi.DeleteComInfoByIds) // 批量删除比赛信息
		comDataRouter.PUT("updateComInfo", comDataApi.UpdateComInfo)    // 更新比赛信息
	}
	{
		comDataRouterWithoutRecord.GET("findComInfo", comDataApi.FindComInfo)        // 根据ID获取比赛信息
		comDataRouterWithoutRecord.GET("getComInfoList", comDataApi.GetComInfoList)  // 获取比赛信息列表
	}
}
