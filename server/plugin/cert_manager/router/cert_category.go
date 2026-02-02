package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type certCategory struct{}

func (r *certCategory) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	group := private.Group("certCategory").Use(middleware.OperationRecord())
	{
		group.POST("createCertCategory", apiCertCategory.CreateCertCategory)
		group.DELETE("deleteCertCategory", apiCertCategory.DeleteCertCategory)
		group.PUT("updateCertCategory", apiCertCategory.UpdateCertCategory)
		group.GET("getCertCategoryList", apiCertCategory.GetCertCategoryList)
		group.POST("batchUpdateCertCategory", apiCertCategory.BatchUpdateCertCategory)
	}
}
