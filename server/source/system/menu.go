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

	// 定义所有菜单(顶层分组)
	allMenus := []SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "permission", Name: "permission", Component: "view/routerHolder.vue", Sort: 2, Meta: Meta{Title: "权限管理", Icon: "perm-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "systemConfig", Name: "systemConfig", Component: "view/routerHolder.vue", Sort: 3, Meta: Meta{Title: "系统设置", Icon: "config-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "monitor", Name: "monitor", Component: "view/routerHolder.vue", Sort: 4, Meta: Meta{Title: "运维监控", Icon: "monitor-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "media", Name: "media", Component: "view/routerHolder.vue", Sort: 5, Meta: Meta{Title: "媒体管理", Icon: "folder-opened"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "systemTools", Name: "systemTools", Component: "view/routerHolder.vue", Sort: 6, Meta: Meta{Title: "编程辅助", Icon: "cpu"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "ai", Name: "ai", Component: "view/routerHolder.vue", Sort: 7, Meta: Meta{Title: "AI 工坊", Icon: "ai-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 9, Meta: Meta{Title: "示例文件", Icon: "example-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 8, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 10, Meta: Meta{Title: "关于我们", Icon: "office-building"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 11, Meta: Meta{Title: "官方网站", Icon: "customer-gva"}},
		{MenuLevel: 0, Hidden: true, ParentId: 0, Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 12, Meta: Meta{Title: "个人信息", Icon: "postcard"}},
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
		// 权限管理
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["permission"], Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Icon: "role-gva"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["permission"], Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["permission"], Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Icon: "api-gva", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["permission"], Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Icon: "user"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["permission"], Path: "apiToken", Name: "apiToken", Component: "view/systemTools/apiToken/index.vue", Sort: 5, Meta: Meta{Title: "API Token", Icon: "key"}},

		// 系统设置
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemConfig"], Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 1, Meta: Meta{Title: "配置文件", Icon: "config-file-gva"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemConfig"], Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 2, Meta: Meta{Title: "字典管理", Icon: "notebook"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemConfig"], Path: "sysParams", Name: "sysParams", Component: "view/superAdmin/params/sysParams.vue", Sort: 3, Meta: Meta{Title: "参数管理", Icon: "set-up"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemConfig"], Path: "security", Name: "security", Component: "view/system/security/index.vue", Sort: 4, Meta: Meta{Title: "安全配置", Icon: "security-gva"}},

		// 运维监控
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["monitor"], Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 1, Meta: Meta{Title: "操作历史", Icon: "document"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["monitor"], Path: "loginLog", Name: "loginLog", Component: "view/systemTools/loginLog/index.vue", Sort: 2, Meta: Meta{Title: "登录日志", Icon: "clock"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["monitor"], Path: "sysError", Name: "sysError", Component: "view/systemTools/sysError/sysError.vue", Sort: 3, Meta: Meta{Title: "错误日志", Icon: "error-gva"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["monitor"], Path: "sysVersion", Name: "sysVersion", Component: "view/systemTools/version/version.vue", Sort: 4, Meta: Meta{Title: "版本管理", Icon: "version-gva"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["monitor"], Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 5, Meta: Meta{Title: "服务器状态", Icon: "server"}},

		// 媒体管理
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["media"], Path: "upload", Name: "upload", Component: "view/media/upload.vue", Sort: 1, Meta: Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["media"], Path: "chunkUpload", Name: "chunkUpload", Component: "view/media/chunkUpload.vue", Sort: 2, Meta: Meta{Title: "大文件上传", Icon: "folder-add"}},

		// example子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 1, Meta: Meta{Title: "客户列表（资源示例）", Icon: "service"}},

		// systemTools子菜单(编程辅助)
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoCode", Name: "autoCode", Component: "plugin/auto/view/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "代码生成器", Icon: "magic-stick"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoPkg", Name: "autoPkg", Component: "plugin/auto/view/autoPkg/autoPkg.vue", Sort: 2, Meta: Meta{Title: "模板配置", Icon: "files"}},
		{MenuLevel: 1, Hidden: true, ParentId: menuNameMap["systemTools"], Path: "autoCodeAdmin", Name: "AutoCodeAdmin", Component: "plugin/auto/view/autoCodeAdmin/index.vue", Sort: 3, Meta: Meta{Title: "自动代码管理", Icon: "tools"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "formCreate", Name: "formCreate", Component: "plugin/auto/view/formCreate/index.vue", Sort: 4, Meta: Meta{Title: "表单生成器", Icon: "edit", KeepAlive: true}},
		{MenuLevel: 1, Hidden: true, ParentId: menuNameMap["systemTools"], Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "plugin/auto/view/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "exportTemplate", Name: "exportTemplate", Component: "plugin/auto/view/exportTemplate/exportTemplate.vue", Sort: 6, Meta: Meta{Title: "导出模板", Icon: "download"}},

		// AI 工坊
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["ai"], Path: "mcpTool", Name: "mcpTool", Component: "plugin/ai/view/mcp/mcp.vue", Sort: 1, Meta: Meta{Title: "Mcp Tools模板", Icon: "grid"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["ai"], Path: "mcpTest", Name: "mcpTest", Component: "plugin/ai/view/mcp/mcpTest.vue", Sort: 2, Meta: Meta{Title: "Mcp Tools管理", Icon: "connection"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["ai"], Path: "skills", Name: "Skills", Component: "plugin/ai/view/skills/index.vue", Sort: 3, Meta: Meta{Title: "Skills管理", Icon: "edit-pen"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["ai"], Path: "cli", Name: "Cli", Component: "plugin/ai/view/cli/index.vue", Sort: 4, Meta: Meta{Title: "AI CLI管理", Icon: "monitor", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["ai"], Path: "picture", Name: "picture", Component: "plugin/ai/view/picture/picture.vue", Sort: 5, Meta: Meta{Title: "AI页面绘制", Icon: "picture"}},

		// 插件系统
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "https://plugin.gin-vue-admin.com/", Name: "https://plugin.gin-vue-admin.com/", Component: "https://plugin.gin-vue-admin.com/", Sort: 0, Meta: Meta{Title: "插件市场", Icon: "shop"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "pubPlug", Name: "pubPlug", Component: "view/systemTools/pubPlug/pubPlug.vue", Sort: 3, Meta: Meta{Title: "打包插件", Icon: "suitcase"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 4, Meta: Meta{Title: "邮件插件", Icon: "message"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "anInfo", Name: "anInfo", Component: "plugin/announcement/view/info.vue", Sort: 5, Meta: Meta{Title: "公告管理[示例]", Icon: "bell"}},
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
	if errors.Is(db.Where("path = ?", "dashboard").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
