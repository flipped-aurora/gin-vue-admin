package service

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"go.uber.org/zap"
)

type certAssociation struct{}

func (s *certAssociation) AssociateCertificateWithDomain(certInfo *CertInfo, domain string) error {
	var existingEntity model.CertificateEntity
	err := global.GVA_DB.Where("fingerprint = ?", certInfo.Fingerprint).First(&existingEntity).Error
	
	if err != nil {
		newEntity := &model.CertificateEntity{
			Fingerprint: certInfo.Fingerprint,
			Brand:       certInfo.Brand,
			StartAt:     certInfo.StartAt,
			ExpireAt:    certInfo.ExpireAt,
			DaysRemaining: certInfo.DaysRemaining,
			TlsVersion:  certInfo.TlsVersion,
			IsWildcard:  certInfo.IsWildcard,
			SANs:        strings.Join(certInfo.SANs, ","),
		}
		if err := global.GVA_DB.Create(newEntity).Error; err != nil {
			global.GVA_LOG.Error("创建证书实体失败", zap.Error(err))
			return err
		}
		return global.GVA_DB.Model(&model.DomainAsset{}).Where("domain = ?", domain).Update("cert_id", newEntity.ID).Error
	}
	
	return global.GVA_DB.Model(&model.DomainAsset{}).Where("domain = ?", domain).Update("cert_id", existingEntity.ID).Error
}

func (s *certAssociation) MatchWildcardCertificate(rootDomain string, wildcardCert *model.CertificateEntity) error {
	var subdomains []model.DomainAsset
	err := global.GVA_DB.Where("root_domain = ? AND is_ignored = ?", rootDomain, false).Find(&subdomains).Error
	if err != nil {
		return err
	}
	
	var matchedDomains []string
	sans := strings.Split(wildcardCert.SANs, ",")
	for _, subdomain := range subdomains {
		if s.isWildcardMatch(sans, subdomain.Domain) {
			matchedDomains = append(matchedDomains, subdomain.Domain)
		}
	}
	
	if len(matchedDomains) > 0 {
		if err := global.GVA_DB.Model(&model.DomainAsset{}).
			Where("domain IN ?", matchedDomains).
			Updates(map[string]interface{}{
				"cert_id":   wildcardCert.ID,
				"cert_type": "wildcard",
			}).Error; err != nil {
			global.GVA_LOG.Error("批量关联泛域名证书失败", zap.Error(err))
			return err
		}
		global.GVA_LOG.Info("泛域名证书自动关联", zap.Int("count", len(matchedDomains)), zap.String("root", rootDomain))
	}
	
	return nil
}

func (s *certAssociation) isWildcardMatch(sans []string, domain string) bool {
	for _, san := range sans {
		if strings.HasPrefix(san, "*.") {
			wildcardPattern := strings.TrimPrefix(san, "*.")
			if strings.HasSuffix(domain, wildcardPattern) {
				return true
			}
		}
	}
	return false
}

func (s *certAssociation) ResolveCertificateOverlap(domain string) (*model.CertificateEntity, error) {
	var assets []model.DomainAsset
	err := global.GVA_DB.Where("domain = ?", domain).Find(&assets).Error
	if err != nil {
		return nil, err
	}
	
	if len(assets) == 0 {
		return nil, nil
	}
	
	var certEntities []model.CertificateEntity
	var certIDs []uint
	for _, asset := range assets {
		if asset.CertID != nil {
			certIDs = append(certIDs, *asset.CertID)
		}
	}
	
	if len(certIDs) == 0 {
		return nil, nil
	}
	
	err = global.GVA_DB.Where("id IN ?", certIDs).Find(&certEntities).Error
	if err != nil {
		return nil, err
	}
	
	if len(certEntities) == 1 {
		return &certEntities[0], nil
	}
	
	var shortestCert *model.CertificateEntity
	minDays := -1
	for i := range certEntities {
		if minDays == -1 || certEntities[i].DaysRemaining < minDays {
			minDays = certEntities[i].DaysRemaining
			shortestCert = &certEntities[i]
		}
	}
	
	global.GVA_LOG.Info("检测到证书重叠，选择剩余天数最短的", zap.String("domain", domain), zap.Int("days", minDays))
	return shortestCert, nil
}
