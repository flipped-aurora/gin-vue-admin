package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JobApplyRouter struct {
}

// InitJobApplyRouter 初始化 JobApply 路由信息
func (s *JobApplyRouter) InitJobApplyRouter(Router *gin.RouterGroup) {
	jobApplyRouter := Router.Group("jobApply").Use(middleware.OperationRecord())
	jobApplyRouterWithoutRecord := Router.Group("jobApply")
	h5JobApplyRouterWithoutRecord := Router.Group(global.GetAppApi() + "jobApply")
	var jobApplyApi = v1.ApiGroupApp.ClothingApiGroup.JobApplyApi
	{
		jobApplyRouter.POST("createJobApply", jobApplyApi.CreateJobApply)             // 新建JobApply
		jobApplyRouter.DELETE("deleteJobApply", jobApplyApi.DeleteJobApply)           // 删除JobApply
		jobApplyRouter.DELETE("deleteJobApplyByIds", jobApplyApi.DeleteJobApplyByIds) // 批量删除JobApply
		jobApplyRouter.PUT("updateJobApply", jobApplyApi.UpdateJobApply)              // 更新JobApply
	}
	{
		jobApplyRouterWithoutRecord.GET("findJobApply", jobApplyApi.FindJobApply)       // 根据ID获取JobApply
		jobApplyRouterWithoutRecord.GET("getJobApplyList", jobApplyApi.GetJobApplyList) // 获取JobApply列表
	}
	{
		jobApplyRouterWithoutRecord.GET("findJobApply", jobApplyApi.FindJobApply)        // 根据ID获取JobApply
		h5JobApplyRouterWithoutRecord.POST("createJobApply", jobApplyApi.CreateJobApply) // 新建JobApply
	}
}
