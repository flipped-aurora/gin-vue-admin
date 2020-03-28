package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("t").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TestRouter.POST("createTest", api.CreateTest)     // 新建Test
		TestRouter.POST("deleteTest", api.DeleteTest)   //删除Test
		TestRouter.POST("updateTest", api.UpdateTest)   //更新Test
		TestRouter.POST("findTest ", api.FindTest)           // 根据ID获取Test
		TestRouter.POST("getTestList", api.GetTestList) //获取Test列表
	}
}
