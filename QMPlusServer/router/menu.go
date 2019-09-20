package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitMenuRouter(Router *gin.Engine) {
	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth())
	{
		MenuRouter.POST("getMenu", api.GetMenu)
		MenuRouter.POST("getMenuList", api.GetMenuList)
		MenuRouter.POST("addBaseMenu", api.AddBaseMenu)
		MenuRouter.POST("getBaseMenuTree", api.GetBaseMenuTree)
		MenuRouter.POST("addMenuAuthority", api.AddMenuAuthority)
		MenuRouter.POST("getMenuAuthority", api.GetMenuAuthority)
		MenuRouter.POST("deleteBaseMenu",api.DeleteBaseMenu)
	}
}
