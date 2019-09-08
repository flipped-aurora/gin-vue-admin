package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitAuthorityRouter(Router *gin.Engine) {
	AuthorityRouter := Router.Group("authority").Use(middleware.JWTAuth())
	{
		AuthorityRouter.POST("createAuthority", api.CreateAuthority)
		AuthorityRouter.POST("deleteAuthority", api.DeleteAuthority)
	}
}
