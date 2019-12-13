package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AuthorityRouter.POST("createAuthority", api.CreateAuthority)   //创建角色
		AuthorityRouter.POST("deleteAuthority", api.DeleteAuthority)   //删除角色
		AuthorityRouter.POST("getAuthorityList", api.GetAuthorityList) //获取角色列表
	}
}
