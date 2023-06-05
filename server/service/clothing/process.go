package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type ProcessService struct {
}

// CreateProcess 创建Process记录
// Author [piexlmax](https://github.com/piexlmax)
func (processService *ProcessService) CreateProcess(process *clothing.Process) (err error) {
	err = global.GVA_DB.Create(process).Error
	return err
}

// DeleteProcess 删除Process记录
// Author [piexlmax](https://github.com/piexlmax)
func (processService *ProcessService) DeleteProcess(process clothing.Process) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Process{}).Where("id = ?", process.ID).Update("deleted_by", process.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&process).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteProcessByIds 批量删除Process记录
// Author [piexlmax](https://github.com/piexlmax)
func (processService *ProcessService) DeleteProcessByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Process{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Process{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateProcess 更新Process记录
// Author [piexlmax](https://github.com/piexlmax)
func (processService *ProcessService) UpdateProcess(process clothing.Process) (err error) {
	err = global.GVA_DB.Model(&process).Updates(map[string]interface{}{
		"style_id": process.StyleID,
		"name":     process.Name,
		"price":    process.Price,
		"percent":  process.Percent,
	}).Error
	return err
}

// GetProcess 根据id获取Process记录
// Author [piexlmax](https://github.com/piexlmax)
func (processService *ProcessService) GetProcess(id uint) (process clothing.Process, err error) {
	err = global.GVA_DB.Preload("Style").Where("id = ?", id).First(&process).Error
	return
}

// GetProcessInfoList 分页获取Process记录
// Author [piexlmax](https://github.com/piexlmax)
func (processService *ProcessService) GetProcessInfoList(info clothingReq.ProcessSearch) (list []clothing.Process, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Process{})
	var processs []clothing.Process
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StyleID != 0 {
		db = db.Where("style_id = ?", info.StyleID)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("Style").Limit(limit).Offset(offset).Find(&processs).Error
	return processs, total, err
}
