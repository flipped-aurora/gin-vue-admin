package sysModel

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SysBaseMenu struct {
	gorm.Model
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      string `json:"sort"`
	Meta      `json:"meta"`
	NickName  string        `json:"nickName"`
	Children  []SysBaseMenu `json:"children"`
}

//增加基础路由
func (b *SysBaseMenu) AddBaseMenu() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", b.Name).Find(&SysBaseMenu{}).Error
	if findOne != nil {
		b.NickName = b.Title
		err = qmsql.DEFAULTDB.Create(b).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

//删除基础路由
func (b *SysBaseMenu) DeleteBaseMenu(id float64) (err error) {
	err = qmsql.DEFAULTDB.Where("parent_id = ?", id).First(&SysBaseMenu{}).Error
	if err != nil {
		err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&b).Error
		err = qmsql.DEFAULTDB.Where("menu_id = ?", id).Unscoped().Delete(&SysMenu{}).Error
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//更新路由
func (b *SysBaseMenu) UpdataBaseMenu() (err error) {
	upDataMap := make(map[string]interface{})
	upDataMap["parent_id"] = b.ParentId
	upDataMap["path"] = b.Path
	upDataMap["name"] = b.Name
	upDataMap["hidden"] = b.Hidden
	upDataMap["component"] = b.Component
	upDataMap["title"] = b.Title
	upDataMap["icon"] = b.Icon
	upDataMap["sort"] = b.Sort
	upDataMap["nick_name"] = b.Title
	err = qmsql.DEFAULTDB.Where("id = ?", b.ID).Find(&SysBaseMenu{}).Updates(upDataMap).Error
	err1 := qmsql.DEFAULTDB.Where("menu_id = ?", b.ID).Find(&[]SysMenu{}).Updates(upDataMap).Error
	fmt.Printf("菜单修改时候，关联菜单err1:%v,err:%v", err1, err)
	return err
}

//当前选中角色所拥有的路由
func (b *SysBaseMenu) GetBaseMenuById(id float64) (err error, menu SysBaseMenu) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).First(&menu).Error
	return
}

//获取路由分页
func (b *SysBaseMenu) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(b, info)
	if err != nil {
		return
	} else {
		var menuList []SysBaseMenu
		err = db.Where("parent_id = 0").Order("sort", true).Find(&menuList).Error
		for i := 0; i < len(menuList); i++ {
			err = getBaseChildrenList(&menuList[i])
		}
		return err, menuList, total
	}
}

//获取基础路由树
func (m *SysBaseMenu) GetBaseMenuTree() (err error, menus []SysBaseMenu) {
	err = qmsql.DEFAULTDB.Where(" parent_id = ?", 0).Order("sort", true).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i])
	}
	return err, menus
}

func getBaseChildrenList(menu *SysBaseMenu) (err error) {
	err = qmsql.DEFAULTDB.Where("parent_id = ?", menu.ID).Order("sort", true).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i])
	}
	return err
}
