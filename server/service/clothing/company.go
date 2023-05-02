package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type CompanyService struct {
}

// CreateCompany 创建Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) CreateCompany(company *clothing.Company) (err error) {
	err = global.GVA_DB.Create(company).Error
	return err
}

// DeleteCompany 删除Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService)DeleteCompany(company clothing.Company) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Company{}).Where("id = ?", company.ID).Update("deleted_by", company.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&company).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCompanyByIds 批量删除Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService)DeleteCompanyByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Company{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Company{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCompany 更新Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService)UpdateCompany(company clothing.Company) (err error) {
	err = global.GVA_DB.Save(&company).Error
	return err
}

// GetCompany 根据id获取Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService)GetCompany(id uint) (company clothing.Company, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&company).Error
	return
}

// GetCompanyInfoList 分页获取Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService)GetCompanyInfoList(info clothingReq.CompanySearch) (list []clothing.Company, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.Company{})
    var companys []clothing.Company
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&companys).Error
	return  companys, total, err
}
