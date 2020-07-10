package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateExaCustomer
// @description   create a customer, 创建用户
// @param     e               model.ExaCustomer
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateExaCustomer(e model.ExaCustomer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// @title    DeleteFileChunk
// @description   delete a customer, 删除用户
// @auth                     （2020/04/05  20:22）
// @param     e               model.ExaCustomer
// @return                    error

func DeleteExaCustomer(e model.ExaCustomer) (err error) {
	err = global.GVA_DB.Delete(e).Error
	return err
}

// @title    UpdateExaCustomer
// @description   update a customer, 更新用户
// @param     e               *model.ExaCustomer
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateExaCustomer(e *model.ExaCustomer) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

// @title    GetExaCustomer
// @description   get the info of a costumer , 获取用户信息
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    customer        ExaCustomer

func GetExaCustomer(id uint) (err error, customer model.ExaCustomer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

// @title    GetCustomerInfoList
// @description   get customer list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     sysUserAuthorityID              string
// @param     info            PageInfo
// @return                    error

func GetCustomerInfoList(sysUserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ExaCustomer{})
	var a model.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	err, auth := GetAuthorityInfo(a)
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []model.ExaCustomer
	err = db.Where("sys_user_authority_id in (?)", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in (?)", dataId).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
