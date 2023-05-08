package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type TeamService struct {
}

// CreateTeam 创建Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) CreateTeam(team *clothing.Team) (err error) {
	err = global.GVA_DB.Create(team).Error
	return err
}

// DeleteTeam 删除Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) DeleteTeam(team clothing.Team) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Team{}).Where("id = ?", team.ID).Update("deleted_by", team.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&team).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTeamByIds 批量删除Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) DeleteTeamByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.Team{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Team{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTeam 更新Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) UpdateTeam(team clothing.Team) (err error) {
	err = global.GVA_DB.Save(&team).Error
	return err
}

// GetTeam 根据id获取Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) GetTeam(id uint) (team clothing.Team, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&team).Error
	return
}

// GetTeamInfoList 分页获取Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) GetTeamInfoList(info clothingReq.TeamSearch) (list []clothing.Team, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Team{})
	var teams []clothing.Team
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CompanyID != 0 {
		db = db.Where("company_id = ?", info.CompanyID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&teams).Error
	return teams, total, err
}

func (teamService *TeamService) GetTeamByName(company clothing.Company, name string) (team clothing.Team, err error) {
	err = global.GVA_DB.Where("company_id = ? and name = ?", company.ID, name).First(&team).Error
	return
}
