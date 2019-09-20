package dbModel

import (
	"fmt"
	"main/init/qmsql"
)

type Menu struct {
	BaseMenu
	MenuId      string `json:"menuId"`
	AuthorityId string `json:"-"`
	Children    []Menu `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// 为角色增加menu树
func (m *Menu) AddMenuAuthority(menus []BaseMenu, authorityId string) (err error) {
	var menu Menu
	qmsql.DEFAULTDB.Where("authority_id = ? ", authorityId).Unscoped().Delete(&Menu{})
	for _, v := range menus {
		menu.BaseMenu = v
		menu.AuthorityId = authorityId
		menu.MenuId = fmt.Sprintf("%v", v.ID)
		menu.ID = 0
		err = qmsql.DEFAULTDB.Create(&menu).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// 查看当前角色树
func (m *Menu) GetMenuAuthority(authorityId string) (err error, menus []Menu) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authorityId).Find(&menus).Error
	return err, menus
}

//获取动态路由树
func (m *Menu) GetMenuTree(authorityId float64) (err error, menus []Menu) {
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", authorityId, 0).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i])
	}
	return err, menus
}

func getChildrenList(menu *Menu) (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", menu.AuthorityId, menu.MenuId).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i])
	}
	return err
}
