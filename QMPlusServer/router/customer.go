package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMyCustomerRouter(Router *gin.RouterGroup) {

	// 客户操作
	customerAddressRouter := Router.Group("/customerAddress").Use(middleware.JWTAuth())
	{
		customerAddressRouter.POST("/add", api.AddAddress)
		customerAddressRouter.PATCH("/update", api.UpdateAddress)
		customerAddressRouter.DELETE("/delete", api.DeleteAddress)
	}
	// 管理员操作
	customerRouter := Router.Group("/customer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		customerAddressRouter.GET("/address", api.GetAddressList)
		customerRouter.POST("list", api.GetCustomerList)
		customerRouter.PATCH("/update", api.UpdateCustomer)
		customerRouter.DELETE("/delete", api.DeleteCustomer)
		customerRouter.POST("/getbyid", api.GetCustomerById)
		customerRouter.POST("/add", api.AddCustomer)
	}
	customerBusinessRouter := Router.Group("/cus").Use(middleware.JWTAuth())
	{
		customerBusinessRouter.POST("/upload", api.UploadCusHeaderImg)
		customerBusinessRouter.POST("/cart", api.GetCartList)
		customerBusinessRouter.POST("/addcart", api.AddCart)
		customerBusinessRouter.POST("/reduceCart", api.ReduceCart)
		customerBusinessRouter.POST("/delcart", api.DelCart)
		customerBusinessRouter.POST("/checkStatus", api.CheckStatus)
		customerBusinessRouter.POST("/getAddress", api.GetAddressListByUserId)
	}
}
