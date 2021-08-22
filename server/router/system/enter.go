package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	AutoCodeRouter
	BaseRouter
	CasbinRouter
	DictionaryRouter
	DictionaryDetailRouter
	InitRouter
	JwtRouter
	MenuRouter
	OperationRecordRouter
	SysRouter
	UserRouter
}
