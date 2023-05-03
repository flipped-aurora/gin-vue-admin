package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type ClothService struct {
}

// CreateCloth 创建Cloth记录
// Author [piexlmax](https://github.com/piexlmax)
func (clothService *ClothService) CreateCloth(cloth *clothing.Cloth) (err error) {
	err = global.GVA_DB.Create(cloth).Error
	return err
}

// DeleteCloth 删除Cloth记录
// Author [piexlmax](https://github.com/piexlmax)
func (clothService *ClothService)DeleteCloth(cloth clothing.Cloth) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Cloth{}).Where("id = ?", cloth.ID).Update("deleted_by", cloth.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cloth).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteClothByIds 批量删除Cloth记录
// Author [piexlmax](https://github.com/piexlmax)
func (clothService *ClothService)DeleteClothByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Cloth{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Cloth{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCloth 更新Cloth记录
// Author [piexlmax](https://github.com/piexlmax)
func (clothService *ClothService)UpdateCloth(cloth clothing.Cloth) (err error) {
	err = global.GVA_DB.Save(&cloth).Error
	return err
}

// GetCloth 根据id获取Cloth记录
// Author [piexlmax](https://github.com/piexlmax)
func (clothService *ClothService)GetCloth(id uint) (cloth clothing.Cloth, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cloth).Error
	return
}

// GetClothInfoList 分页获取Cloth记录
// Author [piexlmax](https://github.com/piexlmax)
func (clothService *ClothService)GetClothInfoList(info clothingReq.ClothSearch) (list []clothing.Cloth, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.Cloth{})
    var cloths []clothing.Cloth
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.StyleID != nil {
        db = db.Where("style_id = ?",info.StyleID)
    }
    if info.Color != "" {
        db = db.Where("color = ?",info.Color)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&cloths).Error
	return  cloths, total, err
}
