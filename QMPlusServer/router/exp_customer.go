package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("customer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("createExaCustomer", api.CreateExaCustomer)   // 创建客户
		ApiRouter.POST("updateExaCustomer", api.UpdateExaCustomer)   // 更新客户
		ApiRouter.POST("deleteExaCustomer", api.DeleteExaCustomer) // 删除客户
		ApiRouter.POST("getExaCustomer", api.GetExaCustomer) // 获取单一客户信息
		ApiRouter.POST("getExaCustomerList", api.GetExaCustomerList)   // 获取客户列表
	}
}
