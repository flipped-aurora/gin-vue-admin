package model

import (
	"fmt"
	"gin-vue-admin/global"
)

type SysMenu struct {
	SysBaseMenu
	MenuId      string    `json:"menuId"`
	AuthorityId string    `json:"-"`
	Children    []SysMenu `json:"children"`
}

// @title    AddMenuAuthority
// @description   为角色增加menu树
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (m *SysMenu) AddMenuAuthority(menus []SysBaseMenu, authorityId string) (err error) {
	var menu SysMenu
	global.GVA_DB.Where("authority_id = ? ", authorityId).Unscoped().Delete(&SysMenu{})
	for _, v := range menus {
		menu.SysBaseMenu = v
		menu.AuthorityId = authorityId
		menu.MenuId = fmt.Sprintf("%v", v.ID)
		menu.ID = 0
		err = global.GVA_DB.Create(&menu).Error
		if err != nil {
			return err
		}
	}
	var auth SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	auth.SetMuneAuthority()
	return nil
}

// @title    GetMenuAuthority
// @description   查看当前角色树
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (m *SysMenu) GetMenuAuthority(authorityId string) (err error, menus []SysMenu) {
	err = global.GVA_DB.Where("authority_id = ?", authorityId).Find(&menus).Error
	return err, menus
}

// @title    GetMenuTree
// @description   获取动态菜单树
// @auth                     （2020/04/05  20:22 ）
// @param     newPassword     string
// @return    err             error
func (m *SysMenu) GetMenuTree(authorityId string) (err error, menus []SysMenu) {
	err = global.GVA_DB.Where("authority_id = ? AND parent_id = ?", authorityId, 0).Order("sort", true).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i])
	}
	return err, menus
}

// @title    getChildrenList
// @description   获取子菜单
// @auth                     （2020/04/05  20:22 ）
// @param     newPassword     string
// @return    err             error
func getChildrenList(menu *SysMenu) (err error) {
	err = global.GVA_DB.Where("authority_id = ? AND parent_id = ?", menu.AuthorityId, menu.MenuId).Order("sort", true).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i])
	}
	return err
}
