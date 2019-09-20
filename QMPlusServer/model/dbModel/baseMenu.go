package dbModel

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"main/controller/servers"
	"main/init/qmsql"
	"main/model/modelInterface"
)

type BaseMenu struct {
	gorm.Model
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Meta      `json:"meta"`
	NickName  string     `json:"nickName"`
	Children  []BaseMenu `json:"children"`
}

func (b *BaseMenu) AddBaseMenu() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", b.Name).Find(&BaseMenu{}).Error
	if findOne != nil {
		b.NickName = b.Title
		err = qmsql.DEFAULTDB.Create(b).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

func (b *BaseMenu) DeleteBaseMenu(name string) (err error) {
	err = qmsql.DEFAULTDB.Where("name = ?", name).Delete(&b).Delete(&Menu{}).Error
	return err
}

func (b *BaseMenu) UpdataBaseMenu(name string) (err error) {
	err = qmsql.DEFAULTDB.Where("name = ?", name).Update(&b).Update(&Menu{}).Error
	return err
}

func (b *BaseMenu) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(b, info)
	if err != nil {
		return
	} else {
		var menuList []BaseMenu
		err = db.Find(&menuList).Error
		return err, menuList, total
	}
}

//获取基础路由树
func (m *BaseMenu) GetBaseMenuTree() (err error, menus []BaseMenu) {
	err = qmsql.DEFAULTDB.Where(" parent_id = ?", 0).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i])
	}
	return err, menus
}

func getBaseChildrenList(menu *BaseMenu) (err error) {
	err = qmsql.DEFAULTDB.Where("parent_id = ?", menu.ID).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i])
	}
	return err
}
