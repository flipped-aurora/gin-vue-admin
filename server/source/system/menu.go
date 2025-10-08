package system

import (
	"context"

	. "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i *initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		&SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{}) &&
		m.HasTable(&SysBaseMenuBtn{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 定义所有菜单
	allMenus := []SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 9, Meta: Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: true, ParentId: 0, Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 7, Meta: Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: Meta{Title: "系统工具", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: Meta{Title: "官方网站", Icon: "customer-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 6, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
	}

	// 先创建父级菜单（ParentId = 0 的菜单）
	if err = db.Create(&allMenus).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"父级菜单初始化失败!")
	}

	// 建立菜单映射 - 通过Name查找已创建的菜单及其ID
	menuNameMap := make(map[string]uint)
	for _, menu := range allMenus {
		menuNameMap[menu.Name] = menu.ID
	}

	// 定义子菜单，并设置正确的ParentId
	childMenus := []SysBaseMenu{
		// superAdmin子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: Meta{Title: "字典管理", Icon: "notebook"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: Meta{Title: "操作历史", Icon: "pie-chart"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "sysParams", Name: "sysParams", Component: "view/superAdmin/params/sysParams.vue", Sort: 7, Meta: Meta{Title: "参数管理", Icon: "compass"}},

		// example子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: Meta{Title: "断点续传", Icon: "upload-filled"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: Meta{Title: "客户列表（资源示例）", Icon: "avatar"}},

		// systemTools子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 3, Meta: Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 4, Meta: Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 2, Meta: Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		{MenuLevel: 1, Hidden: true, ParentId: menuNameMap["systemTools"], Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoPkg", Name: "autoPkg", Component: "view/systemTools/autoPkg/autoPkg.vue", Sort: 0, Meta: Meta{Title: "模板配置", Icon: "folder"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "exportTemplate", Name: "exportTemplate", Component: "view/systemTools/exportTemplate/exportTemplate.vue", Sort: 5, Meta: Meta{Title: "导出模板", Icon: "reading"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "picture", Name: "picture", Component: "view/systemTools/autoCode/picture.vue", Sort: 6, Meta: Meta{Title: "AI页面绘制", Icon: "picture-filled"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "mcpTool", Name: "mcpTool", Component: "view/systemTools/autoCode/mcp.vue", Sort: 7, Meta: Meta{Title: "Mcp Tools模板", Icon: "magnet"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "mcpTest", Name: "mcpTest", Component: "view/systemTools/autoCode/mcpTest.vue", Sort: 7, Meta: Meta{Title: "Mcp Tools测试", Icon: "partly-cloudy"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "sysVersion", Name: "sysVersion", Component: "view/systemTools/version/version.vue", Sort: 8, Meta: Meta{Title: "版本管理", Icon: "server"}},

		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "https://plugin.gin-vue-admin.com/", Name: "https://plugin.gin-vue-admin.com/", Component: "https://plugin.gin-vue-admin.com/", Sort: 0, Meta: Meta{Title: "插件市场", Icon: "shop"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "pubPlug", Name: "pubPlug", Component: "view/systemTools/pubPlug/pubPlug.vue", Sort: 3, Meta: Meta{Title: "打包插件", Icon: "files"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 4, Meta: Meta{Title: "邮件插件", Icon: "message"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "anInfo", Name: "anInfo", Component: "plugin/announcement/view/info.vue", Sort: 5, Meta: Meta{Title: "公告管理[示例]", Icon: "scaleToOriginal"}},
	}

	// 创建子菜单
	if err = db.Create(&childMenus).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"子菜单初始化失败!")
	}

	// 组合所有菜单作为返回结果
	allEntities := append(allMenus, childMenus...)
	next = context.WithValue(ctx, i.InitializerName(), allEntities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "autoPkg").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
