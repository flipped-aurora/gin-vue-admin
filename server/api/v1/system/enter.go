package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	AutoCodeHistoryApi
	DictionaryDetailApi
}

var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var autoCodeService = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
var baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var dictionaryService = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var autoCodeHistoryService = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
var dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
