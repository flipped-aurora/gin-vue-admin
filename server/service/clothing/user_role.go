package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type UserRoleService struct {
}

// CreateUserRole 创建UserRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (userRoleService *UserRoleService) CreateUserRole(userRole *clothing.UserRole) (err error) {
	err = global.GVA_DB.Create(userRole).Error
	return err
}

// DeleteUserRole 删除UserRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (userRoleService *UserRoleService)DeleteUserRole(userRole clothing.UserRole) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.UserRole{}).Where("id = ?", userRole.ID).Update("deleted_by", userRole.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&userRole).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteUserRoleByIds 批量删除UserRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (userRoleService *UserRoleService)DeleteUserRoleByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.UserRole{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.UserRole{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateUserRole 更新UserRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (userRoleService *UserRoleService)UpdateUserRole(userRole clothing.UserRole) (err error) {
	err = global.GVA_DB.Save(&userRole).Error
	return err
}

// GetUserRole 根据id获取UserRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (userRoleService *UserRoleService)GetUserRole(id uint) (userRole clothing.UserRole, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userRole).Error
	return
}

// GetUserRoleInfoList 分页获取UserRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (userRoleService *UserRoleService)GetUserRoleInfoList(info clothingReq.UserRoleSearch) (list []clothing.UserRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.UserRole{})
    var userRoles []clothing.UserRole
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&userRoles).Error
	return  userRoles, total, err
}
