package internal

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var Authority = new(authority)

type authority struct{}

// Tree .
// tree: map[uint][]model.SysAuthority => map[父角色id][]子角色
// ids: map[uint][]uint => map[角色id][]子角色id
func (a *authority) Tree(ctx context.Context) (tree map[uint][]model.SysAuthority, ids map[uint][]uint, err error) {
	ids = make(map[uint][]uint, 10)
	tree = make(map[uint][]model.SysAuthority, 10)
	var entities []model.SysAuthority
	err = global.GVA_DB.WithContext(ctx).FindInBatches(&entities, 10, func(tx *gorm.DB, batch int) error {
		for i := 0; i < batch; i++ {
			v1, ok := tree[*entities[i].ParentId]
			if !ok {
				v1 = make([]model.SysAuthority, 0)
			}
			v1 = append(v1, entities[i])
			tree[*entities[i].ParentId] = v1
			v2, o2 := ids[entities[i].AuthorityId]
			if !o2 {
				v2 = make([]uint, 0)
			}
			v2 = append(v2, entities[i].AuthorityId)
			ids[entities[i].AuthorityId] = v2
		}
		return nil
	}).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "查询失败!")
	}
	return tree, ids, nil
}

// Children .
// tree: map[uint][]model.SysAuthority => map[父角色id][]子角色
// entity: *model.SysAuthority => 角色
func (a *authority) Children(tree map[uint][]model.SysAuthority, entity *model.SysAuthority) {
	value, ok := tree[entity.AuthorityId]
	if ok {
		entity.Children = value
		for i := 0; i < len(entity.Children); i++ {
			a.Children(tree, &value[i])
		}
	}
}

// Copy 复制角色
// ids: map[uint][]uint => map[角色id][]子角色id
// creates: map[uint]uint => map[原角色id]新角色id
// entity: *model.SysAuthority => 原角色
// parent: model.SysAuthority => 新角色
func (a *authority) Copy(ids map[uint][]uint, creates map[uint]uint, entity *model.SysAuthority, parent model.SysAuthority) []model.SysAuthority {
	var entities []model.SysAuthority
	a.CopyChildren(ids, creates, entity, parent, &entities, 1)
	return entities
}

// CopyChildren 复制角色的子角色
// ids: map[uint][]uint => map[角色id][]子角色id
// creates: map[uint]uint => map[原角色id]新角色id
// entity: *model.SysAuthority => 原角色
// parent: model.SysAuthority => 新角色
// entities: *[]model.SysAuthority => 新角色集合
// zi: uint => 子角色id
func (a *authority) CopyChildren(ids map[uint][]uint, creates map[uint]uint, entity *model.SysAuthority, parent model.SysAuthority, entities *[]model.SysAuthority, zi uint) {
	var index int
	for i := 0; i < len(entity.Children); i++ {
	again:
		id := parent.AuthorityId*10 + uint(i+index)
		_, ok := ids[id]
		if ok {
			index++
			goto again
		}
		ids[id] = []uint{}
		names := strings.Split(parent.AuthorityName, "子")
		if len(names) == 0 {
			return
		}
		name := names[0]
		for j := 0; j < int(zi); j++ {
			name += "子"
		}
		name += fmt.Sprintf("角色%d", i+1)
		children := model.SysAuthority{
			AuthorityId:   id,
			AuthorityName: name,
			DefaultRouter: entity.Children[i].DefaultRouter,
		}
		creates[entity.Children[i].AuthorityId] = children.AuthorityId
		*entities = append(*entities, children)
		a.CopyChildren(ids, creates, &entity.Children[i], children, entities, zi+1)
	}
}

// Menus .
// authorityMenus: map[uint][]model.SysBaseMenu => map[角色id][]菜单
func (a *authority) Menus(ctx context.Context) (authorityMenus map[uint][]model.SysBaseMenu, err error) {
	menus := make(map[uint]model.SysBaseMenu, 10)           // map[MenuId][]model.SysBaseMenu
	authoritiesMenus := make(map[uint][]uint, 10)           // map[AuthorityId][]MenuId
	authorityMenus = make(map[uint][]model.SysBaseMenu, 10) // map[AuthorityId][]model.SysBaseMenu
	{
		var entities []model.SysBaseMenu
		err = global.GVA_DB.WithContext(ctx).FindInBatches(&entities, 10, func(tx *gorm.DB, batch int) error {
			for i := 0; i < batch; i++ {
				_, ok := menus[entities[i].ID]
				if ok {
					return errors.New("存在相同菜单id")
				}
				menus[entities[i].ID] = entities[i]
			}
			return nil
		}).Error
		if err != nil {
			return nil, errors.Wrap(err, "获取所有菜单失败!")
		}
	} // 获取所有菜单
	{
		authoritiesMenus, err = a.MenuIds(ctx)
		if err != nil {
			return nil, err
		}
	} // 获取角色对应的菜单
	{
		for key, values := range authoritiesMenus {
			v1, o1 := authorityMenus[key]
			if !o1 {
				v1 = make([]model.SysBaseMenu, 0)
			}
			for i := 0; i < len(values); i++ {
				v2, o2 := menus[values[i]]
				if o2 {
					v1 = append(v1, v2)
				}
			}
			authorityMenus[key] = v1
		}
	} // 组装tree
	return authorityMenus, nil
}

