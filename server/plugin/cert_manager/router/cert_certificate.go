package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type certCertificate struct{}

func (r *certCertificate) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("certCertificate").Use(middleware.OperationRecord())
		group.POST("createCertCertificate", apiCertCertificate.CreateCertCertificate)
		group.DELETE("deleteCertCertificate", apiCertCertificate.DeleteCertCertificate)
		group.DELETE("deleteCertCertificateByIds", apiCertCertificate.DeleteCertCertificateByIds)
		group.PUT("updateCertCertificate", apiCertCertificate.UpdateCertCertificate)
		group.POST("probeCertificate", apiCertCertificate.ProbeCertificate)
		group.POST("updateAllCertificates", apiCertCertificate.UpdateAllCertificates)
	}
	{
		group := private.Group("certCertificate")
		group.GET("findCertCertificate", apiCertCertificate.FindCertCertificate)
		group.GET("getCertCertificateList", apiCertCertificate.GetCertCertificateList)
	}
}
