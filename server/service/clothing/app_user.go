package clothing

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type AppUserService struct {
}

// CreateAppUser 创建AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) CreateAppUser(appUser *clothing.AppUser) (err error) {
	err = global.GVA_DB.Create(appUser).Error
	return err
}

// DeleteAppUser 删除AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUser(appUser clothing.AppUser) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.AppUser{}).Where("id = ?", appUser.ID).Update("deleted_by", appUser.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&appUser).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAppUserByIds 批量删除AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUserByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.AppUser{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.AppUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAppUser 更新AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateAppUser(appUser clothing.AppUser) (err error) {
	err = global.GVA_DB.Save(&appUser).Error
	return err
}

// GetAppUser 根据id获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUser(id uint) (appUser clothing.AppUser, err error) {
	err = global.GVA_DB.Preload("Roles", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Role").Preload("Company")
	}).Where("id = ?", id).First(&appUser).Error
	return
}

// GetAppUserInfoList 分页获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUserInfoList(info clothingReq.AppUserSearch) (list []clothing.AppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.AppUser{})
	var appUsers []clothing.AppUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("user_name LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.StartWages != nil && info.EndWages != nil {
		db = db.Where("wages BETWEEN ? AND ? ", info.StartWages, info.EndWages)
	}
	if info.PhoneNum != "" {
		db = db.Where("phone_num = ?", info.PhoneNum)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["wages"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&appUsers).Error
	return appUsers, total, err
}

func (AppUserService *AppUserService) CheckExist(content string, checkType int) error {
	var user clothing.AppUser
	if checkType == 1 {
		if !errors.Is(global.GVA_DB.Where("username = ?", content).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return errors.New("用户名已注册")
		}
	} else {
		if !errors.Is(global.GVA_DB.Where("phone_num = ?", content).First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册
			return errors.New("手机号已注册")
		}
	}
	return nil
}

func (AppUserService *AppUserService) Register(appUser clothing.AppUser) (userInter clothing.AppUser, err error) {
	if err := AppUserService.CheckExist(appUser.Username, 1); err != nil { // 判断用户名是否注册
		return userInter, err
	}
	if err := AppUserService.CheckExist(appUser.PhoneNum, 2); err != nil { // 判断手机号是否注册
		return userInter, err
	}
	appUser.Password = utils.BcryptHash(appUser.Password)
	err = global.GVA_DB.Create(&appUser).Error
	return appUser, err
}

func (appUserService *AppUserService) Login(u *clothing.AppUser) (userInter *clothing.AppUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user clothing.AppUser
	err = global.GVA_DB.Where("username = ? or phone_num = ?", u.Username, u.PhoneNum).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

func (appUserService *AppUserService) GetAppUserList(info clothingReq.UserFilter) (list []clothing.AppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.AppUser{})
	var appUsers []clothing.AppUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("user_name LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.PhoneNum != "" {
		db = db.Where("phone_num = ?", info.PhoneNum)
	}
	if info.CompanyID != 0 {
		userIDs := make([]uint, 0)
		global.GVA_DB.Model(&clothing.UserRole{}).Where("company_id = ?", info.CompanyID).Pluck("user_id", &userIDs)
		db = db.Where("id in ?", userIDs)
	}
	if info.TeamID != 0 {
		userIDs := make([]uint, 0)
		global.GVA_DB.Model(&clothing.TeamUser{}).Where("team_id = ?", info.TeamID).Pluck("user_id", &userIDs)
		db = db.Where("id in ?", userIDs)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appUsers).Error
	return appUsers, total, err
}
