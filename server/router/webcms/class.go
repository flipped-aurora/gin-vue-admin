package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClassRouter struct {
}

// InitClassRouter 初始化 Class 路由信息
func (s *ClassRouter) InitClassRouter(Router *gin.RouterGroup) {
	classRouter := Router.Group("class").Use(middleware.OperationRecord())
	classRouterWithoutRecord := Router.Group("class")
	var classApi = v1.ApiGroupApp.WebcmsApiGroup.ClassApi
	{
		classRouter.POST("createClass", classApi.CreateClass)             // 新建Class
		classRouter.DELETE("deleteClass", classApi.DeleteClass)           // 删除Class
		classRouter.DELETE("deleteClassByIds", classApi.DeleteClassByIds) // 批量删除Class
		classRouter.PUT("updateClass", classApi.UpdateClass)              // 更新Class
	}
	{
		classRouterWithoutRecord.GET("findClass", classApi.FindClass)               // 根据ID获取Class
		classRouterWithoutRecord.GET("getClassList", classApi.GetClassList)         // 获取Class列表
		classRouterWithoutRecord.GET("getClassListById", classApi.GetClassListById) // 通过课程id获取Class列表
	}
}
