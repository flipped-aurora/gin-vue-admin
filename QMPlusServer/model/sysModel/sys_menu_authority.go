package sysModel

import (
	"fmt"
	"gin-vue-admin/init/qmsql"
)

// menu需要构建的点有点多 这里关联关系表直接把所有数据拿过来 用代码实现关联  后期实现主外键模式
type SysMenu struct {
	SysBaseMenu
	MenuId      string    `json:"menuId"`
	AuthorityId string    `json:"-"`
	Children    []SysMenu `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// 为角色增加menu树
func (m *SysMenu) AddMenuAuthority(menus []SysBaseMenu, authorityId string) (err error) {
	var menu SysMenu
	qmsql.DEFAULTDB.Where("authority_id = ? ", authorityId).Unscoped().Delete(&SysMenu{})
	for _, v := range menus {
		menu.SysBaseMenu = v
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
func (m *SysMenu) GetMenuAuthority(authorityId string) (err error, menus []SysMenu) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authorityId).Find(&menus).Error
	return err, menus
}

//获取动态路由树
func (m *SysMenu) GetMenuTree(authorityId string) (err error, menus []SysMenu) {
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", authorityId, 0).Order("sort", true).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i])
	}
	return err, menus
}

func getChildrenList(menu *SysMenu) (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", menu.AuthorityId, menu.MenuId).Order("sort", true).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i])
	}
	return err
}
