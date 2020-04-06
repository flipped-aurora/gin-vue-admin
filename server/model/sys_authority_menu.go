package model

import (
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
// @param     menus           []SysBaseMenu
// @param     authorityId     string
// @return                    error
func (m *SysMenu) AddMenuAuthority(menus []SysBaseMenu, authorityId string) (err error) {
	var auth SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = auth.SetMuneAuthority()
	return err
}

// @title    GetMenuAuthority
// @description   查看当前角色树
// @auth                     （2020/04/05  20:22 ）
// @param     authorityId     string
// @return    err             error
// @return    menus           []SysBaseMenu
func (m *SysMenu) GetMenuAuthority(authorityId string) (err error, menus []SysMenu) {
	SQLstatement := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ?"
	err = global.GVA_DB.Raw(SQLstatement, authorityId).Scan(&menus).Error
	return err, menus
}

// @title    GetMenuTree
// @description   获取动态菜单树
// @auth                     （2020/04/05  20:22 ）
// @param     authorityId     string
// @return    err             error
// @return    menus           []SysMenu
func (m *SysMenu) GetMenuTree(authorityId string) (err error, menus []SysMenu) {
	SQLstatement := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? AND authority_menu.parent_id = ?"

	err = global.GVA_DB.Raw(SQLstatement, authorityId, 0).Scan(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], SQLstatement)
	}
	return err, menus
}

// @title    getChildrenList
// @description   获取子菜单
// @auth                     （2020/04/05  20:22 ）
// @param     menu            *SysMenu
// @param     SQLstatement    string
// @return    err             error
func getChildrenList(menu *SysMenu, SQLstatement string) (err error) {
	err = global.GVA_DB.Raw(SQLstatement, menu.AuthorityId, menu.MenuId).Scan(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], SQLstatement)
	}
	return err
}
