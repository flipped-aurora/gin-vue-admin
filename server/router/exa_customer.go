package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(Router *gin.RouterGroup) {
	CustomerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	{
		CustomerRouter.POST("customer", v1.CreateExaCustomer)     // 创建客户
		CustomerRouter.PUT("customer", v1.UpdateExaCustomer)      // 更新客户
		CustomerRouter.DELETE("customer", v1.DeleteExaCustomer)   // 删除客户
		CustomerRouter.GET("customer", v1.GetExaCustomer)         // 获取单一客户信息
		CustomerRouter.GET("customerList", v1.GetExaCustomerList) // 获取客户列表
	}
}
