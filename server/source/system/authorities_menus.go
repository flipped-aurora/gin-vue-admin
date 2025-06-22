package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenuAuthority = initOrderMenu + initOrderAuthority

type initMenuAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuAuthority, &initMenuAuthority{})
}

func (i *initMenuAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuAuthority) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i *initMenuAuthority) InitializerName() string {
	return "sys_menu_authorities"
}

func (i *initMenuAuthority) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	initAuth := &initAuthority{}
	authorities, ok := ctx.Value(initAuth.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext, "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}

	allMenus, ok := ctx.Value(new(initMenu).InitializerName()).([]sysModel.SysBaseMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx

	// 构建菜单ID映射，方便快速查找
	menuMap := make(map[uint]sysModel.SysBaseMenu)
	for _, menu := range allMenus {
		menuMap[menu.ID] = menu
	}

	// 为不同角色分配不同权限
	// 1. 超级管理员角色(888) - 拥有所有菜单权限
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Replace(allMenus); err != nil {
		return next, errors.Wrap(err, "为超级管理员分配菜单失败")
	}

	// 2. 普通用户角色(8881) - 仅拥有基础功能菜单
	// 仅选择部分父级菜单及其子菜单
	var menu8881 []sysModel.SysBaseMenu

	// 添加仪表盘、关于我们和个人信息菜单
	for _, menu := range allMenus {
		if menu.ParentId == 0 && (menu.Name == "dashboard" || menu.Name == "about" || menu.Name == "person" || menu.Name == "state") {
			menu8881 = append(menu8881, menu)
		}
	}

	if err = db.Model(&authorities[1]).Association("SysBaseMenus").Replace(menu8881); err != nil {
		return next, errors.Wrap(err, "为普通用户分配菜单失败")
	}

	// 3. 测试角色(9528) - 拥有部分菜单权限
	var menu9528 []sysModel.SysBaseMenu

	// 添加所有父级菜单
	for _, menu := range allMenus {
		if menu.ParentId == 0 {
			menu9528 = append(menu9528, menu)
		}
	}

	// 添加部分子菜单 - 系统工具、示例文件等模块的子菜单
	for _, menu := range allMenus {
		parentName := ""
		if menu.ParentId > 0 && menuMap[menu.ParentId].Name != "" {
			parentName = menuMap[menu.ParentId].Name
		}

		if menu.ParentId > 0 && (parentName == "systemTools" || parentName == "example") {
			menu9528 = append(menu9528, menu)
		}
	}

	if err = db.Model(&authorities[2]).Association("SysBaseMenus").Replace(menu9528); err != nil {
		return next, errors.Wrap(err, "为测试角色分配菜单失败")
	}

	return next, nil
}

func (i *initMenuAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	auth := &sysModel.SysAuthority{}
	if ret := db.Model(auth).
		Where("authority_id = ?", 9528).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(auth.SysBaseMenus) > 0
	}
	return false
}
