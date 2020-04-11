package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysAddressRouter(Router *gin.RouterGroup) {
	sysAddressRouter := Router.Group("/sysAddress").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		sysAddressRouter.POST("province", api.GetProvincesList)
		sysAddressRouter.POST("city", api.GetCityListByProvince)
	}
}
