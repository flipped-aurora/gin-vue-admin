package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"go.uber.org/zap"
)

func Data() {
	var err error
	err = service.InitSysApi()
	err = service.InitSysUser()
	err = service.InitExaCustomer()
	err = service.InitCasbinModel()
	err = service.InitSysAuthority()
	err = service.InitSysBaseMenus()
	err = service.InitAuthorityMenu()
	err = service.InitSysDictionary()
	err = service.InitSysAuthorityMenus()
	err = service.InitSysDataAuthorityId()
	err = service.InitSysDictionaryDetail()
	err = service.InitExaFileUploadAndDownload()
	if err != nil {
		global.GVA_LOG.Error("initialize data failed", zap.Any("err", err))
	}
	global.GVA_LOG.Debug("initialize data success")
}
