package clothing

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	goqrcode "github.com/skip2/go-qrcode"
	"gorm.io/gorm"
	"time"
)

type CompanyService struct {
}

// CreateCompany 创建Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) CreateCompany(companyName string, userID, agentID uint) (err error) {
	db := global.GVA_DB
	tx := db.Begin()
	defer tx.Commit()
	var company clothing.Company
	company.Name = companyName
	company.UserID = userID
	status := new(bool)
	*status = true
	company.Status = status
	if err = global.GVA_DB.Create(&company).Error; err != nil {
		tx.Rollback()
		return err
	}
	var userRole clothing.UserRole
	userRole.CompanyID = company.ID
	userRole.RoleID = enum.Boss
	userRole.UserID = userID
	if err = global.GVA_DB.Create(&userRole).Error; err != nil {
		tx.Rollback()
		return err
	}
	if agentID > 0 {
		global.GVA_DB.Model(&clothing.Agent{}).Where("id = ?", agentID).Update("member_count", gorm.Expr("member_count + ?", 1))
	}
	return err
}

// DeleteCompany 删除Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) DeleteCompany(company clothing.Company) (err error) {
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
func (companyService *CompanyService) DeleteCompanyByIds(ids request.IdsReq, deleted_by uint) (err error) {
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
func (companyService *CompanyService) UpdateCompany(company clothing.Company) (err error) {
	err = global.GVA_DB.Save(&company).Error
	return err
}

// GetCompany 根据id获取Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) GetCompany(id uint) (company clothing.Company, err error) {
	err = global.GVA_DB.Preload("User").Preload("Agent").Where("id = ?", id).First(&company).Error
	return
}

// GetCompanyInfoList 分页获取Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) GetCompanyInfoList(info clothingReq.CompanySearch) (list []clothing.Company, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Company{})
	var companys []clothing.Company
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("User").Preload("Agent").Limit(limit).Offset(offset).Find(&companys).Error
	return companys, total, err
}

func (companyService *CompanyService) GetCompanyByName(name string) (clothing.Company, error) {
	var company clothing.Company
	err := global.GVA_DB.Where("name = ?", name).First(&company).Error
	return company, err
}

func (companyService *CompanyService) CreateQrCode(company clothing.Company) (url string, err error) {
	type companyObj struct {
		ID           uint       `json:"ID" form:"ID"`
		UserID       uint       `json:"userID" form:"userID"`
		Name         string     `json:"name" form:"name"`
		ExpirationAt *time.Time `json:"expirationAt" form:"expirationAt"`
	}
	if company.ExpirationAt == nil || time.Now().Sub(*company.ExpirationAt) > 0 {
		return "", errors.New("会员已过期")
	}
	expirationAt := time.Now().AddDate(0, 0, 7)
	if expirationAt.Sub(*company.ExpirationAt).Seconds() > 0 {
		expirationAt = *company.ExpirationAt
	}
	obj := companyObj{
		ID:           company.ID,
		UserID:       company.UserID,
		Name:         company.Name,
		ExpirationAt: &expirationAt,
	}
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	global.GVA_LOG.Sugar().Info(string(b))
	encodeString := base64.StdEncoding.EncodeToString(b)
	data, err := goqrcode.Encode(encodeString, goqrcode.Medium, 256)
	oss := new(upload.Local)
	stream := bytes.NewBuffer(data)
	pic := fmt.Sprintf("qrcode/%d/%d.png", company.ID, company.ID)
	url, _, err = oss.Put(pic, stream)
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return "", err
	}
	global.GVA_DB.Model(&company).Updates(map[string]interface{}{
		"qr_code":               url,
		"qr_code_expiration_at": expirationAt,
	})
	return url, nil
}
