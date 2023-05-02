package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
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
func (appUserService *AppUserService)DeleteAppUser(appUser clothing.AppUser) (err error) {
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
func (appUserService *AppUserService)DeleteAppUserByIds(ids request.IdsReq,deleted_by uint) (err error) {
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
func (appUserService *AppUserService)UpdateAppUser(appUser clothing.AppUser) (err error) {
	err = global.GVA_DB.Save(&appUser).Error
	return err
}

// GetAppUser 根据id获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService)GetAppUser(id uint) (appUser clothing.AppUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appUser).Error
	return
}

// GetAppUserInfoList 分页获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService)GetAppUserInfoList(info clothingReq.AppUserSearch) (list []clothing.AppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.AppUser{})
    var appUsers []clothing.AppUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.UserName != "" {
        db = db.Where("user_name LIKE ?","%"+ info.UserName+"%")
    }
    if info.RoleID != nil {
        db = db.Where("role_id = ?",info.RoleID)
    }
    if info.Nickname != "" {
        db = db.Where("nickname LIKE ?","%"+ info.Nickname+"%")
    }
    if info.CompanyID != nil {
        db = db.Where("company_id = ?",info.CompanyID)
    }
        if info.StartWages != nil && info.EndWages != nil {
            db = db.Where("wages BETWEEN ? AND ? ",info.StartWages,info.EndWages)
        }
    if info.WorkType != nil {
        db = db.Where("work_type = ?",info.WorkType)
    }
	err = db.Count(&total).Error
	if err!=nil {
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
	return  appUsers, total, err
}
