package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"gorm.io/gorm"
)

type CroppingRecordService struct {
}

// CreateCroppingRecord 创建CroppingRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (croppingRecordService *CroppingRecordService) CreateCroppingRecord(croppingRecord *clothing.CroppingRecord) (err error) {
	m := make(map[string]uint, 0)
	for _, list := range croppingRecord.SizeList {
		croppingRecord.Quantity += list.Quantity
		if _, ok := m[list.Size]; !ok {
			m[list.Size] = list.Quantity
		} else {
			m[list.Size] += list.Quantity
		}
	}
	p := make([]clothing.Process, 0) // 工序
	global.GVA_DB.Where("style_id = ?", croppingRecord.StyleID).Find(&p)
	i := make([]clothing.Inventory, 0)       // 库存
	sizeList := make([]clothing.SizeList, 0) // 尺码
	for s, u := range m {
		sizeList = append(sizeList, clothing.SizeList{
			Size:     s,
			Quantity: u,
			Margin:   u,
		})
		for _, process := range p {
			i = append(i, clothing.Inventory{
				ProcessID: process.ID,
				Size:      s,
				Total:     u,
				Margin:    u,
			})
		}
	}
	croppingRecord.SizeList = sizeList
	croppingRecord.Inventory = i
	croppingRecord.Usage = croppingRecord.Length / float64(croppingRecord.Quantity)
	err = global.GVA_DB.Create(&croppingRecord).Error
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
	global.GVA_DB.Where("cropping_record_id = ?", croppingRecord.ID).Unscoped().Delete(&clothing.Inventory{})
	global.GVA_DB.Where("cropping_record_id = ?", croppingRecord.ID).Unscoped().Delete(&clothing.SizeList{})
	jobs := make([]clothing.Job, 0)
	global.GVA_DB.Preload("Team").Where("cropping_id = ?", croppingRecord.ID).Find(&jobs)
	for _, job := range jobs {
		if job.Step == enum.CroppingComplete {
			var wallet clothing.UserWallet
			global.GVA_DB.Where("user_id = ? and company_id = ?", job.UserID, job.Team.CompanyID).FirstOrCreate(&wallet)
			err = global.GVA_DB.Model(&wallet).
				Updates(map[string]interface{}{
					"wages":         gorm.Expr("wages - ?", job.RealIncome),
					"pending_wages": gorm.Expr("pending_wages - ?", job.RealIncome),
				}).Error
		}
	}
	m := make(map[string]uint, 0)
	for _, list := range croppingRecord.SizeList {
		croppingRecord.Quantity += list.Quantity
		if _, ok := m[list.Size]; !ok {
			m[list.Size] = list.Quantity
		} else {
			m[list.Size] += list.Quantity
		}
	}
	p := make([]clothing.Process, 0) // 工序
	global.GVA_DB.Where("style_id = ?", croppingRecord.StyleID).Find(&p)
	i := make([]clothing.Inventory, 0)       // 库存
	sizeList := make([]clothing.SizeList, 0) // 尺码
	for s, u := range m {
		sizeList = append(sizeList, clothing.SizeList{
			Size:     s,
			Quantity: u,
			Margin:   u,
		})
		for _, process := range p {
			i = append(i, clothing.Inventory{
				ProcessID: process.ID,
				Size:      s,
				Total:     u,
				Margin:    u,
			})
		}
	}
	croppingRecord.SizeList = sizeList
	croppingRecord.Inventory = i
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
func (croppingRecordService *CroppingRecordService) GetCroppingRecordInfoList(info clothingReq.CroppingRecordSearch) (list []clothing.CroppingRecord, total, quantity, margin int64, err error) {
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
	ids := make([]uint, 0)
	db2 := db
	err = db2.Debug().Pluck("id", &ids).Error
	if err != nil {
		return
	}
	type sumRes struct {
		Quantity int64 `json:"quantity"`
		Margin   int64 `json:"margin"`
	}
	var s sumRes
	err = global.GVA_DB.Model(&clothing.SizeList{}).Where("cropping_record_id IN ?", ids).Select("sum(quantity) as quantity,sum(margin) as margin").Scan(&s).Error
	if err != nil {
		return
	}
	err = db.Preload("Style").Preload("SizeList").Limit(limit).Offset(offset).Find(&croppingRecords).Error
	return croppingRecords, total, s.Quantity, s.Margin, err
}
