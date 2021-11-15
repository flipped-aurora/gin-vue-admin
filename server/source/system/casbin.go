package system

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

func (c *casbin) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c *casbin) Initialize() error {
	entities := []adapter.CasbinRule{
		{PType: "p", V0: "888", V1: "POST", V2: "/base/login"},
		{PType: "p", V0: "888", V1: "POST", V2: "/user/register"},

		{PType: "p", V0: "888", V1: "POST", V2: "/api/createApi"},
		{PType: "p", V0: "888", V1: "POST", V2: "/api/getApiList"},
		{PType: "p", V0: "888", V1: "POST", V2: "/api/getApiById"},
		{PType: "p", V0: "888", V1: "POST", V2: "/api/deleteApi"},
		{PType: "p", V0: "888", V1: "POST", V2: "/api/updateApi"},
		{PType: "p", V0: "888", V1: "POST", V2: "/api/getAllApis"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/api/deleteApisByIds"},

		{PType: "p", V0: "888", V1: "POST", V2: "/authority/copyAuthority"},
		{PType: "p", V0: "888", V1: "PUT", V2: "/authority/updateAuthority"},
		{PType: "p", V0: "888", V1: "POST", V2: "/authority/createAuthority"},
		{PType: "p", V0: "888", V1: "POST", V2: "/authority/deleteAuthority"},
		{PType: "p", V0: "888", V1: "POST", V2: "/authority/getAuthorityList"},
		{PType: "p", V0: "888", V1: "POST", V2: "/authority/setDataAuthority"},

		{PType: "p", V0: "888", V1: "POST", V2: "/menu/getMenu"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/getMenuList"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/addBaseMenu"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/getBaseMenuTree"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/addMenuAuthority"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/getMenuAuthority"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/deleteBaseMenu"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/updateBaseMenu"},
		{PType: "p", V0: "888", V1: "POST", V2: "/menu/getBaseMenuById"},

		{PType: "p", V0: "888", V1: "GET", V2: "/user/getUserInfo"},
		{PType: "p", V0: "888", V1: "PUT", V2: "/user/setUserInfo"},
		{PType: "p", V0: "888", V1: "POST", V2: "/user/getUserList"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/user/deleteUser"},
		{PType: "p", V0: "888", V1: "POST", V2: "/user/changePassword"},
		{PType: "p", V0: "888", V1: "POST", V2: "/user/setUserAuthority"},
		{PType: "p", V0: "888", V1: "POST", V2: "/user/setUserAuthorities"},

		{PType: "p", V0: "888", V1: "POST", V2: "/fileUploadAndDownload/upload"},
		{PType: "p", V0: "888", V1: "POST", V2: "/fileUploadAndDownload/deleteFile"},
		{PType: "p", V0: "888", V1: "POST", V2: "/fileUploadAndDownload/getFileList"},

		{PType: "p", V0: "888", V1: "POST", V2: "/casbin/updateCasbin"},
		{PType: "p", V0: "888", V1: "POST", V2: "/casbin/getPolicyPathByAuthorityId"},

		{PType: "p", V0: "888", V1: "POST", V2: "/jwt/jsonInBlacklist"},

		{PType: "p", V0: "888", V1: "POST", V2: "/system/getSystemConfig"},
		{PType: "p", V0: "888", V1: "POST", V2: "/system/setSystemConfig"},
		{PType: "p", V0: "888", V1: "POST", V2: "/system/getServerInfo"},

		{PType: "p", V0: "888", V1: "GET", V2: "/customer/customer"},
		{PType: "p", V0: "888", V1: "PUT", V2: "/customer/customer"},
		{PType: "p", V0: "888", V1: "POST", V2: "/customer/customer"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/customer/customer"},
		{PType: "p", V0: "888", V1: "GET", V2: "/customer/customerList"},

		{PType: "p", V0: "888", V1: "GET", V2: "/autoCode/getDB"},
		{PType: "p", V0: "888", V1: "POST", V2: "/autoCode/getMeta"},
		{PType: "p", V0: "888", V1: "POST", V2: "/autoCode/preview"},
		{PType: "p", V0: "888", V1: "GET", V2: "/autoCode/getTables"},
		{PType: "p", V0: "888", V1: "GET", V2: "/autoCode/getColumn"},
		{PType: "p", V0: "888", V1: "POST", V2: "/autoCode/rollback"},
		{PType: "p", V0: "888", V1: "POST", V2: "/autoCode/createTemp"},
		{PType: "p", V0: "888", V1: "POST", V2: "/autoCode/delSysHistory"},
		{PType: "p", V0: "888", V1: "POST", V2: "/autoCode/getSysHistory"},

		{PType: "p", V0: "888", V1: "GET", V2: "/sysDictionaryDetail/findSysDictionaryDetail"},
		{PType: "p", V0: "888", V1: "PUT", V2: "/sysDictionaryDetail/updateSysDictionaryDetail"},
		{PType: "p", V0: "888", V1: "POST", V2: "/sysDictionaryDetail/createSysDictionaryDetail"},
		{PType: "p", V0: "888", V1: "GET", V2: "/sysDictionaryDetail/getSysDictionaryDetailList"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/sysDictionaryDetail/deleteSysDictionaryDetail"},

		{PType: "p", V0: "888", V1: "GET", V2: "/sysDictionary/findSysDictionary"},
		{PType: "p", V0: "888", V1: "PUT", V2: "/sysDictionary/updateSysDictionary"},
		{PType: "p", V0: "888", V1: "GET", V2: "/sysDictionary/getSysDictionaryList"},
		{PType: "p", V0: "888", V1: "POST", V2: "/sysDictionary/createSysDictionary"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/sysDictionary/deleteSysDictionary"},

		{PType: "p", V0: "888", V1: "GET", V2: "/sysOperationRecord/findSysOperationRecord"},
		{PType: "p", V0: "888", V1: "PUT", V2: "/sysOperationRecord/updateSysOperationRecord"},
		{PType: "p", V0: "888", V1: "POST", V2: "/sysOperationRecord/createSysOperationRecord"},
		{PType: "p", V0: "888", V1: "GET", V2: "/sysOperationRecord/getSysOperationRecordList"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/sysOperationRecord/deleteSysOperationRecord"},
		{PType: "p", V0: "888", V1: "DELETE", V2: "/sysOperationRecord/deleteSysOperationRecordByIds"},

		{PType: "p", V0: "888", V1: "POST", V2: "/email/emailTest"},

		{PType: "p", V0: "888", V1: "POST", V2: "/simpleUploader/upload"},
		{PType: "p", V0: "888", V1: "GET", V2: "/simpleUploader/checkFileMd5"},
		{PType: "p", V0: "888", V1: "GET", V2: "/simpleUploader/mergeFileMd5"},

		{PType: "p", V0: "888", V1: "POST", V2: "/excel/importExcel"},
		{PType: "p", V0: "888", V1: "GET", V2: "/excel/loadExcel"},
		{PType: "p", V0: "888", V1: "POST", V2: "/excel/exportExcel"},
		{PType: "p", V0: "888", V1: "GET", V2: "/excel/downloadTemplate"},

		{PType: "p", V0: "8881", V1: "POST", V2: "/base/login"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/user/register"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/api/createApi"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/api/getApiList"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/api/getApiById"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/api/deleteApi"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/api/updateApi"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/api/getAllApis"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/authority/createAuthority"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/authority/deleteAuthority"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/authority/getAuthorityList"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/authority/setDataAuthority"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/getMenu"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/getMenuList"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/addBaseMenu"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/getBaseMenuTree"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/addMenuAuthority"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/getMenuAuthority"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/deleteBaseMenu"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/updateBaseMenu"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/menu/getBaseMenuById"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/user/changePassword"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/user/getUserList"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/user/setUserAuthority"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/fileUploadAndDownload/upload"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/fileUploadAndDownload/getFileList"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/fileUploadAndDownload/deleteFile"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/casbin/updateCasbin"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/casbin/getPolicyPathByAuthorityId"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/jwt/jsonInBlacklist"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/system/getSystemConfig"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/system/setSystemConfig"},
		{PType: "p", V0: "8881", V1: "POST", V2: "/customer/customer"},
		{PType: "p", V0: "8881", V1: "PUT", V2: "/customer/customer"},
		{PType: "p", V0: "8881", V1: "DELETE", V2: "/customer/customer"},
		{PType: "p", V0: "8881", V1: "GET", V2: "/customer/customer"},
		{PType: "p", V0: "8881", V1: "GET", V2: "/customer/customerList"},
		{PType: "p", V0: "8881", V1: "GET", V2: "/user/getUserInfo"},

		{PType: "p", V0: "9528", V1: "POST", V2: "/base/login"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/user/register"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/api/createApi"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/api/getApiList"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/api/getApiById"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/api/deleteApi"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/api/updateApi"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/api/getAllApis"},

		{PType: "p", V0: "9528", V1: "POST", V2: "/authority/createAuthority"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/authority/deleteAuthority"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/authority/getAuthorityList"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/authority/setDataAuthority"},

		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/getMenu"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/getMenuList"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/addBaseMenu"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/getBaseMenuTree"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/addMenuAuthority"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/getMenuAuthority"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/deleteBaseMenu"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/updateBaseMenu"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/menu/getBaseMenuById"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/user/changePassword"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/user/getUserList"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/user/setUserAuthority"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/fileUploadAndDownload/upload"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/fileUploadAndDownload/getFileList"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/fileUploadAndDownload/deleteFile"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/casbin/updateCasbin"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/casbin/getPolicyPathByAuthorityId"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/jwt/jsonInBlacklist"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/system/getSystemConfig"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/system/setSystemConfig"},
		{PType: "p", V0: "9528", V1: "PUT", V2: "/customer/customer"},
		{PType: "p", V0: "9528", V1: "GET", V2: "/customer/customer"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/customer/customer"},
		{PType: "p", V0: "9528", V1: "DELETE", V2: "/customer/customer"},
		{PType: "p", V0: "9528", V1: "GET", V2: "/customer/customerList"},
		{PType: "p", V0: "9528", V1: "POST", V2: "/autoCode/createTemp"},
		{PType: "p", V0: "9528", V1: "GET", V2: "/user/getUserInfo"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, c.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (c *casbin) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where(adapter.CasbinRule{PType: "p", V0: "9528", V1: "GET", V2: "/user/getUserInfo"}).First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
