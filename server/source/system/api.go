package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: global.Translate("system.api.jwtAddedToBlackList")},

		{ApiGroup: global.Translate("system.api.systemUser"), Method: "DELETE", Path: "/user/deleteUser", Description: global.Translate("system.api.deleteUsers")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "POST", Path: "/user/admin_register", Description: global.Translate("system.api.userRegistration")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "POST", Path: "/user/getUserList", Description: global.Translate("system.api.getUsersList")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "PUT", Path: "/user/setUserInfo", Description: global.Translate("system.api.setUserInfo")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "PUT", Path: "/user/setSelfInfo", Description: global.Translate("system.api.setSelfInfo")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "GET", Path: "/user/getUserInfo", Description: global.Translate("system.api.getSelfInfo")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "POST", Path: "/user/setUserAuthorities", Description: global.Translate("system.api.setPermissionGroup")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "POST", Path: "/user/changePassword", Description: global.Translate("system.api.changePassword")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "POST", Path: "/user/setUserAuthority", Description: global.Translate("system.api.modifyUserRole")},
		{ApiGroup: global.Translate("system.api.systemUser"), Method: "POST", Path: "/user/resetPassword", Description: global.Translate("system.api.resetUserPassword")},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: global.Translate("system.api.createAPI")},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: global.Translate("system.api.deleteAPI")},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: global.Translate("system.api.updateAPI")},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: global.Translate("system.api.getAPIList")},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: global.Translate("system.api.getAllAPI")},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: global.Translate("system.api.getAPIByID")},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: global.Translate("system.api.deleteAPIByID")},

		{ApiGroup: global.Translate("system.api.role"), Method: "POST", Path: "/authority/copyAuthority", Description: global.Translate("system.api.copyRole")},
		{ApiGroup: global.Translate("system.api.role"), Method: "POST", Path: "/authority/createAuthority", Description: global.Translate("system.api.createRole")},
		{ApiGroup: global.Translate("system.api.role"), Method: "POST", Path: "/authority/deleteAuthority", Description: global.Translate("system.api.deleteRole")},
		{ApiGroup: global.Translate("system.api.role"), Method: "PUT", Path: "/authority/updateAuthority", Description: global.Translate("system.api.updateRole")},
		{ApiGroup: global.Translate("system.api.role"), Method: "POST", Path: "/authority/getAuthorityList", Description: global.Translate("system.api.getRoleList")},
		{ApiGroup: global.Translate("system.api.role"), Method: "POST", Path: "/authority/setDataAuthority", Description: global.Translate("system.api.setRolePermissions")},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: global.Translate("system.api.changeRoleAPIPermission")},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: global.Translate("system.api.getPermissionList")},

		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/addBaseMenu", Description: global.Translate("system.api.addMenu")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/getMenu", Description: global.Translate("system.api.getMenuTree")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/deleteBaseMenu", Description: global.Translate("system.api.deleteMenu")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/updateBaseMenu", Description: global.Translate("system.api.updateMenu")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/getBaseMenuById", Description: global.Translate("system.api.getMenuByID")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/getMenuList", Description: global.Translate("system.api.getMenuList")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/getBaseMenuTree", Description: global.Translate("system.api.getDynamicRoute")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/getMenuAuthority", Description: global.Translate("system.api.getMenuRole")},
		{ApiGroup: global.Translate("system.api.menu"), Method: "POST", Path: "/menu/addMenuAuthority", Description: global.Translate("system.api.addMenuRole")},

		{ApiGroup: global.Translate("system.api.partialUpload"), Method: "POST", Path: "/fileUploadAndDownload/findFile", Description: "寻找目标文件（秒传）"},
		{ApiGroup: global.Translate("system.api.partialUpload"), Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "断点续传"},
		{ApiGroup: global.Translate("system.api.partialUpload"), Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "断点续传完成"},
		{ApiGroup: global.Translate("system.api.partialUpload"), Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "上传完成移除文件"},

		{ApiGroup: global.Translate("system.api.fileUploadDownload"), Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "文件上传示例"},
		{ApiGroup: global.Translate("system.api.fileUploadDownload"), Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件"},
		{ApiGroup: global.Translate("system.api.fileUploadDownload"), Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "文件名或者备注编辑"},
		{ApiGroup: global.Translate("system.api.fileUploadDownload"), Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表"},

		{ApiGroup: global.Translate("system.api.systemService"), Method: "POST", Path: "/system/getServerInfo", Description: "获取服务器信息"},
		{ApiGroup: global.Translate("system.api.systemService"), Method: "POST", Path: "/system/getSystemConfig", Description: "获取配置文件内容"},
		{ApiGroup: global.Translate("system.api.systemService"), Method: "POST", Path: "/system/setSystemConfig", Description: "设置配置文件内容"},

		{ApiGroup: global.Translate("system.api.customer"), Method: "PUT", Path: "/customer/customer", Description: "更新客户"},
		{ApiGroup: global.Translate("system.api.customer"), Method: "POST", Path: "/customer/customer", Description: "创建客户"},
		{ApiGroup: global.Translate("system.api.customer"), Method: "DELETE", Path: "/customer/customer", Description: "删除客户"},
		{ApiGroup: global.Translate("system.api.customer"), Method: "GET", Path: "/customer/customer", Description: "获取单一客户"},
		{ApiGroup: global.Translate("system.api.customer"), Method: "GET", Path: "/customer/customerList", Description: "获取客户列表"},

		{ApiGroup: global.Translate("system.api.codeGen"), Method: "GET", Path: "/autoCode/getDB", Description: "获取所有数据库"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "GET", Path: "/autoCode/getTables", Description: "获取数据库表"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "POST", Path: "/autoCode/createTemp", Description: "自动化代码"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "POST", Path: "/autoCode/preview", Description: "预览自动化代码"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "GET", Path: "/autoCode/getColumn", Description: "获取所选table的所有字段"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "POST", Path: "/autoCode/createPlug", Description: "自从创建插件包"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "POST", Path: "/autoCode/installPlugin", Description: "安装插件"},
		{ApiGroup: global.Translate("system.api.codeGen"), Method: "POST", Path: "/autoCode/pubPlug", Description: "打包插件"},

		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/createPackage", Description: "生成包(package)"},
		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/getPackage", Description: "获取所有包(package)"},
		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/delPackage", Description: "删除包(package)"},

		{ApiGroup: global.Translate("system.api.codeGenHistory"), Method: "POST", Path: "/autoCode/getMeta", Description: "获取meta信息"},
		{ApiGroup: global.Translate("system.api.codeGenHistory"), Method: "POST", Path: "/autoCode/rollback", Description: "回滚自动生成代码"},
		{ApiGroup: global.Translate("system.api.codeGenHistory"), Method: "POST", Path: "/autoCode/getSysHistory", Description: "查询回滚记录"},
		{ApiGroup: global.Translate("system.api.codeGenHistory"), Method: "POST", Path: "/autoCode/delSysHistory", Description: "删除回滚记录"},

		{ApiGroup: global.Translate("system.api.dictDetails"), Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容"},
		{ApiGroup: global.Translate("system.api.dictDetails"), Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容"},
		{ApiGroup: global.Translate("system.api.dictDetails"), Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容"},
		{ApiGroup: global.Translate("system.api.dictDetails"), Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容"},
		{ApiGroup: global.Translate("system.api.dictDetails"), Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表"},

		{ApiGroup: global.Translate("system.api.dictionary"), Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "新增字典"},
		{ApiGroup: global.Translate("system.api.dictionary"), Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典"},
		{ApiGroup: global.Translate("system.api.dictionary"), Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "更新字典"},
		{ApiGroup: global.Translate("system.api.dictionary"), Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典"},
		{ApiGroup: global.Translate("system.api.dictionary"), Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表"},

		{ApiGroup: global.Translate("system.api.optRecord"), Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录"},
		{ApiGroup: global.Translate("system.api.optRecord"), Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录"},
		{ApiGroup: global.Translate("system.api.optRecord"), Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表"},
		{ApiGroup: global.Translate("system.api.optRecord"), Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录"},
		{ApiGroup: global.Translate("system.api.optRecord"), Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史"},

		{ApiGroup: global.Translate("system.api.resumeUpload"), Method: "POST", Path: "/simpleUploader/upload", Description: "插件版分片上传"},
		{ApiGroup: global.Translate("system.api.resumeUpload"), Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证"},
		{ApiGroup: global.Translate("system.api.resumeUpload"), Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "发送测试邮件"},
		{ApiGroup: "email", Method: "POST", Path: "/email/emailSend", Description: "发送邮件示例"},

		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "设置按钮权限"},
		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "获取已有按钮权限"},
		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "删除按钮"},

		{ApiGroup: "万用表格", Method: "POST", Path: "/chatGpt/getTable", Description: "通过gpt获取内容"},
		{ApiGroup: "万用表格", Method: "POST", Path: "/chatGpt/createSK", Description: "录入sk"},
		{ApiGroup: "万用表格", Method: "GET", Path: "/chatGpt/getSK", Description: "获取sk"},
		{ApiGroup: "万用表格", Method: "DELETE", Path: "/chatGpt/deleteSK", Description: "删除sk"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+" "+global.Translate("general.tabelDataInitFail"))
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
