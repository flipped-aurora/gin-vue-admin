package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup) {
	TestTestRouter := Router.Group("t").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TestTestRouter.POST("createTestTest", v1.CreateTestTest)     // 新建TestTest
		TestTestRouter.DELETE("deleteTestTest", v1.DeleteTestTest)   //删除TestTest
		TestTestRouter.PUT("updateTestTest", v1.UpdateTestTest)   //更新TestTest
		TestTestRouter.GET("findTestTest", v1.FindTestTest)           // 根据ID获取TestTest
		TestTestRouter.GET("getTestTestList", v1.GetTestTestList) //获取TestTest列表
	}
}
