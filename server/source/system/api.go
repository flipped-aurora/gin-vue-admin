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
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "system.api.desc.jwtAddedToBlackList"},

		{ApiGroup: "system.api.group.systemUser", Method: "DELETE", Path: "/user/deleteUser", Description: "system.api.desc.deleteUsers"},
		{ApiGroup: "system.api.group.systemUser", Method: "POST", Path: "/user/admin_register", Description: "system.api.desc.userRegistration"},
		{ApiGroup: "system.api.group.systemUser", Method: "POST", Path: "/user/getUserList", Description: "system.api.desc.getUsersList"},
		{ApiGroup: "system.api.group.systemUser", Method: "PUT", Path: "/user/setUserInfo", Description: "system.api.desc.setUserInfo"},
		{ApiGroup: "system.api.group.systemUser", Method: "PUT", Path: "/user/setSelfInfo", Description: "system.api.desc.setSelfInfo"},
		{ApiGroup: "system.api.group.systemUser", Method: "GET", Path: "/user/getUserInfo", Description: "system.api.desc.getSelfInfo"},
		{ApiGroup: "system.api.group.systemUser", Method: "POST", Path: "/user/setUserAuthorities", Description: "system.api.desc.setPermissionGroup"},
		{ApiGroup: "system.api.group.systemUser", Method: "POST", Path: "/user/changePassword", Description: "system.api.desc.changePassword"},
		{ApiGroup: "system.api.group.systemUser", Method: "POST", Path: "/user/setUserAuthority", Description: "system.api.desc.modifyUserRole"},
		{ApiGroup: "system.api.group.systemUser", Method: "POST", Path: "/user/resetPassword", Description: "system.api.desc.resetUserPassword"},

		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/createApi", Description: "system.api.desc.createAPI"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/deleteApi", Description: "system.api.desc.deleteAPI"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/updateApi", Description: "system.api.desc.updateAPI"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/getApiList", Description: "system.api.desc.getAPIList"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/getAllApis", Description: "system.api.desc.getAllAPI"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/getApiById", Description: "system.api.desc.getAPIByID"},
		{ApiGroup: "system.api.group.api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "system.api.desc.deleteAPIByID"},
		{ApiGroup: "system.api.group.api", Method: "GET", Path: "/api/syncApi", Description: "system.api.desc.getSyncApi"},
		{ApiGroup: "system.api.group.api", Method: "GET", Path: "/api/getApiGroups", Description: "system.api.desc.getRouteGroup"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/enterSyncApi", Description: "system.api.desc.confirmSyncApi"},
		{ApiGroup: "system.api.group.api", Method: "POST", Path: "/api/ignoreApi", Description: "system.api.desc.ignoreApi"},

		{ApiGroup: "system.api.group.role", Method: "POST", Path: "/authority/copyAuthority", Description: "system.api.desc.copyRole"},
		{ApiGroup: "system.api.group.role", Method: "POST", Path: "/authority/createAuthority", Description: "system.api.desc.createRole"},
		{ApiGroup: "system.api.group.role", Method: "POST", Path: "/authority/deleteAuthority", Description: "system.api.desc.deleteRole"},
		{ApiGroup: "system.api.group.role", Method: "PUT", Path: "/authority/updateAuthority", Description: "system.api.desc.updateRole"},
		{ApiGroup: "system.api.group.role", Method: "POST", Path: "/authority/getAuthorityList", Description: "system.api.desc.getRoleList"},
		{ApiGroup: "system.api.group.role", Method: "POST", Path: "/authority/setDataAuthority", Description: "system.api.desc.setRolePermissions"},

		{ApiGroup: "system.api.group.casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "system.api.desc.changeRoleAPIPermission"},
		{ApiGroup: "system.api.group.casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "system.api.desc.getPermissionList"},

		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/addBaseMenu", Description: "system.api.desc.addMenu"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/getMenu", Description: "system.api.desc.getMenuTree"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "system.api.desc.deleteMenu"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/updateBaseMenu", Description: "system.api.desc.updateMenu"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/getBaseMenuById", Description: "system.api.desc.getMenuByID"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/getMenuList", Description: "system.api.desc.getMenuList"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "system.api.desc.getDynamicRoute"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/getMenuAuthority", Description: "system.api.desc.getMenuRole"},
		{ApiGroup: "system.api.group.menu", Method: "POST", Path: "/menu/addMenuAuthority", Description: "system.api.desc.addMenuRole"},

		{ApiGroup: "system.api.group.partialUpload", Method: "POST", Path: "/fileUploadAndDownload/findFile", Description: "system.api.desc.findTargetFile"},
		{ApiGroup: "system.api.group.partialUpload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "system.api.desc.breakpointContinue"},
		{ApiGroup: "system.api.group.partialUpload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "system.api.desc.breakpointContinueFinish"},
		{ApiGroup: "system.api.group.partialUpload", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "system.api.desc.removeFileAfterUpload"},

		{ApiGroup: "system.api.group.fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "system.api.desc.fileUploadExample"},
		{ApiGroup: "system.api.group.fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "system.api.desc.deleteFile"},
		{ApiGroup: "system.api.group.fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "system.api.desc.editFileNameOrRemark"},
		{ApiGroup: "system.api.group.fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "system.api.desc.getUploadFileList"},

		{ApiGroup: "system.api.group.systemService", Method: "POST", Path: "/system/getServerInfo", Description: "system.api.desc.getServerInfo"},
		{ApiGroup: "system.api.group.systemService", Method: "POST", Path: "/system/getSystemConfig", Description: "system.api.desc.getConfigFileContent"},
		{ApiGroup: "system.api.group.systemService", Method: "POST", Path: "/system/setSystemConfig", Description: "system.api.desc.setConfigFileContent"},

		{ApiGroup: "system.api.group.customer", Method: "PUT", Path: "/customer/customer", Description: "system.api.desc.updateCustomer"},
		{ApiGroup: "system.api.group.customer", Method: "POST", Path: "/customer/customer", Description: "system.api.desc.createCustomer"},
		{ApiGroup: "system.api.group.customer", Method: "DELETE", Path: "/customer/customer", Description: "system.api.desc.deleteCustomer"},
		{ApiGroup: "system.api.group.customer", Method: "GET", Path: "/customer/customer", Description: "system.api.desc.getSingleCustomer"},
		{ApiGroup: "system.api.group.customer", Method: "GET", Path: "/customer/customerList", Description: "system.api.desc.getCustomerList"},

		{ApiGroup: "system.api.group.customer", Method: "GET", Path: "/autoCode/getDB", Description: "system.api.desc.getAllDatabases"},
		{ApiGroup: "system.api.group.customer", Method: "GET", Path: "/autoCode/getTables", Description: "system.api.desc.getDatabaseTables"},
		{ApiGroup: "system.api.group.customer", Method: "POST", Path: "/autoCode/createTemp", Description: "system.api.desc.autoCode"},
		{ApiGroup: "system.api.group.customer", Method: "POST", Path: "/autoCode/preview", Description: "system.api.desc.previewAutoCode"},
		{ApiGroup: "system.api.group.customer", Method: "GET", Path: "/autoCode/getColumn", Description: "system.api.desc.getSelectedTableFields"},
		{ApiGroup: "system.api.group.customer", Method: "POST", Path: "/autoCode/createPlug", Description: "system.api.desc.createPluginPackage"},
		{ApiGroup: "system.api.group.customer", Method: "POST", Path: "/autoCode/installPlugin", Description: "system.api.desc.installPlugin"},
		{ApiGroup: "system.api.group.customer", Method: "POST", Path: "/autoCode/pubPlug", Description: "system.api.desc.packagePlugin"},

		{ApiGroup: "system.api.group.templateConfiguration", Method: "POST", Path: "/autoCode/createPackage", Description: "system.api.desc.configurationTemplates"},
		{ApiGroup: "system.api.group.templateConfiguration", Method: "GET", Path: "/autoCode/getTemplates", Description: "system.api.desc.getTemplateFile"},
		{ApiGroup: "system.api.group.templateConfiguration", Method: "POST", Path: "/autoCode/getPackage", Description: "system.api.desc.getAllTemplates"},
		{ApiGroup: "system.api.group.templateConfiguration", Method: "POST", Path: "/autoCode/delPackage", Description: "system.api.desc.deleteTemplate"},

		{ApiGroup: "system.api.group.codeGenHistory", Method: "POST", Path: "/autoCode/getMeta", Description: "system.api.desc.getMetaInfo"},
		{ApiGroup: "system.api.group.codeGenHistory", Method: "POST", Path: "/autoCode/rollback", Description: "system.api.desc.rollbackAutoGeneratedCode"},
		{ApiGroup: "system.api.group.codeGenHistory", Method: "POST", Path: "/autoCode/getSysHistory", Description: "system.api.desc.queryRollbackRecord"},
		{ApiGroup: "system.api.group.codeGenHistory", Method: "POST", Path: "/autoCode/delSysHistory", Description: "system.api.desc.deleteRollbackRecord"},
		{ApiGroup: "system.api.group.codeGenHistory", Method: "POST", Path: "/autoCode/addFunc", Description: "system.api.desc.addTemplateMethod"},

		{ApiGroup: "system.api.group.dictDetails", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "system.api.desc.updateDictionaryContent"},
		{ApiGroup: "system.api.group.dictDetails", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "system.api.desc.createDictionaryContent"},
		{ApiGroup: "system.api.group.dictDetails", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "system.api.desc.deleteDictionaryContent"},
		{ApiGroup: "system.api.group.dictDetails", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "system.api.desc.getDictionaryContentById"},
		{ApiGroup: "system.api.group.dictDetails", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "system.api.desc.getDictionaryContentList"},

		{ApiGroup: "system.api.group.dictionary", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "system.api.desc.createDictionary"},
		{ApiGroup: "system.api.group.dictionary", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "system.api.desc.deleteDictionary"},
		{ApiGroup: "system.api.group.dictionary", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "system.api.desc.updateDictionary"},
		{ApiGroup: "system.api.group.dictionary", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "system.api.desc.getDictionaryById"},
		{ApiGroup: "system.api.group.dictionary", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "system.api.desc.getDictionaryList"},

		{ApiGroup: "system.api.group.optRecord", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "system.api.desc.createOperationRecord"},
		{ApiGroup: "system.api.group.optRecord", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "system.api.desc.getOperationRecordById"},
		{ApiGroup: "system.api.group.optRecord", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "system.api.desc.getOperationRecordList"},
		{ApiGroup: "system.api.group.optRecord", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "system.api.desc.deleteOperationRecord"},
		{ApiGroup: "system.api.group.optRecord", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "system.api.desc.batchDeleteOperationHistory"},

		{ApiGroup: "system.api.group.resumeUpload", Method: "POST", Path: "/simpleUploader/upload", Description: "system.api.desc.pluginVersionResumableUpload"},
		{ApiGroup: "system.api.group.resumeUpload", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "system.api.desc.fileIntegrityCheck"},
		{ApiGroup: "system.api.group.resumeUpload", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "system.api.desc.mergeFileAfterUpload"},

		{ApiGroup: "system.api.group.email", Method: "POST", Path: "/email/emailTest", Description: "system.api.desc.sendTestEmail"},
		{ApiGroup: "system.api.group.email", Method: "POST", Path: "/email/sendEmail", Description: "system.api.desc.sendEmail"},

		{ApiGroup: "system.api.group.buttonAuthority", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "system.api.desc.setButtonPermission"},
		{ApiGroup: "system.api.group.buttonAuthority", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "system.api.desc.getExistingButtonPermission"},
		{ApiGroup: "system.api.group.buttonAuthority", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "system.api.desc.deleteButton"},

		{ApiGroup: "system.api.group.tableTemplate", Method: "POST", Path: "/sysExportTemplate/createSysExportTemplate", Description: "system.api.desc.createExportTemplate"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplate", Description: "system.api.desc.deleteExportTemplate"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplateByIds", Description: "system.api.desc.batchDeleteExportTemplate"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "PUT", Path: "/sysExportTemplate/updateSysExportTemplate", Description: "system.api.desc.updateExportTemplate"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "GET", Path: "/sysExportTemplate/findSysExportTemplate", Description: "system.api.desc.getExportTemplateById"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "GET", Path: "/sysExportTemplate/getSysExportTemplateList", Description: "system.api.desc.getExportTemplateList"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "GET", Path: "/sysExportTemplate/exportExcel", Description: "system.api.desc.exportExcel"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "GET", Path: "/sysExportTemplate/exportTemplate", Description: "system.api.desc.downloadTemplate"},
		{ApiGroup: "system.api.group.tableTemplate", Method: "POST", Path: "/sysExportTemplate/importExcel", Description: "system.api.desc.importExcel"},

		{ApiGroup: "system.api.group.announcement", Method: "POST", Path: "/info/createInfo", Description: "system.api.desc.newAnnouncement"},
		{ApiGroup: "system.api.group.announcement", Method: "DELETE", Path: "/info/deleteInfo", Description: "system.api.desc.deleteAnnouncement"},
		{ApiGroup: "system.api.group.announcement", Method: "DELETE", Path: "/info/deleteInfoByIds", Description: "system.api.desc.batchDeleteAnnouncement"},
		{ApiGroup: "system.api.group.announcement", Method: "PUT", Path: "/info/updateInfo", Description: "system.api.desc.updateAnnouncement"},
		{ApiGroup: "system.api.group.announcement", Method: "GET", Path: "/info/findInfo", Description: "system.api.desc.getAnnouncementByID"},
		{ApiGroup: "system.api.group.announcement", Method: "GET", Path: "/info/getInfoList", Description: "system.api.desc.getAnnouncementList"},
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
