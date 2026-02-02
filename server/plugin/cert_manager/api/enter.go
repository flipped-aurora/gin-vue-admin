package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/service"

var (
	ApiGroupApp            = new(apiGroup)
	serviceCertCertificate = service.ServiceGroupApp.CertCertificate
	serviceCertCategory    = service.ServiceGroupApp.CertCategory
)

type apiGroup struct {
	CertCertificate certCertificate
	CertCategory    certCategory
	CertAdvanced    certAdvanced
}
