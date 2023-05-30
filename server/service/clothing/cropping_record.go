package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type CroppingRecordService struct {
}

// CreateCroppingRecord 创建CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) CreateCroppingRecord(croppingRecord *clothing.CroppingRecord) (err error) {
	for _, list := range croppingRecord.SizeList {
		croppingRecord.Quantity += list.Quantity
	}
	croppingRecord.Usage = croppingRecord.Length / float64(croppingRecord.Quantity)
	err = global.GVA_DB.Create(croppingRecord).Error
	return err
}

// DeleteCroppingRecord 删除CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) DeleteCroppingRecord(croppingRecord clothing.CroppingRecord) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.CroppingRecord{}).Where("id = ?", croppingRecord.ID).Update("deleted_by", croppingRecord.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&croppingRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCroppingRecordByIds 批量删除CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) DeleteCroppingRecordByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.CroppingRecord{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.CroppingRecord{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCroppingRecord 更新CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) UpdateCroppingRecord(croppingRecord clothing.CroppingRecord) (err error) {
	global.GVA_DB.Where("cropping_record_id = ?", croppingRecord.ID).Unscoped().Delete(&clothing.SizeList{})
	for _, list := range croppingRecord.SizeList {
		croppingRecord.Quantity += list.Quantity
	}
	croppingRecord.Usage = croppingRecord.Length / float64(croppingRecord.Quantity)
	err = global.GVA_DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&croppingRecord).Error
	return err
}

// GetCroppingRecord 根据id获取CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) GetCroppingRecord(id uint) (croppingRecord clothing.CroppingRecord, err error) {
	err = global.GVA_DB.Preload("Style").Preload("SizeList").Where("id = ?", id).First(&croppingRecord).Error
	return
}

// GetCroppingRecordInfoList 分页获取CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) GetCroppingRecordInfoList(info clothingReq.CroppingRecordSearch) (list []clothing.CroppingRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.CroppingRecord{})
	var croppingRecords []clothing.CroppingRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CompanyID != 0 {
		db = db.Where("company_id = ?", info.CompanyID)
	}
	if info.StyleID != 0 {
		db = db.Where("style_id = ?", info.StyleID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("Style").Preload("SizeList").Limit(limit).Offset(offset).Find(&croppingRecords).Error
	return croppingRecords, total, err
}
