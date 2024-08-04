package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
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

func (i initMenu) InitializerName() string {
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
	entities := []SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "system.menu.dashboard", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 9, Meta: Meta{Title: "system.menu.about", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: Meta{Title: "system.menu.adminTools", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "system.menu.roleManage", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "system.menu.menuManage", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "system.menu.apiManage", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "system.menu.userManage", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: Meta{Title: "system.menu.dictManage", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: Meta{Title: "system.menu.operationLog", Icon: "pie-chart"}},
		{MenuLevel: 0, Hidden: true, ParentId: 0, Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: Meta{Title: "system.menu.personalInfo", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 7, Meta: Meta{Title: "system.menu.examples", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: 11, Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: Meta{Title: "system.menu.mediaLibUpDown", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: 11, Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: Meta{Title: "system.menu.breakPoint", Icon: "upload-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: 11, Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: Meta{Title: "system.menu.customerList", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: Meta{Title: "system.menu.systemTools", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: 15, Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "system.menu.autoCode", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 15, Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 3, Meta: Meta{Title: "system.menu.formCreator", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 15, Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 4, Meta: Meta{Title: "system.menu.system", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: 15, Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 2, Meta: Meta{Title: "system.menu.autoCodeManage", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: true, ParentId: 15, Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "system.menu.autoCodeEdit" + " - ${id}", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: 15, Path: "autoPkg", Name: "autoPkg", Component: "view/systemTools/autoPkg/autoPkg.vue", Sort: 0, Meta: Meta{Title: "system.menu.templateConfig", Icon: "folder"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: Meta{Title: "system.menu.website", Icon: "customer-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "system.menu.serverStatus", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 6, Meta: Meta{Title: "system.menu.pluginSystem", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: 24, Path: "https://plugin.gin-vue-admin.com/", Name: "https://plugin.gin-vue-admin.com/", Component: "https://plugin.gin-vue-admin.com/", Sort: 0, Meta: Meta{Title: "system.menu.pluginMarket", Icon: "shop"}},
		{MenuLevel: 0, Hidden: false, ParentId: 24, Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "system.menu.pluginInstall", Icon: "box"}},
		{MenuLevel: 0, Hidden: false, ParentId: 24, Path: "pubPlug", Name: "pubPlug", Component: "view/systemTools/pubPlug/pubPlug.vue", Sort: 3, Meta: Meta{Title: "system.menu.packagePlugin", Icon: "files"}},
		{MenuLevel: 0, Hidden: false, ParentId: 24, Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 4, Meta: Meta{Title: "system.menu.emailPlugin", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: 15, Path: "exportTemplate", Name: "exportTemplate", Component: "view/systemTools/exportTemplate/exportTemplate.vue", Sort: 5, Meta: Meta{Title: "system.menu.tableTemplate", Icon: "reading"}},
		{MenuLevel: 0, Hidden: false, ParentId: 24, Path: "anInfo", Name: "anInfo", Component: "plugin/announcement/view/info.vue", Sort: 5, Meta: Meta{Title: "system.menu.announcementManage", Icon: "scaleToOriginal"}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+" "+global.Translate("general.tabelDataInitFail"))
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
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
