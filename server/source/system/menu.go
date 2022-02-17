package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "sys_base_menus"
}

func (m *menu) Initialize() error {
	entities := []system.SysBaseMenu{
		{GVA_MODEL: global.GVA_MODEL{ID: 1}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: global.Translate("system.menu.dashboard"), Icon: "odometer"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 2}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: global.Translate("system.menu.about"), Icon: "info-filled"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 3}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: global.Translate("system.menu.adminTools"), Icon: "user"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 4}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: global.Translate("system.menu.roleManage"), Icon: "avatar"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 5}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: global.Translate("system.menu.menuManage"), Icon: "tickets", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 6}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: global.Translate("system.menu.apiManage"), Icon: "platform", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 7}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: global.Translate("system.menu.userManage"), Icon: "coordinate"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 8}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: global.Translate("system.menu.personalInfo"), Icon: "message"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 9}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: global.Translate("system.menu.examples"), Icon: "management"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 10}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: system.Meta{Title: global.Translate("system.menu.excelImportExport"), Icon: "takeaway-box"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 11}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: system.Meta{Title: global.Translate("system.menu.mediaLibUpDown"), Icon: "upload"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 12}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: system.Meta{Title: global.Translate("system.menu.breakPoint"), Icon: "upload-filled"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 13}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: system.Meta{Title: global.Translate("system.menu.customerList"), Icon: "avatar"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 14}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: system.Meta{Title: global.Translate("system.menu.systemTools"), Icon: "tools"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 15}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: system.Meta{Title: global.Translate("system.menu.autoCode"), Icon: "cpu", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 16}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: system.Meta{Title: global.Translate("system.menu.formCreator"), Icon: "magic-stick", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 17}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: system.Meta{Title: global.Translate("system.menu.system"), Icon: "operation"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 18}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: system.Meta{Title: global.Translate("system.menu.dictMange"), Icon: "notebook"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 19}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: system.Meta{Title: global.Translate("system.menu.dictDetail"), Icon: "order"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 20}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: global.Translate("system.menu.operationLog"), Icon: "pie-chart"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 21}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: system.Meta{Title: global.Translate("system.menu.resumableUpload"), Icon: "upload"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 22}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: system.Meta{Title: global.Translate("system.menu.website"), Icon: "home-filled"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 23}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: global.Translate("system.menu.serverStatus"), Icon: "cloudy"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 24}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: system.Meta{Title: global.Translate("system.menu.autoCodeManage"), Icon: "magic-stick"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 25}, MenuLevel: 0, Hidden: true, ParentId: "14", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: system.Meta{Title: global.Translate("system.menu.autoCodeEdit"), Icon: "magic-stick"}},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+" "+global.Translate("general.tabelDataInitFail"))
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("path = ?", "autoCodeEdit/:id").First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
