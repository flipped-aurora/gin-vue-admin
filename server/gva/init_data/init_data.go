package init_data

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
}

type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string
	DataAuthorityIdAuthorityId string
}

func InitSysApi() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysApi{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", "用户登录", "base", "POST"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/register", "用户注册", "base", "POST"},
		{gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", "创建api", "api", "POST"},
		{gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", "获取api列表", "api", "POST"},
		{gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", "获取api详细信息", "api", "POST"},
		{gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", "删除Api", "api", "POST"},
		{gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", "更新Api", "api", "POST"},
		{gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", "获取所有api", "api", "POST"},
		{gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", "创建角色", "authority", "POST"},
		{gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", "删除角色", "authority", "POST"},
		{gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", "获取角色列表", "authority", "POST"},
		{gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", "获取菜单树", "menu", "POST"},
		{gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", "分页获取基础menu列表", "menu", "POST"},
		{gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", "新增菜单", "menu", "POST"},
		{gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", "获取用户动态路由", "menu", "POST"},
		{gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", "增加menu和角色关联关系", "menu", "POST"},
		{gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", "获取指定角色menu", "menu", "POST"},
		{gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", "删除菜单", "menu", "POST"},
		{gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", "更新菜单", "menu", "POST"},
		{gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", "根据id获取菜单", "menu", "POST"},
		{gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", "修改密码", "user", "POST"},
		{gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", "获取用户列表", "user", "POST"},
		{gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", "修改用户角色", "user", "POST"},
		{gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/upload", "文件上传示例", "fileUploadAndDownload", "POST"},
		{gorm.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/getFileList", "获取上传文件列表", "fileUploadAndDownload", "POST"},
		{gorm.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", "更改角色api权限", "casbin", "POST"},
		{gorm.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", "获取权限列表", "casbin", "POST"},
		{gorm.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/deleteFile", "删除文件", "fileUploadAndDownload", "POST"},
		{gorm.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/jwt/jsonInBlacklist", "jwt加入黑名单", "jwt", "POST"},
		{gorm.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", "设置角色资源权限", "authority", "POST"},
		{gorm.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", "获取配置文件内容", "system", "POST"},
		{gorm.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", "设置配置文件内容", "system", "POST"},
		{gorm.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "创建客户", "customer", "POST"},
		{gorm.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "更新客户", "customer", "PUT"},
		{gorm.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "删除客户", "customer", "DELETE"},
		{gorm.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "获取单一客户", "customer", "GET"},
		{gorm.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", "获取客户列表", "customer", "GET"},
		{gorm.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", "RESTFUL模式测试", "casbin", "GET"},
		{gorm.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/createTemp", "自动化代码", "autoCode", "POST"},
		{gorm.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", "更新角色信息", "authority", "PUT"},
		{gorm.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", "拷贝角色", "authority", "POST"},
		{gorm.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", "删除用户", "user", "DELETE"},
		{gorm.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/createSysDictionaryDetail", "新增字典内容", "sysDictionaryDetail", "POST"},
		{gorm.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/deleteSysDictionaryDetail", "删除字典内容", "sysDictionaryDetail", "DELETE"},
		{gorm.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/updateSysDictionaryDetail", "更新字典内容", "sysDictionaryDetail", "PUT"},
		{gorm.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/findSysDictionaryDetail", "根据ID获取字典内容", "sysDictionaryDetail", "GET"},
		{gorm.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/getSysDictionaryDetailList", "获取字典内容列表", "sysDictionaryDetail", "GET"},
		{gorm.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/createSysDictionary", "新增字典", "sysDictionary", "POST"},
		{gorm.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/deleteSysDictionary", "删除字典", "sysDictionary", "DELETE"},
		{gorm.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/updateSysDictionary", "更新字典", "sysDictionary", "PUT"},
		{gorm.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/findSysDictionary", "根据ID获取字典", "sysDictionary", "GET"},
		{gorm.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/getSysDictionaryList", "获取字典列表", "sysDictionary", "GET"},
		{gorm.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", "新增操作记录", "sysOperationRecord", "POST"},
		{gorm.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", "删除操作记录", "sysOperationRecord", "DELETE"},
		{gorm.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", "根据ID获取操作记录", "sysOperationRecord", "GET"},
		{gorm.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", "获取操作记录列表", "sysOperationRecord", "GET"},
		{gorm.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getTables", "获取数据库表", "autoCode", "GET"},
		{gorm.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getDB", "获取所有数据库", "autoCode", "GET"},
		{gorm.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getColume", "获取所选table的所有字段", "autoCode", "GET"},
		{gorm.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", "批量删除操作历史", "sysOperationRecord", "DELETE"},
		{gorm.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/upload", "插件版分片上传", "simpleUploader", "POST"},
		{gorm.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/checkFileMd5", "文件完整度验证", "simpleUploader", "GET"},
		{gorm.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/mergeFileMd5", "上传完成合并文件", "simpleUploader", "GET"},
		{gorm.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", "设置用户信息", "user", "PUT"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysUser() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysUser{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/1571627762timg.jpg", AuthorityId: "888"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", NickName: "QMPlusUser", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitExaCustomer() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.ExaCustomer{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CustomerName: "测试客户", CustomerPhoneData: "1761111111", SysUserID: 1, SysUserAuthorityID: "888"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitCasbinModel() (err error) {
	if !global.GVA_DB.Migrator().HasTable("casbin_rule") {
		if err := global.GVA_DB.Migrator().CreateTable(&gormadapter.CasbinRule{}); err != nil {
			return err
		}
	}
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.CasbinModel{
		{"p", "888", "/base/login", "POST"},
		{"p", "888", "/base/register", "POST"},
		{"p", "888", "/api/createApi", "POST"},
		{"p", "888", "/api/getApiList", "POST"},
		{"p", "888", "/api/getApiById", "POST"},
		{"p", "888", "/api/deleteApi", "POST"},
		{"p", "888", "/api/updateApi", "POST"},
		{"p", "888", "/api/getAllApis", "POST"},
		{"p", "888", "/authority/createAuthority", "POST"},
		{"p", "888", "/authority/deleteAuthority", "POST"},
		{"p", "888", "/authority/getAuthorityList", "POST"},
		{"p", "888", "/authority/setDataAuthority", "POST"},
		{"p", "888", "/authority/updateAuthority", "PUT"},
		{"p", "888", "/authority/copyAuthority", "POST"},
		{"p", "888", "/menu/getMenu", "POST"},
		{"p", "888", "/menu/getMenuList", "POST"},
		{"p", "888", "/menu/addBaseMenu", "POST"},
		{"p", "888", "/menu/getBaseMenuTree", "POST"},
		{"p", "888", "/menu/addMenuAuthority", "POST"},
		{"p", "888", "/menu/getMenuAuthority", "POST"},
		{"p", "888", "/menu/deleteBaseMenu", "POST"},
		{"p", "888", "/menu/updateBaseMenu", "POST"},
		{"p", "888", "/menu/getBaseMenuById", "POST"},
		{"p", "888", "/user/changePassword", "POST"},
		{"p", "888", "/user/getUserList", "POST"},
		{"p", "888", "/user/setUserAuthority", "POST"},
		{"p", "888", "/user/deleteUser", "DELETE"},
		{"p", "888", "/fileUploadAndDownload/upload", "POST"},
		{"p", "888", "/fileUploadAndDownload/getFileList", "POST"},
		{"p", "888", "/fileUploadAndDownload/deleteFile", "POST"},
		{"p", "888", "/casbin/updateCasbin", "POST"},
		{"p", "888", "/casbin/getPolicyPathByAuthorityId", "POST"},
		{"p", "888", "/casbin/casbinTest/:pathParam", "GET"},
		{"p", "888", "/jwt/jsonInBlacklist", "POST"},
		{"p", "888", "/system/getSystemConfig", "POST"},
		{"p", "888", "/system/setSystemConfig", "POST"},
		{"p", "888", "/customer/customer", "POST"},
		{"p", "888", "/customer/customer", "PUT"},
		{"p", "888", "/customer/customer", "DELETE"},
		{"p", "888", "/customer/customer", "GET"},
		{"p", "888", "/customer/customerList", "GET"},
		{"p", "888", "/autoCode/createTemp", "POST"},
		{"p", "888", "/autoCode/getTables", "GET"},
		{"p", "888", "/autoCode/getDB", "GET"},
		{"p", "888", "/autoCode/getColume", "GET"},
		{"p", "888", "/sysDictionaryDetail/createSysDictionaryDetail", "POST"},
		{"p", "888", "/sysDictionaryDetail/deleteSysDictionaryDetail", "DELETE"},
		{"p", "888", "/sysDictionaryDetail/updateSysDictionaryDetail", "PUT"},
		{"p", "888", "/sysDictionaryDetail/findSysDictionaryDetail", "GET"},
		{"p", "888", "/sysDictionaryDetail/getSysDictionaryDetailList", "GET"},
		{"p", "888", "/sysDictionary/createSysDictionary", "POST"},
		{"p", "888", "/sysDictionary/deleteSysDictionary", "DELETE"},
		{"p", "888", "/sysDictionary/updateSysDictionary", "PUT"},
		{"p", "888", "/sysDictionary/findSysDictionary", "GET"},
		{"p", "888", "/sysDictionary/getSysDictionaryList", "GET"},
		{"p", "888", "/sysOperationRecord/createSysOperationRecord", "POST"},
		{"p", "888", "/sysOperationRecord/deleteSysOperationRecord", "DELETE"},
		{"p", "888", "/sysOperationRecord/updateSysOperationRecord", "PUT"},
		{"p", "888", "/sysOperationRecord/findSysOperationRecord", "GET"},
		{"p", "888", "/sysOperationRecord/getSysOperationRecordList", "GET"},
		{"p", "888", "/sysOperationRecord/deleteSysOperationRecordByIds", "DELETE"},
		{"p", "888", "/user/setUserInfo", "PUT"},
		{"p", "8881", "/base/login", "POST"},
		{"p", "8881", "/base/register", "POST"},
		{"p", "8881", "/api/createApi", "POST"},
		{"p", "8881", "/api/getApiList", "POST"},
		{"p", "8881", "/api/getApiById", "POST"},
		{"p", "8881", "/api/deleteApi", "POST"},
		{"p", "8881", "/api/updateApi", "POST"},
		{"p", "8881", "/api/getAllApis", "POST"},
		{"p", "8881", "/authority/createAuthority", "POST"},
		{"p", "8881", "/authority/deleteAuthority", "POST"},
		{"p", "8881", "/authority/getAuthorityList", "POST"},
		{"p", "8881", "/authority/setDataAuthority", "POST"},
		{"p", "8881", "/menu/getMenu", "POST"},
		{"p", "8881", "/menu/getMenuList", "POST"},
		{"p", "8881", "/menu/addBaseMenu", "POST"},
		{"p", "8881", "/menu/getBaseMenuTree", "POST"},
		{"p", "8881", "/menu/addMenuAuthority", "POST"},
		{"p", "8881", "/menu/getMenuAuthority", "POST"},
		{"p", "8881", "/menu/deleteBaseMenu", "POST"},
		{"p", "8881", "/menu/updateBaseMenu", "POST"},
		{"p", "8881", "/menu/getBaseMenuById", "POST"},
		{"p", "8881", "/user/changePassword", "POST"},
		{"p", "8881", "/user/getUserList", "POST"},
		{"p", "8881", "/user/setUserAuthority", "POST"},
		{"p", "8881", "/fileUploadAndDownload/upload", "POST"},
		{"p", "8881", "/fileUploadAndDownload/getFileList", "POST"},
		{"p", "8881", "/fileUploadAndDownload/deleteFile", "POST"},
		{"p", "8881", "/casbin/updateCasbin", "POST"},
		{"p", "8881", "/casbin/getPolicyPathByAuthorityId", "POST"},
		{"p", "8881", "/jwt/jsonInBlacklist", "POST"},
		{"p", "8881", "/system/getSystemConfig", "POST"},
		{"p", "8881", "/system/setSystemConfig", "POST"},
		{"p", "8881", "/customer/customer", "POST"},
		{"p", "8881", "/customer/customer", "PUT"},
		{"p", "8881", "/customer/customer", "DELETE"},
		{"p", "8881", "/customer/customer", "GET"},
		{"p", "8881", "/customer/customerList", "GET"},
		{"p", "9528", "/base/login", "POST"},
		{"p", "9528", "/base/register", "POST"},
		{"p", "9528", "/api/createApi", "POST"},
		{"p", "9528", "/api/getApiList", "POST"},
		{"p", "9528", "/api/getApiById", "POST"},
		{"p", "9528", "/api/deleteApi", "POST"},
		{"p", "9528", "/api/updateApi", "POST"},
		{"p", "9528", "/api/getAllApis", "POST"},
		{"p", "9528", "/authority/createAuthority", "POST"},
		{"p", "9528", "/authority/deleteAuthority", "POST"},
		{"p", "9528", "/authority/getAuthorityList", "POST"},
		{"p", "9528", "/authority/setDataAuthority", "POST"},
		{"p", "9528", "/menu/getMenu", "POST"},
		{"p", "9528", "/menu/getMenuList", "POST"},
		{"p", "9528", "/menu/addBaseMenu", "POST"},
		{"p", "9528", "/menu/getBaseMenuTree", "POST"},
		{"p", "9528", "/menu/addMenuAuthority", "POST"},
		{"p", "9528", "/menu/getMenuAuthority", "POST"},
		{"p", "9528", "/menu/deleteBaseMenu", "POST"},
		{"p", "9528", "/menu/updateBaseMenu", "POST"},
		{"p", "9528", "/menu/getBaseMenuById", "POST"},
		{"p", "9528", "/user/changePassword", "POST"},
		{"p", "9528", "/user/getUserList", "POST"},
		{"p", "9528", "/user/setUserAuthority", "POST"},
		{"p", "9528", "/fileUploadAndDownload/upload", "POST"},
		{"p", "9528", "/fileUploadAndDownload/getFileList", "POST"},
		{"p", "9528", "/fileUploadAndDownload/deleteFile", "POST"},
		{"p", "9528", "/casbin/updateCasbin", "POST"},
		{"p", "9528", "/casbin/getPolicyPathByAuthorityId", "POST"},
		{"p", "9528", "/jwt/jsonInBlacklist", "POST"},
		{"p", "9528", "/system/getSystemConfig", "POST"},
		{"p", "9528", "/system/setSystemConfig", "POST"},
		{"p", "9528", "/customer/customer", "POST"},
		{"p", "9528", "/customer/customer", "PUT"},
		{"p", "9528", "/customer/customer", "DELETE"},
		{"p", "9528", "/customer/customer", "GET"},
		{"p", "9528", "/customer/customerList", "GET"},
		{"p", "9528", "/autoCode/createTemp", "POST"},
	}
	if tx.Table("casbin_rule").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysAuthority() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysAuthority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysBaseMenus() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysBaseMenu{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{Title: "关于我们", Icon: "info"}},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: model.Meta{Title: "示例文件", Icon: "s-management"}},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "table", Name: "table", Component: "view/example/table/table.vue", Sort: 1, Meta: model.Meta{Title: "表格示例", Icon: "s-order"}},
		{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "form", Name: "form", Component: "view/example/form/form.vue", Sort: 2, Meta: model.Meta{Title: "表单示例", Icon: "document"}},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "rte", Name: "rte", Component: "view/example/rte/rte.vue", Sort: 3, Meta: model.Meta{Title: "富文本编辑器", Icon: "reading"}},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: model.Meta{Title: "excel导入导出", Icon: "s-marketing"}},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: model.Meta{Title: "上传下载", Icon: "upload"}},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: model.Meta{Title: "断点续传", Icon: "upload"}},
		{Model: gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: model.Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "iconList", Name: "iconList", Component: "view/iconList/index.vue", Sort: 2, Meta: model.Meta{Title: "图标集合", Icon: "star-on"}},
		{Model: gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "notebook-2"}},
		{Model: gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "s-order"}},
		{Model: gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
		{Model: gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: model.Meta{Title: "断点续传（插件版）", Icon: "upload"}},
		{Model: gorm.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: false, Component: "/", Sort: 0, Meta: model.Meta{Title: "官方网站", Icon: "s-home"}},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitAuthorityMenu() (err error) {
	return global.GVA_DB.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`, `sys_base_menus`.`updated_at` AS `updated_at`, `sys_base_menus`.`deleted_at` AS `deleted_at`, `sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`, `sys_base_menus`.`title`  AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on ((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)))").Error
}

func InitSysDictionary() (err error) {
	status := new(bool)
	*status = true
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysDictionary{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysAuthorityMenus() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []SysAuthorityMenus{
		{"888", 1},
		{"888", 2},
		{"888", 3},
		{"888", 4},
		{"888", 5},
		{"888", 6},
		{"888", 7},
		{"888", 8},
		{"888", 9},
		{"888", 10},
		{"888", 11},
		{"888", 12},
		{"888", 13},
		{"888", 14},
		{"888", 15},
		{"888", 16},
		{"888", 17},
		{"888", 18},
		{"888", 19},
		{"888", 20},
		{"888", 21},
		{"888", 22},
		{"888", 23},
		{"888", 24},
		{"888", 25},
		{"888", 26},
		{"8881", 1},
		{"8881", 2},
		{"8881", 8},
		{"8881", 17},
		{"8881", 18},
		{"8881", 19},
		{"8881", 20},
		{"9528", 1},
		{"9528", 2},
		{"9528", 3},
		{"9528", 4},
		{"9528", 5},
		{"9528", 6},
		{"9528", 7},
		{"9528", 8},
		{"9528", 9},
		{"9528", 10},
		{"9528", 11},
		{"9528", 12},
		{"9528", 13},
		{"9528", 14},
		{"9528", 15},
		{"9528", 17},
		{"9528", 18},
		{"9528", 19},
		{"9528", 20},
	}
	if tx.Table("sys_authority_menus").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDataAuthorityId() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []SysDataAuthorityId{
		{"888", "888"},
		{"888", "8881"},
		{"888", "9528"},
		{"9528", "8881"},
		{"9528", "9528"},
	}
	if tx.Table("sys_data_authority_ids").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDictionaryDetail() (err error) {
	status := new(bool)
	*status = true
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysDictionaryDetail{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "smallint", 1, status, 1, 2},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumint", 2, status, 2, 2},
		{gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "int", 3, status, 3, 2},
		{gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "bigint", 4, status, 4, 2},
		{gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "data", 0, status, 0, 3},
		{gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "time", 1, status, 1, 3},
		{gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "year", 2, status, 2, 3},
		{gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "datetime", 3, status, 3, 3},
		{gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "timestamp", 5, status, 5, 3},
		{gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "float", 0, status, 0, 4},
		{gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "double", 1, status, 1, 4},
		{gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "decimal", 2, status, 2, 4},
		{gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "char", 0, status, 0, 5},
		{gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "varchar", 1, status, 1, 5},
		{gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyblob", 2, status, 2, 5},
		{gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinytext", 3, status, 3, 5},
		{gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "text", 4, status, 4, 5},
		{gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "blob", 5, status, 5, 5},
		{gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumblob", 6, status, 6, 5},
		{gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumtext", 7, status, 7, 5},
		{gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longblob", 8, status, 8, 5},
		{gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longtext", 9, status, 9, 5},
		{gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyint", 0, status, 0, 6},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitExaFileUploadAndDownload() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.ExaFileUploadAndDownload{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitData() {
	var err error
	err = InitSysApi()
	err = InitSysUser()
	err = InitExaCustomer()
	err = InitCasbinModel()
	err = InitSysAuthority()
	err = InitSysBaseMenus()
	err = InitAuthorityMenu()
	err = InitSysDictionary()
	err = InitSysAuthorityMenus()
	err = InitSysDataAuthorityId()
	err = InitSysDictionaryDetail()
	err = InitExaFileUploadAndDownload()
	if err != nil {
		global.GVA_LOG.Error("initialize data failed", zap.Any("err", err))
	}
	global.GVA_LOG.Debug("initialize data success")
}
