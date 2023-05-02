package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type AppRoleService struct {
}

// CreateAppRole 创建AppRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRoleService *AppRoleService) CreateAppRole(appRole *clothing.AppRole) (err error) {
	err = global.GVA_DB.Create(appRole).Error
	return err
}

// DeleteAppRole 删除AppRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRoleService *AppRoleService)DeleteAppRole(appRole clothing.AppRole) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.AppRole{}).Where("id = ?", appRole.ID).Update("deleted_by", appRole.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&appRole).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAppRoleByIds 批量删除AppRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRoleService *AppRoleService)DeleteAppRoleByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.AppRole{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.AppRole{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAppRole 更新AppRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRoleService *AppRoleService)UpdateAppRole(appRole clothing.AppRole) (err error) {
	err = global.GVA_DB.Save(&appRole).Error
	return err
}

// GetAppRole 根据id获取AppRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRoleService *AppRoleService)GetAppRole(id uint) (appRole clothing.AppRole, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appRole).Error
	return
}

// GetAppRoleInfoList 分页获取AppRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRoleService *AppRoleService)GetAppRoleInfoList(info clothingReq.AppRoleSearch) (list []clothing.AppRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.AppRole{})
    var appRoles []clothing.AppRole
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&appRoles).Error
	return  appRoles, total, err
}
