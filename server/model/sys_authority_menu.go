package model

import (
	"gin-vue-admin/global"
)

// menu需要构建的点有点多 这里关联关系表直接把所有数据拿过来 用代码实现关联  后期实现主外键模式
type SysMenu struct {
	SysBaseMenu
	MenuID      string    `json:"menuId"`
	AuthorityId string    `json:"-"`
	Children    []SysMenu `json:"children"`
}

// 为角色增加menu树
func (m *SysMenu) AddMenuAuthority(menus []SysBaseMenu, authorityId string) (err error) {
	var auth SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	auth.SetMuneAuthority()
	return nil
}

// 查看当前角色树
func (m *SysMenu) GetMenuAuthority(authorityId string) (err error, menus []SysBaseMenu) {
	var a SysAuthority
	err = global.GVA_DB.Preload("SysBaseMenus").Where("authority_id = ?", authorityId).First(&a).Error
	return err, a.SysBaseMenus
}

//获取动态路由树
func (m *SysMenu) GetMenuTree(authorityId string) (err error, menus []SysMenu) {
	SQLstatement := "SELECT authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? AND authority_menu.parent_id = ?"

	err = global.GVA_DB.Raw(SQLstatement, authorityId, 0).Scan(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], SQLstatement)
	}
	return err, menus
}

func getChildrenList(menu *SysMenu, SQLstatement string) (err error) {
	err = global.GVA_DB.Raw(SQLstatement, menu.AuthorityId, menu.MenuID).Scan(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], SQLstatement)
	}
	return err
}
