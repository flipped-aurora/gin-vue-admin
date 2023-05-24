package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type RechargeOptionService struct {
}

// CreateRechargeOption 创建RechargeOption记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeOptionService *RechargeOptionService) CreateRechargeOption(rechargeOption *clothing.RechargeOption) (err error) {
	err = global.GVA_DB.Create(rechargeOption).Error
	return err
}

// DeleteRechargeOption 删除RechargeOption记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeOptionService *RechargeOptionService)DeleteRechargeOption(rechargeOption clothing.RechargeOption) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.RechargeOption{}).Where("id = ?", rechargeOption.ID).Update("deleted_by", rechargeOption.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&rechargeOption).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRechargeOptionByIds 批量删除RechargeOption记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeOptionService *RechargeOptionService)DeleteRechargeOptionByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.RechargeOption{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.RechargeOption{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRechargeOption 更新RechargeOption记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeOptionService *RechargeOptionService)UpdateRechargeOption(rechargeOption clothing.RechargeOption) (err error) {
	err = global.GVA_DB.Save(&rechargeOption).Error
	return err
}

// GetRechargeOption 根据id获取RechargeOption记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeOptionService *RechargeOptionService)GetRechargeOption(id uint) (rechargeOption clothing.RechargeOption, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rechargeOption).Error
	return
}

// GetRechargeOptionInfoList 分页获取RechargeOption记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeOptionService *RechargeOptionService)GetRechargeOptionInfoList(info clothingReq.RechargeOptionSearch) (list []clothing.RechargeOption, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.RechargeOption{})
    var rechargeOptions []clothing.RechargeOption
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&rechargeOptions).Error
	return  rechargeOptions, total, err
}
