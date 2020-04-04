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

//创建用户
func (e *ExaCustomer) CreateExaCustomer() (err error) {
	err = global.GVA_DB.Create(e).Error
	return err
}

//删除用户
func (e *ExaCustomer) DeleteExaCustomer() (err error) {
	err = global.GVA_DB.Delete(e).Error
	return err
}

//更新用户
func (e *ExaCustomer) UpdateExaCustomer() (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//获取用户信息
func (e *ExaCustomer) GetExaCustomer() (err error, customer ExaCustomer) {
	err = global.GVA_DB.Where("id = ?", e.ID).First(&customer).Error
	return
}

//获取用户列表
// 分页获取数据
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
