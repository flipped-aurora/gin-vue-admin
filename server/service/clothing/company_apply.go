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

type CompanyApplyService struct {
}

// CreateCompanyApply 创建CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService) CreateCompanyApply(companyApply *clothing.CompanyApply) (err error) {
	err = global.GVA_DB.Create(companyApply).Error
	return err
}

// DeleteCompanyApply 删除CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService) DeleteCompanyApply(companyApply clothing.CompanyApply) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.CompanyApply{}).Where("id = ?", companyApply.ID).Update("deleted_by", companyApply.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&companyApply).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCompanyApplyByIds 批量删除CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService) DeleteCompanyApplyByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.CompanyApply{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.CompanyApply{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCompanyApply 更新CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService) UpdateCompanyApply(companyApply clothing.CompanyApply) (err error) {
	err = global.GVA_DB.Save(&companyApply).Error
	return err
}

// GetCompanyApply 根据id获取CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService) GetCompanyApply(id uint) (companyApply clothing.CompanyApply, err error) {
	err = global.GVA_DB.Preload("Company").Preload("User").Where("id = ?", id).First(&companyApply).Error
	return
}

// GetCompanyApplyInfoList 分页获取CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService) GetCompanyApplyInfoList(info clothingReq.CompanyApplySearch) (list []clothing.CompanyApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.CompanyApply{})
	var companyApplys []clothing.CompanyApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&companyApplys).Error
	return companyApplys, total, err
}

func (companyApplyService *CompanyApplyService) JoinCompany(roleID uint, userID uint, company clothing.Company) (err error) {
	var role clothing.UserRole
	role.RoleID = roleID
	role.UserID = userID
	role.CompanyID = company.ID
	err = global.GVA_DB.Where(&role).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&role).Error
		if err == nil {
			var wallet clothing.UserWallet
			wallet.CompanyID = company.ID
			wallet.UserID = userID
			err = global.GVA_DB.Create(&wallet).Error
			global.GVA_DB.Model(&company).Update("clerk_count", gorm.Expr("clerk_count + ?", 1))
		}
	}
	return err
}

func (companyApplyService *CompanyApplyService) OptApply(apply clothing.CompanyApply, status int, company clothing.Company) (err error) {
	if apply.Status != nil && *apply.Status != enum.ApplyPending {
		return errors.New("已处理")
	}
	if status == enum.ApplyReject {
		err = global.GVA_DB.Model(&apply).Update("status", enum.ApplyReject).Error
		return
	}
	switch apply.RoleID {
	case enum.Tailor:
		err = companyApplyService.JoinCompany(enum.Tailor, apply.UserID, company)
	case enum.GroupLeader:
		var team clothing.Team
		team.CompanyID = company.ID
		team.UserID = apply.UserID
		team.Name = apply.Remark
		err = global.GVA_DB.Create(&team).Error
		if err == nil {
			if err = companyApplyService.JoinCompany(enum.GroupLeader, apply.UserID, company); err == nil {
				var teamUser clothing.TeamUser
				teamUser.TeamID = team.ID
				teamUser.UserID = apply.UserID
				err = global.GVA_DB.Where(&teamUser).First(&teamUser).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = global.GVA_DB.Create(&teamUser).Error
				}
			}
		}
	default:
		err = errors.New("角色类型错误")
	}
	if err == nil {
		err = global.GVA_DB.Model(&apply).Update("status", status).Error
	}
	return err
}
