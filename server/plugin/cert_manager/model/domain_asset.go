package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DomainAsset struct {
	global.GVA_MODEL
	Domain       string         `json:"domain" gorm:"column:domain;comment:监控域名;not null;index:idx_domain,priority:1"`
	ParentID     *uint          `json:"parentId" gorm:"column:parent_id;comment:父域名ID;index:idx_parent_id"`
	RootDomain   string         `json:"rootDomain" gorm:"column:root_domain;comment:根域名;index:idx_root_domain"`
	Category     string         `json:"category" gorm:"column:category;comment:项目分类"`
	IsIgnored    bool           `json:"isIgnored" gorm:"column:is_ignored;comment:是否忽略;default:false"`
	CertID       *uint          `json:"certId" gorm:"column:cert_id;comment:关联证书ID;index:idx_cert_id"`
	CertType     string         `json:"certType" gorm:"column:cert_type;comment:证书类型(wildcard/single)"`
	DomainStatus int            `json:"domainStatus" gorm:"column:domain_status;comment:域名状态(1:正常, 3:异常);default:1"`
	CertStatus   int            `json:"certStatus" gorm:"column:cert_status;comment:证书状态(1:正常, 3:异常);default:1"`
	IP           string         `json:"ip" gorm:"column:ip;comment:解析IP"`
	Source       string         `json:"source" gorm:"column:source;comment:发现来源"`
	LastProbedAt *time.Time     `json:"lastProbedAt" gorm:"column:last_probed_at;comment:最后探测时间"`
	Children     []*DomainAsset `json:"children" gorm:"-"`
}

func (DomainAsset) TableName() string {
	return "gva_cert_domain_assets"
}
