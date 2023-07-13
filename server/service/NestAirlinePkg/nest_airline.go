package NestAirlinePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg"
	NestAirlinePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type NestAirlineService struct {
}

// CreateNestAirline 创建NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) CreateNestAirline(NtAirline *NestAirlinePkg.NestAirline) (err error) {
	err = global.GVA_DB.Create(NtAirline).Error
	return err
}

// DeleteNestAirline 删除NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) DeleteNestAirline(NtAirline NestAirlinePkg.NestAirline) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestAirlinePkg.NestAirline{}).Where("id = ?", NtAirline.ID).Update("deleted_by", NtAirline.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&NtAirline).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteNestAirlineByIds 批量删除NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) DeleteNestAirlineByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestAirlinePkg.NestAirline{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&NestAirlinePkg.NestAirline{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateNestAirline 更新NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) UpdateNestAirline(NtAirline NestAirlinePkg.NestAirline) (err error) {
	err = global.GVA_DB.Save(&NtAirline).Error
	return err
}

// GetNestAirline 根据id获取NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) GetNestAirline(id uint) (NtAirline NestAirlinePkg.NestAirline, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&NtAirline).Error
	return
}

// GetNestAirlineInfoList 分页获取NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) GetNestAirlineInfoList(info NestAirlinePkgReq.NestAirlineSearch) (list []NestAirlinePkg.NestAirline, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&NestAirlinePkg.NestAirline{})
	var NtAirlines []NestAirlinePkg.NestAirline
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&NtAirlines).Error
	return NtAirlines, total, err
}
