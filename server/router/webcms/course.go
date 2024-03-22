package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CourseRouter struct {
}

// InitCourseRouter 初始化 Course 路由信息
func (s *CourseRouter) InitCourseRouter(Router *gin.RouterGroup) {
	courseRouter := Router.Group("course").Use(middleware.OperationRecord())
	courseRouterWithoutRecord := Router.Group("course")
	var courseApi = v1.ApiGroupApp.WebcmsApiGroup.CourseApi
	{
		courseRouter.POST("createCourse", courseApi.CreateCourse)             // 新建Course
		courseRouter.DELETE("deleteCourse", courseApi.DeleteCourse)           // 删除Course
		courseRouter.DELETE("deleteCourseByIds", courseApi.DeleteCourseByIds) // 批量删除Course
		courseRouter.PUT("updateCourse", courseApi.UpdateCourse)              // 更新Course
	}
	{
		courseRouterWithoutRecord.GET("findCourse", courseApi.FindCourse)       // 根据ID获取Course
		courseRouterWithoutRecord.GET("getCourseList", courseApi.GetCourseList) // 获取Course列表
	}
}
