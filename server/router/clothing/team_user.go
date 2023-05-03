package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamUserRouter struct {
}

// InitTeamUserRouter 初始化 TeamUser 路由信息
func (s *TeamUserRouter) InitTeamUserRouter(Router *gin.RouterGroup) {
	teamUserRouter := Router.Group("teamUser").Use(middleware.OperationRecord())
	teamUserRouterWithoutRecord := Router.Group("teamUser")
	var teamUserApi = v1.ApiGroupApp.ClothingApiGroup.TeamUserApi
	{
		teamUserRouter.POST("createTeamUser", teamUserApi.CreateTeamUser)   // 新建TeamUser
		teamUserRouter.DELETE("deleteTeamUser", teamUserApi.DeleteTeamUser) // 删除TeamUser
		teamUserRouter.DELETE("deleteTeamUserByIds", teamUserApi.DeleteTeamUserByIds) // 批量删除TeamUser
		teamUserRouter.PUT("updateTeamUser", teamUserApi.UpdateTeamUser)    // 更新TeamUser
	}
	{
		teamUserRouterWithoutRecord.GET("findTeamUser", teamUserApi.FindTeamUser)        // 根据ID获取TeamUser
		teamUserRouterWithoutRecord.GET("getTeamUserList", teamUserApi.GetTeamUserList)  // 获取TeamUser列表
	}
}
