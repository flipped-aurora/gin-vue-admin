package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type UserWalletService struct {
}

// CreateUserWallet 创建UserWallet记录
// Author [piexlmax](https://github.com/piexlmax)
func (userWalletService *UserWalletService) CreateUserWallet(userWallet *clothing.UserWallet) (err error) {
	err = global.GVA_DB.Create(userWallet).Error
	return err
}

// DeleteUserWallet 删除UserWallet记录
// Author [piexlmax](https://github.com/piexlmax)
func (userWalletService *UserWalletService) DeleteUserWallet(userWallet clothing.UserWallet) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.UserWallet{}).Where("id = ?", userWallet.ID).Update("deleted_by", userWallet.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&userWallet).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteUserWalletByIds 批量删除UserWallet记录
// Author [piexlmax](https://github.com/piexlmax)
func (userWalletService *UserWalletService) DeleteUserWalletByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.UserWallet{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.UserWallet{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUserWallet 更新UserWallet记录
// Author [piexlmax](https://github.com/piexlmax)
func (userWalletService *UserWalletService) UpdateUserWallet(userWallet clothing.UserWallet) (err error) {
	err = global.GVA_DB.Save(&userWallet).Error
	return err
}

// GetUserWallet 根据id获取UserWallet记录
// Author [piexlmax](https://github.com/piexlmax)
func (userWalletService *UserWalletService) GetUserWallet(id uint) (userWallet clothing.UserWallet, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userWallet).Error
	return
}

// GetUserWalletInfoList 分页获取UserWallet记录
// Author [piexlmax](https://github.com/piexlmax)
func (userWalletService *UserWalletService) GetUserWalletInfoList(info clothingReq.UserWalletSearch) (list []clothing.UserWallet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.UserWallet{})
	var userWallets []clothing.UserWallet
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserID != 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.CompanyID != 0 {
		db = db.Where("company_id = ?", info.CompanyID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userWallets).Error
	return userWallets, total, err
}
