package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type CertCertificate struct {
	global.GVA_MODEL
	Domain         string     `json:"domain" gorm:"column:domain;comment:监控域名;not null;index:idx_domain,priority:1"`
	Brand          string     `json:"brand" gorm:"column:brand;comment:证书品牌/颁发者"`
	Category       string     `json:"category" gorm:"column:category;comment:项目分类"`
	StartAt        *time.Time `json:"startAt" gorm:"column:start_at;comment:证书签发时间"`
	ExpireAt       *time.Time `json:"expireAt" gorm:"column:expire_at;comment:证书到期时间"`
	DaysRemaining  int        `json:"-" gorm:"column:days_remaining;comment:证书剩余天数"`
	DomainStartAt  *time.Time `json:"domainStartAt" gorm:"column:domain_start_at;comment:域名注册时间"`
	DomainExpireAt *time.Time `json:"domainExpireAt" gorm:"column:domain_expire_at;comment:域名到期时间"`
	TlsVersion     string     `json:"tlsVersion" gorm:"column:tls_version;comment:TLS版本"`
	DomainStatus   int        `json:"domainStatus" gorm:"column:domain_status;comment:域名状态(1:正常, 3:异常);default:1"`
	CertStatus     int        `json:"certStatus" gorm:"column:cert_status;comment:证书状态(1:正常, 3:异常);default:1"`
	ErrorLog       string     `json:"errorLog" gorm:"column:error_log;type:text;comment:错误日志"`
	CertID         *uint      `json:"certId" gorm:"column:cert_id;comment:关联证书实体ID;index:idx_cert_id"`
	ParentID       *uint      `json:"parentId" gorm:"column:parent_id;comment:父域名ID;index:idx_parent_id"`
}

func (CertCertificate) TableName() string {
	return "gva_cert_certificates"
}
