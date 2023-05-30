package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type SizeListService struct {
}

// CreateSizeList 创建SizeList记录
// Author [piexlmax](https://github.com/piexlmax)
func (sizeListService *SizeListService) CreateSizeList(sizeList *clothing.SizeList) (err error) {
	err = global.GVA_DB.Create(sizeList).Error
	return err
}

// DeleteSizeList 删除SizeList记录
// Author [piexlmax](https://github.com/piexlmax)
func (sizeListService *SizeListService)DeleteSizeList(sizeList clothing.SizeList) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.SizeList{}).Where("id = ?", sizeList.ID).Update("deleted_by", sizeList.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&sizeList).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteSizeListByIds 批量删除SizeList记录
// Author [piexlmax](https://github.com/piexlmax)
func (sizeListService *SizeListService)DeleteSizeListByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.SizeList{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.SizeList{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateSizeList 更新SizeList记录
// Author [piexlmax](https://github.com/piexlmax)
func (sizeListService *SizeListService)UpdateSizeList(sizeList clothing.SizeList) (err error) {
	err = global.GVA_DB.Save(&sizeList).Error
	return err
}

// GetSizeList 根据id获取SizeList记录
// Author [piexlmax](https://github.com/piexlmax)
func (sizeListService *SizeListService)GetSizeList(id uint) (sizeList clothing.SizeList, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sizeList).Error
	return
}

// GetSizeListInfoList 分页获取SizeList记录
// Author [piexlmax](https://github.com/piexlmax)
func (sizeListService *SizeListService)GetSizeListInfoList(info clothingReq.SizeListSearch) (list []clothing.SizeList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.SizeList{})
    var sizeLists []clothing.SizeList
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&sizeLists).Error
	return  sizeLists, total, err
}
