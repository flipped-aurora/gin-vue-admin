package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
)

type ExaCustomer struct {
	gorm.Model
	CustomerName       string  `json:"customerName"`
	CustomerPhoneData  string  `json:"customerPhoneData"`
	SysUserID          uint    `json:"sysUserId"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID"`
	SysUser            SysUser `json:"sysUser"`
}

// @title    CreateExaCustomer
// @description   create a customer, 创建用户
// @auth                     （2020/04/05  20:22 ）
// @return    err             error
func (e *ExaCustomer) CreateExaCustomer() (err error) {
	err = global.GVA_DB.Create(e).Error
	return err
}

// @title    DeleteFileChunk
// @description   delete a customer, 删除用户
// @auth                     （2020/04/05  20:22 ）
// @return                    error
func (e *ExaCustomer) DeleteExaCustomer() (err error) {
	err = global.GVA_DB.Delete(e).Error
	return err
}

// @title    UpdateExaCustomer
// @description   update a customer, 更新用户
// @auth                     （2020/04/05  20:22 ）
// @return                    error
func (e *ExaCustomer) UpdateExaCustomer() (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

// @title    GetExaCustomer
// @description   get the info of a costumer , 获取用户信息
// @auth                     （2020/04/05  20:22 ）
// @return                    error
// @return    customer        ExaCustomer
func (e *ExaCustomer) GetExaCustomer() (err error, customer ExaCustomer) {
	err = global.GVA_DB.Where("id = ?", e.ID).First(&customer).Error
	return
}

// @title    GetInfoList
// @description   get customer list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22 ）
// @param     info            PageInfo
// @return                    error
func (e *ExaCustomer) GetInfoList(info PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	if err != nil {
		return
	} else {
		var a SysAuthority
		a.AuthorityId = e.SysUserAuthorityID
		err, auth := a.GetAuthorityInfo()
		var dataId []string
		for _, v := range auth.DataAuthorityId {
			dataId = append(dataId, v.AuthorityId)
		}
		var CustomerList []ExaCustomer
		err = db.Where("sys_user_authority_id in (?)", dataId).Find(&CustomerList).Count(&total).Error
		if err != nil {
			return err, CustomerList, total
		} else {
			err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in (?)", dataId).Find(&CustomerList).Error
		}
		return err, CustomerList, total
	}
}
