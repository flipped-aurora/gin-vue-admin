package home

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type HomeRouter struct {
}

func (h *HomeRouter) InitHomeRouter(Router *gin.RouterGroup) {
	var homeApi = v1.ApiGroupApp.HomeApiGroup.HomeApi
	Router.GET("/", homeApi.Home)
	Router.GET("/search", homeApi.Search)
	Router.POST("/message", homeApi.PostMessage)
	Router.GET("/captcha", homeApi.GetCaptcha)
}
