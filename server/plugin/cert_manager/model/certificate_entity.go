package model

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type CertificateEntity struct {
	global.GVA_MODEL
	Fingerprint string     `json:"fingerprint" gorm:"column:fingerprint;comment:证书指纹(SHA-256);uniqueIndex:idx_fingerprint;size:64"`
	Brand       string     `json:"brand" gorm:"column:brand;comment:证书品牌/颁发者"`
	StartAt     *time.Time `json:"startAt" gorm:"column:start_at;comment:证书签发时间"`
	ExpireAt    *time.Time `json:"expireAt" gorm:"column:expire_at;comment:证书到期时间"`
	DaysRemaining int        `json:"daysRemaining" gorm:"column:days_remaining;comment:证书剩余天数"`
	TlsVersion  string     `json:"tlsVersion" gorm:"column:tls_version;comment:TLS版本"`
	IsWildcard  bool       `json:"isWildcard" gorm:"column:is_wildcard;comment:是否泛域名证书;default:false"`
	SANs        string     `json:"sans" gorm:"column:sans;type:text;comment:Subject Alternative Names"`
}

func (CertificateEntity) TableName() string {
	return "gva_cert_entities"
}

func (c *CertificateEntity) CalculateFingerprint(certData []byte) string {
	hash := sha256.Sum256(certData)
	return hex.EncodeToString(hash[:])
}
