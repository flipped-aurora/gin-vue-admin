package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
)

type DomainAssetSearch struct {
	model.DomainAsset
	Domain       string `json:"domain" form:"domain" gorm:"column:domain;comment:监控域名"`
	RootDomain   string `json:"rootDomain" form:"rootDomain"`
	Category     string `json:"category" form:"category"`
	IsIgnored    *bool  `json:"isIgnored" form:"isIgnored"`
	DomainStatus *int   `json:"domainStatus" form:"domainStatus"`
	CertStatus   *int   `json:"certStatus" form:"certStatus"`
	request.PageInfo
}
