package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamRouter struct {
}

// InitTeamRouter 初始化 Team 路由信息
func (s *TeamRouter) InitTeamRouter(Router *gin.RouterGroup) {
	teamRouter := Router.Group("team").Use(middleware.OperationRecord())
	teamRouterWithoutRecord := Router.Group("team")
	h5TeamRouterWithoutRecord := Router.Group(global.GetAppApi() + "team")
	var teamApi = v1.ApiGroupApp.ClothingApiGroup.TeamApi
	{
		teamRouter.POST("createTeam", teamApi.CreateTeam)             // 新建Team
		teamRouter.DELETE("deleteTeam", teamApi.DeleteTeam)           // 删除Team
		teamRouter.DELETE("deleteTeamByIds", teamApi.DeleteTeamByIds) // 批量删除Team
		teamRouter.PUT("updateTeam", teamApi.UpdateTeam)              // 更新Team
	}
	{
		teamRouterWithoutRecord.GET("findTeam", teamApi.FindTeam)       // 根据ID获取Team
		teamRouterWithoutRecord.GET("getTeamList", teamApi.GetTeamList) // 获取Team列表
	}
	{
		h5TeamRouterWithoutRecord.GET("getTeamList", teamApi.GetTeamList)      // 获取Team列表
		h5TeamRouterWithoutRecord.POST("joinTeam", teamApi.JoinTeam)           // 组长拉用户进组
		h5TeamRouterWithoutRecord.DELETE("deleteMember", teamApi.DeleteMember) // 组长拉用户进组
	}
}
