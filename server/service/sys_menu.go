package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    GetMenuTree
// @description   获取动态菜单树
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return    err             error
// @return    menus           []model.SysMenu
func GetMenuTree(authorityId string) (err error, menus []model.SysMenu) {
	//sql := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM sys_menus authority_menu WHERE authority_menu.authority_id = ? AND authority_menu.parent_id = ?"
	sql := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.id menu_id,am.sys_authority_authority_id authority_id FROM sys_base_menus authority_menu left join  sys_authority_menus am on authority_menu.id = am.sys_base_menu_id WHERE am.sys_authority_authority_id = ? AND authority_menu.parent_id = ? "
	err = global.GVA_DB.Raw(sql, authorityId, 0).Scan(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], sql)
	}
	return err, menus
}

// @title    getChildrenList
// @description   获取子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysMenu
// @param     sql             string
// @return    err             error
func getChildrenList(menu *model.SysMenu, sql string) (err error) {
	err = global.GVA_DB.Raw(sql, menu.AuthorityId, menu.MenuId).Scan(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], sql)
	}
	return err
}

// @title    GetInfoList
// @description   获取路由分页
// @auth                     （2020/04/05  20:22）
// @param     info            request.PageInfo
// @return    err             error
// @return    list            interface{}
// @return    total           int
func GetInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var menuList []model.SysBaseMenu
	err = db.Limit(limit).Offset(offset).Where("parent_id = 0").Order("sort", true).Find(&menuList).Error
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i])
	}
	return err, menuList, total
}

// @title    getBaseChildrenList
// @description   get children of menu, 获取菜单的子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysBaseMenu
// @return    err             error
func getBaseChildrenList(menu *model.SysBaseMenu) (err error) {
	err = global.GVA_DB.Where("parent_id = ?", menu.ID).Order("sort", true).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i])
	}
	return err
}

// @title    AddBaseMenu
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysBaseMenu
// @return    err             error
//增加基础路由
func AddBaseMenu(menu model.SysBaseMenu) (err error) {
	findOne := global.GVA_DB.Where("name = ?", menu.Name).Find(&model.SysBaseMenu{}).Error
	if findOne != nil {
		err = global.GVA_DB.Create(&menu).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

// @title    GetBaseMenuTree
// @description   获取基础路由树
// @auth                     （2020/04/05  20:22）
// @return    err              error
// @return    menus            []SysBaseMenu
func GetBaseMenuTree() (err error, menus []model.SysBaseMenu) {
	err = global.GVA_DB.Where(" parent_id = ?", 0).Order("sort", true).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i])
	}
	return err, menus
}

// @title    AddMenuAuthority
// @description   为角色增加menu树
// @auth                     （2020/04/05  20:22）
// @param     menus           []model.SysBaseMenu
// @param     authorityId     string
// @return                    error
func AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) (err error) {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

// @title    GetMenuAuthority
// @description   查看当前角色树
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return    err             error
// @return    menus           []SysBaseMenu
func GetMenuAuthority(authorityId string) (err error, menus []model.SysMenu) {
	//sql := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM sys_menus authority_menu WHERE authority_menu.authority_id = ?"
	sql := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.id menu_id,am.sys_authority_authority_id authority_id FROM sys_base_menus authority_menu left join  sys_authority_menus am on authority_menu.id = am.sys_base_menu_id WHERE am.sys_authority_authority_id = ? "
	err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
