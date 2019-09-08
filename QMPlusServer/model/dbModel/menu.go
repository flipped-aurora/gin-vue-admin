package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
)

type Menu struct {
	gorm.Model  `json:"-"`
	MenuLevel   uint   `json:"-"`
	AuthorityId uint   `json:"-"`
	ParentId    uint   `json:"parentId"`
	Path        string `json:"path"`
	Name        string `json:"name"`
	Hidden      bool   `json:"hidden"`
	Component   string `json:"component"`
	Meta        `json:"meta"`
	Children    []Menu `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
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
	err = qmsql.DEFAULTDB.Where("authority_id = ? AND parent_id = ?", menu.AuthorityId, menu.ID).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i])
	}
	return err
}
