package example

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // 创建客户
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
