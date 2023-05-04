package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JobQuestionRouter struct {
}

// InitJobQuestionRouter 初始化 JobQuestion 路由信息
func (s *JobQuestionRouter) InitJobQuestionRouter(Router *gin.RouterGroup) {
	jobQuestionRouter := Router.Group("jobQuestion").Use(middleware.OperationRecord())
	jobQuestionRouterWithoutRecord := Router.Group("jobQuestion")
	var jobQuestionApi = v1.ApiGroupApp.ClothingApiGroup.JobQuestionApi
	{
		jobQuestionRouter.POST("createJobQuestion", jobQuestionApi.CreateJobQuestion)   // 新建JobQuestion
		jobQuestionRouter.DELETE("deleteJobQuestion", jobQuestionApi.DeleteJobQuestion) // 删除JobQuestion
		jobQuestionRouter.DELETE("deleteJobQuestionByIds", jobQuestionApi.DeleteJobQuestionByIds) // 批量删除JobQuestion
		jobQuestionRouter.PUT("updateJobQuestion", jobQuestionApi.UpdateJobQuestion)    // 更新JobQuestion
	}
	{
		jobQuestionRouterWithoutRecord.GET("findJobQuestion", jobQuestionApi.FindJobQuestion)        // 根据ID获取JobQuestion
		jobQuestionRouterWithoutRecord.GET("getJobQuestionList", jobQuestionApi.GetJobQuestionList)  // 获取JobQuestion列表
	}
}
