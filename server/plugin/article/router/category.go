package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Category = new(category)

type category struct{}

func (r *category) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("article").Use(middleware.OperationRecord())
		group.POST("createCategory", apiCategory.CreateCategory)
		group.DELETE("deleteCategory", apiCategory.DeleteCategory)
		group.PUT("updateCategory", apiCategory.UpdateCategory)
	}
	{
		group := private.Group("article")
		group.GET("getCategoryList", apiCategory.GetCategoryList)
		group.GET("getAllCategories", apiCategory.GetAllCategories)
	}
}
