package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitAuthorityRouter(Router *gin.Engine) {
	AuthorityRouter := Router.Group("authority").Use(middleware.JWTAuth())
	{
		AuthorityRouter.POST("createAuthority", api.CreateAuthority) //创建角色
		AuthorityRouter.POST("deleteAuthority", api.DeleteAuthority) //删除角色
		AuthorityRouter.POST("getAuthorityList",api.GetAuthorityList) //获取角色列表
	}
}
