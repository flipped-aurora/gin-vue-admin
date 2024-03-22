package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RecruitRouter struct {
}

// InitRecruitRouter 初始化 Recruit 路由信息
func (s *RecruitRouter) InitRecruitRouter(Router *gin.RouterGroup) {
	recruitRouter := Router.Group("recruit").Use(middleware.OperationRecord())
	recruitRouterWithoutRecord := Router.Group("recruit")
	var recruitApi = v1.ApiGroupApp.WebcmsApiGroup.RecruitApi
	{
		recruitRouter.POST("createRecruit", recruitApi.CreateRecruit)   // 新建Recruit
		recruitRouter.DELETE("deleteRecruit", recruitApi.DeleteRecruit) // 删除Recruit
		recruitRouter.DELETE("deleteRecruitByIds", recruitApi.DeleteRecruitByIds) // 批量删除Recruit
		recruitRouter.PUT("updateRecruit", recruitApi.UpdateRecruit)    // 更新Recruit
	}
	{
		recruitRouterWithoutRecord.GET("findRecruit", recruitApi.FindRecruit)        // 根据ID获取Recruit
		recruitRouterWithoutRecord.GET("getRecruitList", recruitApi.GetRecruitList)  // 获取Recruit列表
	}
}
