package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMyCustomerRouter(Router *gin.RouterGroup) {
	// 管理员操作
	customerAddressRouter := Router.Group("/customerAddress").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		customerAddressRouter.POST("/address", api.GetAddressList)
		customerAddressRouter.POST("/add", api.AddAddress)
		customerAddressRouter.PATCH("/update", api.UpdateAddress)
		customerAddressRouter.DELETE("/delete", api.DeleteAddress)
	}

	customerRouter := Router.Group("/customer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		customerRouter.POST("list", api.GetCustomerList)
		customerRouter.PATCH("/update", api.UpdateCustomer)
		customerRouter.DELETE("/delete", api.DeleteCustomer)
		customerRouter.POST("/getbyid", api.GetCustomerById)
		customerRouter.POST("/add", api.AddCustomer)
	}

	customerOrderRouter := Router.Group("/customerOrder").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		customerOrderRouter.POST("/order", api.GetOrderList)
		customerOrderRouter.DELETE("/delOrder", api.DeleteOrder)
		customerOrderRouter.POST("/getbyid", api.GetOrderByOrderId)

	}
	// 客户操作
	customerBusinessRouter := Router.Group("/cus").Use(middleware.JWTAuth())
	{
		customerBusinessRouter.POST("/upload", api.UploadCusHeaderImg)
		customerBusinessRouter.POST("/cart", api.GetCartList)
		customerBusinessRouter.POST("/addcart", api.AddCart)
		customerBusinessRouter.POST("/reduceCart", api.ReduceCart)
		customerBusinessRouter.POST("/delcart", api.DelCart)
		customerBusinessRouter.POST("/checkStatus", api.CheckStatus)
		customerBusinessRouter.POST("/getAddress", api.GetAddressListByUserId)
		customerBusinessRouter.POST("/getOrder", api.GetOrderListByOrderType)
		customerBusinessRouter.POST("/addOrder", api.AddOrder)
		customerBusinessRouter.POST("/delOrder", api.DeleteOrder)
		customerBusinessRouter.POST("/orderDetail", api.GetOrderByOrderId)
	}
}
