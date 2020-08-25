package model

import (
	"gorm.io/gorm"
	"time"
)

type SysBaseMenu struct {
	gorm.Model
	MenuLevel     uint   `json:"-"`
	ParentId      string `json:"parentId" gorm:"comment:'父菜单ID'"`
	Path          string `json:"path" gorm:"comment:'路由path'"`
	Name          string `json:"name" gorm:"comment:'路由name'"`
	Hidden        bool   `json:"hidden" gorm:"comment:'是否在列表隐藏'"`
	Component     string `json:"component" gorm:"comment:'对应前端文件路径'"`
	Sort          int    `json:"sort" gorm:"comment:'排序标记'"`
	Meta          `json:"meta" gorm:"comment:'附加属性'"`
	SysAuthoritys []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter `json:"parameters"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:'是否缓存'"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:'是否是基础路由（开发中）'"`
	Title       string `json:"title" gorm:"comment:'菜单名'"`
	Icon        string `json:"icon" gorm:"comment:'菜单图标'"`
}

type SysBaseMenuParameter struct {
	gorm.Model
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"commit:'地址栏携带参数为params还是query'"`
	Key           string `json:"key" gorm:"commit:'地址栏携带参数的key'"`
	Value         string `json:"value" gorm:"commit:'地址栏携带参数的值'"`
}

func SysBaseMenusData() []SysBaseMenu {
	return []SysBaseMenu{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "setting"}},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7,Meta:Meta{Title: "关于我们", Icon: "info"}},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3,Meta:Meta{Title: "超级管理员", Icon: "user-solid"}},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1,Meta:Meta{Title: "角色管理", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2,Meta:Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3,Meta:Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4,Meta:Meta{Title: "用户管理", Icon: "coordinate"}},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4,Meta:Meta{Title: "个人信息", Icon: "message-solid"}},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()},  MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6,Meta:Meta{Title: "示例文件", Icon: "s-management"}},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "table", Name: "table", Component: "view/example/table/table.vue", Sort: 1,Meta:Meta{Title: "表格示例", Icon: "s-order"}},
		{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "form", Name: "form", Component: "view/example/form/form.vue", Sort: 2,Meta:Meta{Title: "表单示例", Icon: "document"}},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "rte", Name: "rte", Component: "view/example/rte/rte.vue", Sort: 3,Meta:Meta{Title: "富文本编辑器", Icon: "reading"}},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4,Meta:Meta{Title: "excel导入导出", Icon: "s-marketing"}},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5,Meta:Meta{Title: "上传下载", Icon: "upload"}},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6,Meta:Meta{Title: "断点续传", Icon: "upload"}},
		{Model: gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7,Meta:Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5,Meta:Meta{Title: "系统工具", Icon: "s-cooperation"}},
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17",Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1,Meta:Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17",Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2,Meta:Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17",Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3,Meta:Meta{Title: "系统配置", Icon: "s-operation"}},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "iconList", Name: "iconList", Component: "view/iconList/index.vue", Sort: 2,Meta:Meta{Title: "图标集合", Icon: "star-on"}},
		{Model: gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5,Meta:Meta{Title: "字典管理", Icon: "notebook-2"}},
		{Model: gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1,Meta:Meta{Title: "字典详情", Icon: "s-order"}},
		{Model: gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6,Meta:Meta{Title: "操作历史", Icon: "time"}},
		{Model: gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6,Meta:Meta{Title: "断点续传（插件版）", Icon: "upload"}},
	}
}
