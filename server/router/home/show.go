package home

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShowRouter struct {
}

func (h *HomeRouter) InitShowRouter(Router *gin.RouterGroup) {
	var showApi = v1.ApiGroupApp.HomeApiGroup.ShowApi
	Router.GET("/show/:catid/:id", showApi.Show)
}
