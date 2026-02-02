package service

import (
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model/request"
	"go.uber.org/zap"
)

type certCertificate struct {
	sem chan struct{}
}

func (s *certCertificate) CreateCertCertificate(cert *model.CertCertificate) (err error) {
	err = global.GVA_DB.Create(cert).Error
	return err
}

func (s *certCertificate) DeleteCertCertificate(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.CertCertificate{}, "id = ?", ID).Error
	return err
}

func (s *certCertificate) DeleteCertCertificateByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.CertCertificate{}, "id in ?", IDs).Error
	return err
}

func (s *certCertificate) UpdateCertCertificate(cert model.CertCertificate) (err error) {
	err = global.GVA_DB.Model(&model.CertCertificate{}).Where("id = ?", cert.ID).Updates(&cert).Error
	return err
}

func (s *certCertificate) GetCertCertificate(ID string) (cert model.CertCertificate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cert).Error
	return
}

func (s *certCertificate) GetCertCertificateList(info request.CertCertificateSearch) (list []model.CertCertificate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.CertCertificate{})
	var certs []model.CertCertificate

	if info.Domain != "" {
		db = db.Where("domain LIKE ?", "%"+info.Domain+"%")
	}
	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}
	if info.DomainStatus != nil {
		db = db.Where("domain_status = ?", *info.DomainStatus)
	}
	if info.CertStatus != nil {
		db = db.Where("cert_status = ?", *info.CertStatus)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("days_remaining ASC").Find(&certs).Error
	return certs, total, err
}

func (s *certCertificate) getSem() chan struct{} {
	if s.sem == nil {
		concurrency := 5
		// 未来可以从配置加载：concurrency = global.GVA_CONFIG.CertManager.Concurrency
		s.sem = make(chan struct{}, concurrency)
	}
	return s.sem
}

func (s *certCertificate) ProbeAndUpdateCertificate(domain string) error {
	s.getSem() <- struct{}{}
	defer func() { <-s.sem }()

	// 1. 证书探测
	certInfo, certErr := s.ProbeCertificate(domain)
	// 2. 域名 WHOIS 探测
	domainInfo, domainErr := s.ProbeDomainWhois(domain)

	var existingCert model.CertCertificate
	result := global.GVA_DB.Where("domain = ?", domain).First(&existingCert)

	// 3. 严重错误判定：两者均彻底失败
	if certErr != nil && domainErr != nil {
		if result.Error == nil {
			existingCert.DomainStatus = 3 // 异常
			existingCert.CertStatus = 3   // 异常
			existingCert.ErrorLog = "证书与域名探测均失败，请查看系统日志获取详情"
			global.GVA_DB.Save(&existingCert)
		}
		global.GVA_LOG.Error("证书与域名探测全面失败", zap.String("domain", domain), zap.Error(certErr), zap.NamedError("domainErr", domainErr))
		return certErr
	}

	// 4. 严格状态判定逻辑
	domainStatus := 1
	certStatus := 1

	evalStatus := func(expireAt *time.Time) int {
		if expireAt == nil || expireAt.IsZero() {
			return 3 // 无数据视为异常
		}
		days := int(time.Until(*expireAt).Hours() / 24)
		if days <= 0 {
			return 3 // 异常（已过期）
		}
		return 1 // 正常
	}

	if certInfo != nil {
		certStatus = evalStatus(certInfo.ExpireAt)
		if err := s.associateCertificateWithDomain(certInfo, domain); err != nil {
			global.GVA_LOG.Error("证书关联失败", zap.Error(err))
		}
	} else if certErr != nil {
		certStatus = 3
	}

	if domainInfo != nil {
		domainStatus = evalStatus(domainInfo.ExpireAt)
	} else if domainErr != nil {
		domainStatus = 3
	}

	if result.Error != nil {
		// 新增记录
		newCert := &model.CertCertificate{
			Domain:       domain,
			DomainStatus: domainStatus,
			CertStatus:   certStatus,
		}
		if certInfo != nil {
			newCert.Brand = certInfo.Brand
			newCert.StartAt = certInfo.StartAt
			newCert.ExpireAt = certInfo.ExpireAt
			newCert.DaysRemaining = certInfo.DaysRemaining
			newCert.TlsVersion = certInfo.TlsVersion
		}
		if domainInfo != nil {
			newCert.DomainStartAt = domainInfo.StartAt
			newCert.DomainExpireAt = domainInfo.ExpireAt
		}
		return global.GVA_DB.Create(newCert).Error
	} else {
		// 更新记录
		existingCert.DomainStatus = domainStatus
		existingCert.CertStatus = certStatus
		existingCert.ErrorLog = ""
		if certInfo != nil {
			existingCert.Brand = certInfo.Brand
			existingCert.StartAt = certInfo.StartAt
			existingCert.ExpireAt = certInfo.ExpireAt
			existingCert.DaysRemaining = certInfo.DaysRemaining
			existingCert.TlsVersion = certInfo.TlsVersion
		}
		if domainInfo != nil {
			existingCert.DomainStartAt = domainInfo.StartAt
			existingCert.DomainExpireAt = domainInfo.ExpireAt
		}
		return global.GVA_DB.Save(&existingCert).Error
	}
}

func (s *certCertificate) UpdateAllCertificates() error {
	var certs []model.CertCertificate
	if err := global.GVA_DB.Find(&certs).Error; err != nil {
		return err
	}

	for _, cert := range certs {
		if err := s.ProbeAndUpdateCertificate(cert.Domain); err != nil {
			global.GVA_DB.Model(&cert).Update("error_log", err.Error())
		}
	}
	return nil
}

func (s *certCertificate) associateCertificateWithDomain(certInfo *CertInfo, domain string) error {
	if certInfo == nil {
		return nil
	}

	fingerprint := certInfo.Fingerprint
	if fingerprint == "" {
		return nil
	}

	sansStr := strings.Join(certInfo.SANs, ",")

	var certEntity model.CertificateEntity
	err := global.GVA_DB.Where("fingerprint = ?", fingerprint).First(&certEntity).Error

	if err != nil {
		certEntity = model.CertificateEntity{
			Fingerprint:   fingerprint,
			IsWildcard:    certInfo.IsWildcard,
			SANs:          sansStr,
			Brand:         certInfo.Brand,
			StartAt:       certInfo.StartAt,
			ExpireAt:      certInfo.ExpireAt,
			DaysRemaining: certInfo.DaysRemaining,
			TlsVersion:    certInfo.TlsVersion,
		}
		if err := global.GVA_DB.Create(&certEntity).Error; err != nil {
			return err
		}
	}

	rootDomain := extractRootDomain(domain)

	var domainAsset model.DomainAsset
	err = global.GVA_DB.Where("domain = ?", domain).First(&domainAsset).Error

	if err != nil {
		domainAsset = model.DomainAsset{
			Domain:       domain,
			RootDomain:   rootDomain,
			CertID:       &certEntity.ID,
			CertType:     getCertType(certInfo.IsWildcard),
			DomainStatus: 1,
			CertStatus:   1,
		}
		return global.GVA_DB.Create(&domainAsset).Error
	} else {
		domainAsset.CertID = &certEntity.ID
		domainAsset.CertType = getCertType(certInfo.IsWildcard)
		return global.GVA_DB.Save(&domainAsset).Error
	}
}

func extractRootDomain(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) >= 2 {
		return strings.Join(parts[len(parts)-2:], ".")
	}
	return domain
}

func getCertType(isWildcard bool) string {
	if isWildcard {
		return "泛域名证书"
	}
	return "独立证书"
}
