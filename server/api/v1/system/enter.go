package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SystemApiApi
	AuthorityApi
	AutoCodeApi
	BaseApi
	CasbinApi
	DictionaryApi
	DictionaryDetailApi
	SystemApi
	DBApi
	JwtApi
	OperationRecordApi
	AuthorityMenuApi
}

var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var autoCodeService = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
var autoCodeHistoryService = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
var dictionaryService = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
var dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
var initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
