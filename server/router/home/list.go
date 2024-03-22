package home

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ListRouter struct {
}

func (h *HomeRouter) InitListRouter(Router *gin.RouterGroup) {
	var listApi = v1.ApiGroupApp.HomeApiGroup.ListApi
	Router.GET("/list/:catid", listApi.List)
}
