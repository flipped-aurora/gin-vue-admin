package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
)

type RecruitService struct {
}

// CreateRecruit 创建Recruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (recruitService *RecruitService) CreateRecruit(recruit *webcms.Recruit) (err error) {
	err = global.GVA_DB.Create(recruit).Error
	return err
}

// DeleteRecruit 删除Recruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (recruitService *RecruitService) DeleteRecruit(recruit webcms.Recruit) (err error) {
	err = global.GVA_DB.Delete(&recruit).Error
	return err
}

// DeleteRecruitByIds 批量删除Recruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (recruitService *RecruitService) DeleteRecruitByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Recruit{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRecruit 更新Recruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (recruitService *RecruitService) UpdateRecruit(recruit webcms.Recruit) (err error) {
	err = global.GVA_DB.Save(&recruit).Error
	return err
}

// GetRecruit 根据id获取Recruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (recruitService *RecruitService) GetRecruit(id uint) (recruit webcms.Recruit, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&recruit).Error
	return
}

// GetRecruitInfoList 分页获取Recruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (recruitService *RecruitService) GetRecruitInfoList(info webcmsReq.RecruitSearch, siteid string) (list []webcms.Recruit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Recruit{}).Where("siteid", siteid)
	var recruits []webcms.Recruit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CateId != "0" {
		db.Where("cate_id = ?", info.CateId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&recruits).Error
	return recruits, total, err
}
