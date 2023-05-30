package clothing

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type StyleService struct {
}

// CreateStyle 创建Style记录
// Author [piexlmax](https://github.com/piexlmax)
func (styleService *StyleService) CreateStyle(style *clothing.Style) (err error) {
	var s clothing.Style
	if err := global.GVA_DB.Where("name = ? and company_id = ?", style.Name, style.CompanyID).First(&s).Error; err == nil {
		return errors.New("款式已存在")
	}
	err = global.GVA_DB.Create(style).Error
	return err
}

// DeleteStyle 删除Style记录
// Author [piexlmax](https://github.com/piexlmax)
func (styleService *StyleService) DeleteStyle(style clothing.Style) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Style{}).Where("id = ?", style.ID).Update("deleted_by", style.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&style).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteStyleByIds 批量删除Style记录
// Author [piexlmax](https://github.com/piexlmax)
func (styleService *StyleService) DeleteStyleByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Style{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Style{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateStyle 更新Style记录
// Author [piexlmax](https://github.com/piexlmax)
func (styleService *StyleService) UpdateStyle(style clothing.Style) (err error) {
	err = global.GVA_DB.Model(&style).Updates(map[string]interface{}{
		"name":  style.Name,
		"price": style.Price,
	}).Error
	return err
}

// GetStyle 根据id获取Style记录
// Author [piexlmax](https://github.com/piexlmax)
func (styleService *StyleService) GetStyle(id uint) (style clothing.Style, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&style).Error
	return
}

// GetStyleInfoList 分页获取Style记录
// Author [piexlmax](https://github.com/piexlmax)
func (styleService *StyleService) GetStyleInfoList(info clothingReq.StyleSearch) (list []clothing.Style, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Style{})
	var styles []clothing.Style
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CompanyID != 0 {
		db = db.Where("company_id = ?", info.CompanyID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&styles).Error
	return styles, total, err
}
