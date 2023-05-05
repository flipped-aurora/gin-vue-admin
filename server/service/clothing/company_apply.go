package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
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
func (companyApplyService *CompanyApplyService)DeleteCompanyApply(companyApply clothing.CompanyApply) (err error) {
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
func (companyApplyService *CompanyApplyService)DeleteCompanyApplyByIds(ids request.IdsReq,deleted_by uint) (err error) {
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
func (companyApplyService *CompanyApplyService)UpdateCompanyApply(companyApply clothing.CompanyApply) (err error) {
	err = global.GVA_DB.Save(&companyApply).Error
	return err
}

// GetCompanyApply 根据id获取CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService)GetCompanyApply(id uint) (companyApply clothing.CompanyApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&companyApply).Error
	return
}

// GetCompanyApplyInfoList 分页获取CompanyApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyApplyService *CompanyApplyService)GetCompanyApplyInfoList(info clothingReq.CompanyApplySearch) (list []clothing.CompanyApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.CompanyApply{})
    var companyApplys []clothing.CompanyApply
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&companyApplys).Error
	return  companyApplys, total, err
}
