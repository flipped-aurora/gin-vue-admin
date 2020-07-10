package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("customer").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		ApiRouter.POST("customer", v1.CreateExaCustomer)     // 创建客户
		ApiRouter.PUT("customer", v1.UpdateExaCustomer)      // 更新客户
		ApiRouter.DELETE("customer", v1.DeleteExaCustomer)   // 删除客户
		ApiRouter.GET("customer", v1.GetExaCustomer)         // 获取单一客户信息
		ApiRouter.GET("customerList", v1.GetExaCustomerList) // 获取客户列表
	}
}
