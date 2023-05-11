package reservation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/reservation"
	reservationReq "github.com/flipped-aurora/gin-vue-admin/server/model/reservation/request"
	"gorm.io/gorm"
)

type EqtInfoService struct {
}

// CreateEqtInfo 创建EqtInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (eqtInfoService *EqtInfoService) CreateEqtInfo(eqtInfo *reservation.EqtInfo) (err error) {
	err = global.GVA_DB.Create(eqtInfo).Error
	return err
}

// DeleteEqtInfo 删除EqtInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (eqtInfoService *EqtInfoService) DeleteEqtInfo(eqtInfo reservation.EqtInfo) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&reservation.EqtInfo{}).Where("id = ?", eqtInfo.ID).Update("deleted_by", eqtInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&eqtInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteEqtInfoByIds 批量删除EqtInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (eqtInfoService *EqtInfoService) DeleteEqtInfoByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&reservation.EqtInfo{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&reservation.EqtInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateEqtInfo 更新EqtInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (eqtInfoService *EqtInfoService) UpdateEqtInfo(eqtInfo reservation.EqtInfo) (err error) {
	err = global.GVA_DB.Save(&eqtInfo).Error
	return err
}

// GetEqtInfo 根据id获取EqtInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (eqtInfoService *EqtInfoService) GetEqtInfo(id uint) (eqtInfo reservation.EqtInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eqtInfo).Error
	return
}

// GetEqtInfoInfoList 分页获取EqtInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (eqtInfoService *EqtInfoService) GetEqtInfoInfoList(info reservationReq.EqtInfoSearch) (list []reservation.EqtInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&reservation.EqtInfo{})
	var eqtInfos []reservation.EqtInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.EqtClass != nil {
		db = db.Where("eqt_class = ?", info.EqtClass)
	}
	if info.EqtCardNo != "" {
		db = db.Where("eqt_card_no = ?", info.EqtCardNo)
	}
	if info.EqtSN != "" {
		db = db.Where("eqt_sn = ?", info.EqtSN)
	}
	if info.EqtName != "" {
		db = db.Where("eqt_name = ?", info.EqtName)
	}
	if info.EqtModel != "" {
		db = db.Where("eqt_model = ?", info.EqtModel)
	}
	if info.EqtStatus != nil {
		db = db.Where("eqt_status = ?", info.EqtStatus)
	}
	if info.EqtStockStatus != nil {
		db = db.Where("eqt_stock_status = ?", info.EqtStockStatus)
	}
	if info.CurrentDept != nil {
		db = db.Where("current_dept = ?", info.CurrentDept)
	}
	if info.StartServiceTime != nil && info.EndServiceTime != nil {
		db = db.Where("service_time BETWEEN ? AND ? ", info.StartServiceTime, info.EndServiceTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&eqtInfos).Error
	return eqtInfos, total, err
}
