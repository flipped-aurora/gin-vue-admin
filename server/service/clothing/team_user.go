package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type TeamUserService struct {
}

// CreateTeamUser 创建TeamUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamUserService *TeamUserService) CreateTeamUser(teamUser *clothing.TeamUser) (err error) {
	err = global.GVA_DB.Create(teamUser).Error
	return err
}

// DeleteTeamUser 删除TeamUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamUserService *TeamUserService)DeleteTeamUser(teamUser clothing.TeamUser) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.TeamUser{}).Where("id = ?", teamUser.ID).Update("deleted_by", teamUser.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&teamUser).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteTeamUserByIds 批量删除TeamUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamUserService *TeamUserService)DeleteTeamUserByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.TeamUser{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.TeamUser{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateTeamUser 更新TeamUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamUserService *TeamUserService)UpdateTeamUser(teamUser clothing.TeamUser) (err error) {
	err = global.GVA_DB.Save(&teamUser).Error
	return err
}

// GetTeamUser 根据id获取TeamUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamUserService *TeamUserService)GetTeamUser(id uint) (teamUser clothing.TeamUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamUser).Error
	return
}

// GetTeamUserInfoList 分页获取TeamUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamUserService *TeamUserService)GetTeamUserInfoList(info clothingReq.TeamUserSearch) (list []clothing.TeamUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.TeamUser{})
    var teamUsers []clothing.TeamUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&teamUsers).Error
	return  teamUsers, total, err
}
