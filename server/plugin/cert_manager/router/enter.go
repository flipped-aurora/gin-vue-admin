package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/api"

var (
	RouterGroupApp     = new(routerGroup)
	apiCertCertificate = api.ApiGroupApp.CertCertificate
	apiCertCategory    = api.ApiGroupApp.CertCategory
	apiCertAdvanced    = api.ApiGroupApp.CertAdvanced
)

type routerGroup struct {
	CertCertificate    certCertificate
	CertCategory       certCategory
	CertAdvancedRouter certAdvancedRouter
}
