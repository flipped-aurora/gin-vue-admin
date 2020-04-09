package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAuthority
// @description   创建一个角色
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func CreateAuthority(a model.SysAuthority) (err error, authority model.SysAuthority) {
	err = global.GVA_DB.Create(&a).Error
	return err, a
}

// @title    DeleteAuthority
// @description   删除角色
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
// 删除角色
func DeleteAuthority(a model.SysAuthority) (err error) {
	err = global.GVA_DB.Where("authority_id = ?", a.AuthorityId).Find(&model.SysUser{}).Error
	if err == nil {
		err = errors.New("此角色有用户正在使用禁止删除")
		return
	}
	err = global.GVA_DB.Where("parent_id = ?", a.AuthorityId).Find(&model.SysAuthority{}).Error
	if err == nil {
		err = errors.New("此角色存在子角色不允许删除")
		return
	}
	db := global.GVA_DB.Preload("SysBaseMenus").Where("authority_id = ?", a.AuthorityId).First(a).Unscoped().Delete(a)
	if len(a.SysBaseMenus) > 0 {
		err = db.Association("SysBaseMenus").Delete(a.SysBaseMenus).Error
	} else {
		err = db.Error
	}
	ClearCasbin(0, a.AuthorityId)
	return err
}

// @title    GetInfoList
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
// 分页获取数据
func GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k, _ := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

// @title    GetAuthorityInfo
// @description   获取所有角色信息
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func GetAuthorityInfo(a model.SysAuthority) (err error, sa model.SysAuthority) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", a.AuthorityId).First(&sa).Error
	return err, sa
}

// @title    SetDataAuthority
// @description   设置角色资源权限
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func SetDataAuthority(a model.SysAuthority) error {
	var s model.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", a.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&a.DataAuthorityId).Error
	return err
}

// @title    SetMenuAuthority
// @description   菜单与角色绑定
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func SetMenuAuthority(a *model.SysAuthority) error {
	var s model.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", a.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&a.SysBaseMenus).Error
	return err
}

// @title    findChildrenAuthority
// @description   查询子角色
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func findChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k, _ := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
