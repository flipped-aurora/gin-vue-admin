package model

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SysBaseMenu struct {
	gorm.Model
	MenuLevel     uint   `json:"-"`
	ParentId      string `json:"parentId"`
	Path          string `json:"path"`
	Name          string `json:"name"`
	Hidden        bool   `json:"hidden"`
	Component     string `json:"component"`
	Sort          string `json:"sort"`
	Meta          `json:"meta"`
	NickName      string         `json:"nickName"`
	SysAuthoritys []SysAuthority `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu  `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

//增加基础路由
func (b *SysBaseMenu) AddBaseMenu() (err error) {
	findOne := global.GVA_DB.Where("name = ?", b.Name).Find(&SysBaseMenu{}).Error
	if findOne != nil {
		b.NickName = b.Title
		err = global.GVA_DB.Create(b).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

//删除基础路由
func (b *SysBaseMenu) DeleteBaseMenu(id float64) (err error) {
	err = global.GVA_DB.Where("parent_id = ?", id).First(&SysBaseMenu{}).Error
	if err != nil {
		err = global.GVA_DB.Where("id = ?", id).Delete(&b).Error
		err = global.GVA_DB.Where("menu_id = ?", id).Unscoped().Delete(&SysMenu{}).Error
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//更新路由
func (b *SysBaseMenu) UpdateBaseMenu() (err error) {
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = b.ParentId
	upDateMap["path"] = b.Path
	upDateMap["name"] = b.Name
	upDateMap["hidden"] = b.Hidden
	upDateMap["component"] = b.Component
	upDateMap["title"] = b.Title
	upDateMap["icon"] = b.Icon
	upDateMap["sort"] = b.Sort
	upDateMap["nick_name"] = b.Title
	err = global.GVA_DB.Where("id = ?", b.ID).Find(&SysBaseMenu{}).Updates(upDateMap).Error
	err1 := global.GVA_DB.Where("menu_id = ?", b.ID).Find(&[]SysMenu{}).Updates(upDateMap).Error
	fmt.Printf("菜单修改时候，关联菜单err1:%v,err:%v", err1, err)
	return err
}

//当前选中角色所拥有的路由
func (b *SysBaseMenu) GetBaseMenuById(id float64) (err error, menu SysBaseMenu) {
	err = global.GVA_DB.Where("id = ?", id).First(&menu).Error
	return
}

//获取路由分页
func (b *SysBaseMenu) GetInfoList(info PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	if err != nil {
		return
	} else {
		var menuList []SysBaseMenu
		err = db.Limit(limit).Offset(offset).Where("parent_id = 0").Order("sort", true).Find(&menuList).Error
		for i := 0; i < len(menuList); i++ {
			err = getBaseChildrenList(&menuList[i])
		}
		return err, menuList, total
	}
}

//获取基础路由树
func (m *SysBaseMenu) GetBaseMenuTree() (err error, menus []SysBaseMenu) {
	err = global.GVA_DB.Where(" parent_id = ?", 0).Order("sort", true).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i])
	}
	return err, menus
}

func getBaseChildrenList(menu *SysBaseMenu) (err error) {
	err = global.GVA_DB.Where("parent_id = ?", menu.ID).Order("sort", true).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i])
	}
	return err
}
