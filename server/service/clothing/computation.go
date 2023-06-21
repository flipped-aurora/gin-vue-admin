package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
	"math"
)

type ComputationService struct {
}

// CreateComputation 创建Computation记录
// Author [piexlmax](https://github.com/piexlmax)
func (computationService *ComputationService) CreateComputation(computation *clothing.Computation) (err error) {
	err = global.GVA_DB.Create(computation).Error
	return err
}

// DeleteComputation 删除Computation记录
// Author [piexlmax](https://github.com/piexlmax)
func (computationService *ComputationService) DeleteComputation(computation clothing.Computation) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Computation{}).Where("id = ?", computation.ID).Update("deleted_by", computation.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&computation).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteComputationByIds 批量删除Computation记录
// Author [piexlmax](https://github.com/piexlmax)
func (computationService *ComputationService) DeleteComputationByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Computation{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Computation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateComputation 更新Computation记录
// Author [piexlmax](https://github.com/piexlmax)
func (computationService *ComputationService) UpdateComputation(computation clothing.Computation) (err error) {
	err = global.GVA_DB.Save(&computation).Error
	return err
}

// GetComputation 根据id获取Computation记录
// Author [piexlmax](https://github.com/piexlmax)
func (computationService *ComputationService) GetComputation(id uint) (computation clothing.Computation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&computation).Error
	return
}

// GetComputationInfoList 分页获取Computation记录
// Author [piexlmax](https://github.com/piexlmax)
func (computationService *ComputationService) GetComputationInfoList(info clothingReq.ComputationSearch) (list []clothing.Computation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Computation{})
	var computations []clothing.Computation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&computations).Error
	return computations, total, err
}

func (computationService *ComputationService) DoComputation(computation *clothing.Computation) (err error) {
	computation.BarLength45 = math.Sqrt(computation.ClothWidth * computation.ClothWidth * 2)
	computation.BarLength90 = computation.ClothWidth
	computation.BarLength180 = computation.BarLength
	computation.ClothLength45 = computation.BarLength / computation.BarLength45 * computation.BarWidth * float64(computation.CountNum)
	computation.ClothLength90 = computation.BarWidth * float64(computation.CountNum)
	computation.ClothLength180 = float64(computation.CountNum) / (computation.ClothWidth / computation.BarWidth) * computation.BarLength180
	err = global.GVA_DB.Create(computation).Error
	return err
}
