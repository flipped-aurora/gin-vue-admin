package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WebconfigRouter struct {
}

func (w *WebconfigRouter) InitWebConfigRouter(Router *gin.RouterGroup) {
	webconfigRouter := Router.Use(middleware.OperationRecord())
	var webconfigApi = v1.ApiGroupApp.WebcmsApiGroup.WebconfigApi
	{
		webconfigRouter.POST("/createWebconfig", webconfigApi.CreateWebconfig)
		webconfigRouter.GET("/getWebconfig", webconfigApi.GetWebconfig)
		webconfigRouter.DELETE("/delWebconfig", webconfigApi.DeleteWebconfig)
		webconfigRouter.POST("/setWebconfig", webconfigApi.SetWebconfig)
		webconfigRouter.POST("/changeWebconfig", webconfigApi.ChangeWebconfig)
	}
}
