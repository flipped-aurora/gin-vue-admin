package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// @title    DeleteBaseMenu
// @description   删除基础路由
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error
func  DeleteBaseMenu(id float64) (err error) {
	err = global.GVA_DB.Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := global.GVA_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		if len(menu.SysAuthoritys) > 0 {
			err = db.Association("SysAuthoritys").Delete(menu.SysAuthoritys).Error
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// @title    UpdateBaseMenu
// @description   更新路由
// @auth                     （2020/04/05  20:22）
// @param     menu            model.SysBaseMenu
// @return    err             error
func  UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort
	err = global.GVA_DB.Where("id = ?", menu.ID).Find(&model.SysBaseMenu{}).Updates(upDateMap).Error
	global.GVA_LOG.Debug("菜单修改时候，关联菜单err:%v", err)
	return err
}

// @title    GetBaseMenuById
// @description   get current menus, 返回当前选中menu
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error
func GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = global.GVA_DB.Where("id = ?", id).First(&menu).Error
	return
}
