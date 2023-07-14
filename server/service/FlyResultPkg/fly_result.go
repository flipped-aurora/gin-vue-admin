package FlyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg"
	FlyResultPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type FlyResultService struct {
}

// CreateFlyResult 创建FlyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (FlyRtService *FlyResultService) CreateFlyResult(FlyRt *FlyResultPkg.FlyResult) (err error) {
	err = global.GVA_DB.Create(FlyRt).Error
	return err
}

// DeleteFlyResult 删除FlyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (FlyRtService *FlyResultService) DeleteFlyResult(FlyRt FlyResultPkg.FlyResult) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&FlyResultPkg.FlyResult{}).Where("id = ?", FlyRt.ID).Update("deleted_by", FlyRt.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&FlyRt).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteFlyResultByIds 批量删除FlyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (FlyRtService *FlyResultService) DeleteFlyResultByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&FlyResultPkg.FlyResult{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&FlyResultPkg.FlyResult{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateFlyResult 更新FlyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (FlyRtService *FlyResultService) UpdateFlyResult(FlyRt FlyResultPkg.FlyResult) (err error) {
	err = global.GVA_DB.Save(&FlyRt).Error
	return err
}

// GetFlyResult 根据id获取FlyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (FlyRtService *FlyResultService) GetFlyResult(id uint) (FlyRt FlyResultPkg.FlyResult, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FlyRt).Error
	return
}

// GetFlyResultInfoList 分页获取FlyResult记录
// Author [piexlmax](https://github.com/piexlmax)
func (FlyRtService *FlyResultService) GetFlyResultInfoList(info FlyResultPkgReq.FlyResultSearch) (list []FlyResultPkg.FlyResult, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&FlyResultPkg.FlyResult{})
	var FlyRts []FlyResultPkg.FlyResult
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&FlyRts).Error
	return FlyRts, total, err
}
