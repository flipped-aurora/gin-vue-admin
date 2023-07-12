package Nestrolepkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	nestrolepkg "github.com/flipped-aurora/gin-vue-admin/server/model/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	nestrolepkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/nestrolepkg/request"
	"gorm.io/gorm"
)

type NestRoleService struct {
}

// CreateNestRole 创建NestRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestroleService *NestRoleService) CreateNestRole(nestrole *nestrolepkg.NestRole) (err error) {
	err = global.GVA_DB.Create(nestrole).Error
	return err
}

// DeleteNestRole 删除NestRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestroleService *NestRoleService) DeleteNestRole(nestrole nestrolepkg.NestRole) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&nestrolepkg.NestRole{}).Where("id = ?", nestrole.ID).Update("deleted_by", nestrole.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&nestrole).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteNestRoleByIds 批量删除NestRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestroleService *NestRoleService) DeleteNestRoleByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&nestrolepkg.NestRole{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&nestrolepkg.NestRole{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateNestRole 更新NestRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestroleService *NestRoleService) UpdateNestRole(nestrole nestrolepkg.NestRole) (err error) {
	err = global.GVA_DB.Save(&nestrole).Error
	return err
}

// GetNestRole 根据id获取NestRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestroleService *NestRoleService) GetNestRole(id uint) (nestrole nestrolepkg.NestRole, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&nestrole).Error
	return
}

// GetNestRoleInfoList 分页获取NestRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (nestroleService *NestRoleService) GetNestRoleInfoList(info nestrolepkgReq.NestRoleSearch) (list []nestrolepkg.NestRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&nestrolepkg.NestRole{})
	var nestroles []nestrolepkg.NestRole
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RoleidSearch != "" {
		db = db.Where("roleid IN (?)", info.RoleidSearch)
	}
	if info.Nestid != "" {
		db = db.Where("nestid LIKE ?", "%"+info.Nestid+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&nestroles).Error
	return nestroles, total, err
}
