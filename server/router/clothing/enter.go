package clothing

type RouterGroup struct {
	CompanyRouter
	AppRoleRouter
	TeamRouter
	AppUserRouter
	UserRoleRouter
	UserWalletRouter
}
