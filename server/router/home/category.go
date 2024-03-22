package home

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (h *HomeRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	var categoryApi = v1.ApiGroupApp.HomeApiGroup.CategoryApi
	Router.GET("/category/:id", categoryApi.Category)
}
