package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model/request"
)

type domainAssetService struct{}

func (s *domainAssetService) GetDomainAssetList(info request.DomainAssetSearch) (list []model.DomainAsset, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DomainAsset{})

	if info.Domain != "" {
		db = db.Where("domain LIKE ?", "%"+info.Domain+"%")
	}
	if info.RootDomain != "" {
		db = db.Where("root_domain = ?", info.RootDomain)
	}
	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}
	if info.IsIgnored != nil {
		db = db.Where("is_ignored = ?", *info.IsIgnored)
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

	err = db.Order("created_at DESC").Find(&list).Error
	return
}
