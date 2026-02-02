package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
)

type CertCertificateSearch struct {
	model.CertCertificate
	Domain       string `json:"domain" form:"domain" gorm:"column:domain;comment:监控域名"`
	Category     string `json:"category" form:"category"`
	DomainStatus *int   `json:"domainStatus" form:"domainStatus"`
	CertStatus   *int   `json:"certStatus" form:"certStatus"`
	request.PageInfo
}