// MenuIds .
// authorityMenus: map[uint][]uint => map[角色id][]菜单id
func (a *authority) MenuIds(ctx context.Context) (authorityMenus map[uint][]uint, err error) {
	authorityMenus = make(map[uint][]uint, 10)
	{
		var entities []model.SysAuthorityMenu
		err = global.GVA_DB.WithContext(ctx).FindInBatches(&entities, 10, func(tx *gorm.DB, batch int) error {
			for i := 0; i < batch; i++ {
				menuId, _ := strconv.ParseUint(entities[i].MenuId, 10, 64)
				authorityId, _ := strconv.ParseUint(entities[i].AuthorityId, 10, 64)
				value, ok := authorityMenus[uint(authorityId)]
				if !ok {
					value = make([]uint, 0)
				}
				value = append(value, uint(menuId))
				authorityMenus[uint(authorityId)] = value
			}
			return nil
		}).Error
		if err != nil {
			return nil, errors.Wrap(err, "查询角色关联菜单关联表失败!")
		}
	}
	return authorityMenus, nil
}

// CopyMenus 复制角色的角色菜单
// creates: map[uint]uint => map[原角色id]新角色id
func (a *authority) CopyMenus(ctx context.Context, creates map[uint]uint) ([]model.SysAuthorityMenu, error) {
	menus, err := a.MenuIds(ctx)
	if err != nil {
		return nil, err
	}
	entities := make([]model.SysAuthorityMenu, len(creates)*len(menus))
	for key, value := range creates {
		menu, ok := menus[key]
		if ok {
			for j := 0; j < len(menu); j++ {
				menuId := strconv.FormatUint(uint64(menu[j]), 10)
				authorityId := strconv.FormatUint(uint64(value), 10)
				entities = append(entities, model.SysAuthorityMenu{
					MenuId:      menuId,
					AuthorityId: authorityId,
				})
			}
		}
	}
	return entities, nil
}

func (a *authority) MenuButtons(ctx context.Context) (authorityMenuButtons map[uint][]model.SysBaseMenuBtn, err error) {
	menuButtons := make(map[uint]model.SysBaseMenuBtn, 10)           // map[MenuButtonId]model.SysAuthorityBtn
	authoritiesMenuButtons := make(map[uint][]uint, 10)              // map[AuthorityId][]MenuButtonId
	authorityMenuButtons = make(map[uint][]model.SysBaseMenuBtn, 10) // map[AuthorityId][]model.SysAuthorityBtn
	{
		var entities []model.SysBaseMenuBtn
		err = global.GVA_DB.WithContext(ctx).FindInBatches(&entities, 10, func(tx *gorm.DB, batch int) error {
			for i := 0; i < batch; i++ {
				_, ok := menuButtons[entities[i].ID]
				if ok {
					return errors.New("存在相同菜单按钮id")
				}
				menuButtons[entities[i].ID] = entities[i]
			}
			return nil
		}).Error
	} // 获取所有菜单按钮
	{
		authoritiesMenuButtons, err = a.MenuButtonIds(ctx)
		if err != nil {
			return nil, err
		}
	} // 获取角色对应的菜单按钮
	{
		for key, values := range authoritiesMenuButtons {
			v1, o1 := authorityMenuButtons[key]
			if !o1 {
				v1 = make([]model.SysBaseMenuBtn, 0)
			}
			for i := 0; i < len(values); i++ {
				v2, o2 := menuButtons[values[i]]
				if o2 {
					v1 = append(v1, v2)
				}
			}
			authorityMenuButtons[key] = v1
		}
	} // 组装 authorityMenuButtons
	return authorityMenuButtons, nil
}

func (a *authority) MenuButtonIds(ctx context.Context) (authorityMenuButtons map[uint][]uint, err error) {
	authorityMenuButtons = make(map[uint][]uint, 10)
	{
		var entities []model.SysAuthorityBtn
		err = global.GVA_DB.WithContext(ctx).FindInBatches(&entities, 10, func(tx *gorm.DB, batch int) error {
			for i := 0; i < batch; i++ {
				value, ok := authorityMenuButtons[entities[i].AuthorityId]
				if !ok {
					value = make([]uint, 0)
				}
				value = append(value, entities[i].SysBaseMenuBtnID)
				authorityMenuButtons[entities[i].AuthorityId] = value
			}
			return nil
		}).Error
		if err != nil {
			return nil, errors.Wrap(err, "查询角色关联菜单按钮关联表失败!")
		}
	}
	return authorityMenuButtons, nil
}

// CopyMenuButtons 复制角色的角色菜单按钮
// creates: map[uint]uint => map[原角色id]新角色id
func (a *authority) CopyMenuButtons(ctx context.Context, creates map[uint]uint) ([]model.SysAuthorityBtn, error) {
	menuButtons, err := a.MenuButtons(ctx)
	if err != nil {
		return nil, err
	}
	entities := make([]model.SysAuthorityBtn, len(creates)*len(menuButtons))
	for key, value := range creates {
		menuButton, ok := menuButtons[key]
		if ok {
			for i := 0; i < len(menuButton); i++ {
				entities = append(entities, model.SysAuthorityBtn{
					AuthorityId:      value,
					SysMenuID:        menuButton[i].SysBaseMenuID,
					SysBaseMenuBtnID: menuButton[i].ID,
				})
			}
		}
	}
	return entities, nil
}

// CopyCasbin .
// creates: map[uint]uint => map[原角色id]新角色id
func (a *authority) CopyCasbin(rules map[uint][]string, creates map[uint]uint) [][]string {
	entities := make([][]string, len(rules)*len(creates))
	for key, value := range creates {
		values, ok := rules[key]
		if ok {
			if len(values) >= 4 {
				values[1] = strconv.Itoa(int(value))
				entities = append(entities, values)
			}
		}
	}
	return entities
}
