package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamApplyRouter struct {
}

// InitTeamApplyRouter 初始化 TeamApply 路由信息
func (s *TeamApplyRouter) InitTeamApplyRouter(Router *gin.RouterGroup) {
	teamApplyRouter := Router.Group("teamApply").Use(middleware.OperationRecord())
	teamApplyRouterWithoutRecord := Router.Group("teamApply")
	h5TeamApplyRouterWithoutRecord := Router.Group(global.GetAppApi() + "teamApply")
	var teamApplyApi = v1.ApiGroupApp.ClothingApiGroup.TeamApplyApi
	{
		teamApplyRouter.POST("createTeamApply", teamApplyApi.CreateTeamApply)             // 新建TeamApply
		teamApplyRouter.DELETE("deleteTeamApply", teamApplyApi.DeleteTeamApply)           // 删除TeamApply
		teamApplyRouter.DELETE("deleteTeamApplyByIds", teamApplyApi.DeleteTeamApplyByIds) // 批量删除TeamApply
		teamApplyRouter.PUT("updateTeamApply", teamApplyApi.UpdateTeamApply)              // 更新TeamApply
	}
	{
		teamApplyRouterWithoutRecord.GET("findTeamApply", teamApplyApi.FindTeamApply)       // 根据ID获取TeamApply
		teamApplyRouterWithoutRecord.GET("getTeamApplyList", teamApplyApi.GetTeamApplyList) // 获取TeamApply列表
	}
	{
		h5TeamApplyRouterWithoutRecord.GET("findTeamApply", teamApplyApi.FindTeamApply) // 审核申请
		h5TeamApplyRouterWithoutRecord.PUT("optApply", teamApplyApi.OptApply)           // 审核申请
	}
}
