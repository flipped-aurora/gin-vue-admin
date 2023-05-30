package clothing

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"gorm.io/gorm"
)

type TeamApplyService struct {
}

// CreateTeamApply 创建TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) CreateTeamApply(teamApply *clothing.TeamApply) (err error) {
	err = global.GVA_DB.Create(teamApply).Error
	return err
}

// DeleteTeamApply 删除TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) DeleteTeamApply(teamApply clothing.TeamApply) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.TeamApply{}).Where("id = ?", teamApply.ID).Update("deleted_by", teamApply.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&teamApply).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTeamApplyByIds 批量删除TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) DeleteTeamApplyByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.TeamApply{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.TeamApply{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTeamApply 更新TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) UpdateTeamApply(teamApply clothing.TeamApply) (err error) {
	err = global.GVA_DB.Save(&teamApply).Error
	return err
}

// GetTeamApply 根据id获取TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) GetTeamApply(id uint) (teamApply clothing.TeamApply, err error) {
	err = global.GVA_DB.Preload("Team").Preload("User").Where("id = ?", id).First(&teamApply).Error
	return
}

// GetTeamApplyInfoList 分页获取TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) GetTeamApplyInfoList(info clothingReq.TeamApplySearch) (list []clothing.TeamApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.TeamApply{})
	var teamApplys []clothing.TeamApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("Team").Preload("User").Limit(limit).Offset(offset).Find(&teamApplys).Error
	return teamApplys, total, err
}

func (teamApplyService *TeamApplyService) OptApply(apply clothing.TeamApply, status int, team clothing.Team) (err error) {
	if !team.Company.CheckLimit() {
		return errors.New("名额不足或会员已到期")
	}
	if apply.Status != nil && *apply.Status != enum.ApplyPending {
		return errors.New("已处理")
	}
	if status == enum.ApplyReject {
		err = global.GVA_DB.Model(&apply).Update("status", enum.ApplyReject).Error
		return
	}
	var role clothing.UserRole
	role.RoleID = apply.RoleID
	role.UserID = apply.UserID
	role.CompanyID = team.CompanyID
	err = global.GVA_DB.Where(&role).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&role).Error
		if err == nil {
			var wallet clothing.UserWallet
			wallet.CompanyID = team.CompanyID
			wallet.UserID = apply.UserID
			err = global.GVA_DB.Create(&wallet).Error
			global.GVA_DB.Model(&clothing.Company{}).Where("id = ?", team.CompanyID).Update("clerk_count", gorm.Expr("clerk_count + ?", 1))
		}
	}
	var teamUser clothing.TeamUser
	teamUser.TeamID = team.ID
	teamUser.UserID = apply.UserID
	err = global.GVA_DB.Where(&teamUser).First(&teamUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&teamUser).Error
	}
	if err == nil {
		err = global.GVA_DB.Model(&apply).Update("status", status).Error
	}
	return err
}
