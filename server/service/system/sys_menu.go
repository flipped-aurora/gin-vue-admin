package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: treeMap map[string][]system.SysMenu, err error

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[uint][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	var btns []system.SysAuthorityBtn
	treeMap = make(map[uint][]system.SysMenu)

	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}

	err = global.GVA_DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}

	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      baseMenu[i].ID,
			Parameters:  baseMenu[i].Parameters,
		})
	}

	err = global.GVA_DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]uint)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.SysBaseMenu.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: menus []system.SysMenu, err error

func (menuService *MenuService) GetMenuTree(authorityId uint) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.SysMenu, treeMap map[string][]model.SysMenu
//@return: err error

func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[uint][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetInfoList
//@description: 获取路由分页
//@return: list interface{}, total int64,err error

func (menuService *MenuService) GetInfoList(authorityID uint) (list interface{}, err error) {
	var menuList []system.SysBaseMenu
	treeMap, err := menuService.getBaseMenuTreeMap(authorityID)
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[uint][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.SysBaseMenu
//@return: error

func (menuService *MenuService) AddBaseMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: treeMap map[string][]system.SysBaseMenu, err error

func (menuService *MenuService) getBaseMenuTreeMap(authorityID uint) (treeMap map[uint][]system.SysBaseMenu, err error) {
	parentAuthorityID, err := AuthorityServiceApp.GetParentAuthorityID(authorityID)
	if err != nil {
		return nil, err
	}

	var allMenus []system.SysBaseMenu
	treeMap = make(map[uint][]system.SysBaseMenu)
	db := global.GVA_DB.Order("sort").Preload("MenuBtn").Preload("Parameters")

	// 当开启了严格的树角色并且父角色不为0时需要进行菜单筛选
	if global.GVA_CONFIG.System.UseStrictAuth && parentAuthorityID != 0 {
		var authorityMenus []system.SysAuthorityMenu
		err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityID).Find(&authorityMenus).Error
		if err != nil {
			return nil, err
		}
		var menuIds []string
		for i := range authorityMenus {
			menuIds = append(menuIds, authorityMenus[i].MenuId)
		}
		db = db.Where("id in (?)", menuIds)
	}

	err = db.Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: menus []system.SysBaseMenu, err error

func (menuService *MenuService) GetBaseMenuTree(authorityID uint) (menus []system.SysBaseMenu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap(authorityID)
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error

func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, adminAuthorityID, authorityId uint) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus

	err = AuthorityServiceApp.CheckAuthorityIDAuth(adminAuthorityID, authorityId)
	if err != nil {
		return err
	}

	var authority system.SysAuthority
	_ = global.GVA_DB.First(&authority, "authority_id = ?", adminAuthorityID).Error
	var menuIds []string

	// 当开启了严格的树角色并且父角色不为0时需要进行菜单筛选
	if global.GVA_CONFIG.System.UseStrictAuth && *authority.ParentId != 0 {
		var authorityMenus []system.SysAuthorityMenu
		err = global.GVA_DB.Where("sys_authority_authority_id = ?", adminAuthorityID).Find(&authorityMenus).Error
		if err != nil {
			return err
		}
		for i := range authorityMenus {
			menuIds = append(menuIds, authorityMenus[i].MenuId)
		}

		for i := range menus {
			hasMenu := false
			for j := range menuIds {
				idStr := strconv.Itoa(int(menus[i].ID))
				if idStr == menuIds[j] {
					hasMenu = true
				}
			}
			if !hasMenu {
				return errors.New("添加失败,请勿跨级操作")
			}
		}
	}

	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: menus []system.SysMenu, err error

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (menus []system.SysMenu, err error) {
	var baseMenu []system.SysBaseMenu
	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", info.AuthorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}

	err = global.GVA_DB.Where("id in (?) ", MenuIds).Order("sort").Find(&baseMenu).Error

	for i := range baseMenu {
		menus = append(menus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: info.AuthorityId,
			MenuId:      baseMenu[i].ID,
			Parameters:  baseMenu[i].Parameters,
		})
	}
	return menus, err
}

// UserAuthorityDefaultRouter 用户角色默认路由检查
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (menuService *MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.GVA_DB.Model(&system.SysAuthorityMenu{}).Where("sys_authority_authority_id = ?", user.AuthorityId).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am system.SysBaseMenu
	err = global.GVA_DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
