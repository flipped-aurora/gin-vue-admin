package UserTeemlinkPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/UserTeemlinkPkg"
	UserTeemlinkPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/UserTeemlinkPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type UserTeemlinkService struct {
}

// CreateUserTeemlink 创建UserTeemlink记录
// Author [piexlmax](https://github.com/piexlmax)
func (usertlService *UserTeemlinkService) CreateUserTeemlink(usertl *UserTeemlinkPkg.UserTeemlink) (err error) {
	err = global.GVA_DB.Create(usertl).Error
	return err
}

// DeleteUserTeemlink 删除UserTeemlink记录
// Author [piexlmax](https://github.com/piexlmax)
func (usertlService *UserTeemlinkService) DeleteUserTeemlink(usertl UserTeemlinkPkg.UserTeemlink) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&UserTeemlinkPkg.UserTeemlink{}).Where("id = ?", usertl.ID).Update("deleted_by", usertl.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&usertl).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteUserTeemlinkByIds 批量删除UserTeemlink记录
// Author [piexlmax](https://github.com/piexlmax)
func (usertlService *UserTeemlinkService) DeleteUserTeemlinkByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&UserTeemlinkPkg.UserTeemlink{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&UserTeemlinkPkg.UserTeemlink{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUserTeemlink 更新UserTeemlink记录
// Author [piexlmax](https://github.com/piexlmax)
func (usertlService *UserTeemlinkService) UpdateUserTeemlink(usertl UserTeemlinkPkg.UserTeemlink) (err error) {
	err = global.GVA_DB.Save(&usertl).Error
	return err
}

// GetUserTeemlink 根据id获取UserTeemlink记录
// Author [piexlmax](https://github.com/piexlmax)
func (usertlService *UserTeemlinkService) GetUserTeemlink(id uint) (usertl UserTeemlinkPkg.UserTeemlink, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&usertl).Error
	return
}

func (usertlService *UserTeemlinkService) GetUserTeemlinkByTlID(tlid string) (usertl UserTeemlinkPkg.UserTeemlink, err error) {
	err = global.GVA_DB.Where("tl_user_id = ?", tlid).First(&usertl).Error
	return
}

// GetUserTeemlinkInfoList 分页获取UserTeemlink记录
// Author [piexlmax](https://github.com/piexlmax)
func (usertlService *UserTeemlinkService) GetUserTeemlinkInfoList(info UserTeemlinkPkgReq.UserTeemlinkSearch) (list []UserTeemlinkPkg.UserTeemlink, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&UserTeemlinkPkg.UserTeemlink{})
	var usertls []UserTeemlinkPkg.UserTeemlink
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.TlUserID != "" {
		db = db.Where("tl_user_id LIKE ?", "%"+info.TlUserID+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&usertls).Error
	return usertls, total, err
}
