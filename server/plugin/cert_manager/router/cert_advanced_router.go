package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/api"
	"github.com/gin-gonic/gin"
)

type certAdvancedRouter struct{}

var CertAdvancedRouter = new(certAdvancedRouter)

func (r *certAdvancedRouter) InitCertAdvancedRouter(Router *gin.RouterGroup) {
	advancedGroup := Router.Group("cert/advanced")
	{
		advancedGroup.POST("discoverSubdomains", api.CertAdvancedApi.DiscoverSubdomains)
		advancedGroup.POST("batchDiscoverSubdomains", api.CertAdvancedApi.BatchDiscoverSubdomains)
		advancedGroup.GET("domainTree", api.CertAdvancedApi.GetDomainTree)
		advancedGroup.POST("reProbeDomainTree", api.CertAdvancedApi.ReProbeDomainTree)
		advancedGroup.POST("ignoreDomain", api.CertAdvancedApi.IgnoreDomain)
		advancedGroup.POST("batchReprobe", api.CertAdvancedApi.BatchReprobe)
		advancedGroup.POST("batchIgnore", api.CertAdvancedApi.BatchIgnore)
		advancedGroup.GET("domainAssetList", api.CertAdvancedApi.GetDomainAssetList)
		advancedGroup.GET("exportSubdomainReport", api.CertAdvancedApi.ExportSubdomainReport)
	}
}
