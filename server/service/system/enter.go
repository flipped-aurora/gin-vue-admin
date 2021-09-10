package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
	AutoCodeService
	AutoCodeHistoryService
	BaseMenuService
	CasbinService
	DictionaryService
	DictionaryDetailService
	InitDBService
	MenuService
	OperationRecordService
	SystemConfigService
	UserService
}
