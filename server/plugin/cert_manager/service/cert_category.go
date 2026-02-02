package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/model/request"
)

type certCategory struct{}

func (s *certCategory) CreateCertCategory(cc model.CertCategory) error {
	return global.GVA_DB.Create(&cc).Error
}

func (s *certCategory) DeleteCertCategory(cc model.CertCategory) error {
	return global.GVA_DB.Delete(&cc).Error
}

func (s *certCategory) UpdateCertCategory(cc model.CertCategory) error {
	return global.GVA_DB.Save(&cc).Error
}

func (s *certCategory) GetCertCategory(id uint) (cc model.CertCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cc).Error
	return
}

func (s *certCategory) GetCertCategoryList(info request.CertCategorySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.CertCategory{})
	var categories []model.CertCategory
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&categories).Error
	return categories, total, err
}

func (s *certCategory) BatchUpdateCertCategory(info request.BatchUpdateCategory) error {
	return global.GVA_DB.Model(&model.CertCertificate{}).Where("id IN ?", info.IDs).Update("category", info.Category).Error
}
