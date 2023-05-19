package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserWalletRouter struct {
}

// InitUserWalletRouter 初始化 UserWallet 路由信息
func (s *UserWalletRouter) InitUserWalletRouter(Router *gin.RouterGroup) {
	userWalletRouter := Router.Group("userWallet").Use(middleware.OperationRecord())
	userWalletRouterWithoutRecord := Router.Group("userWallet")
	h5UserWalletRouterWithoutRecord := Router.Group(global.GetAppApi() + "userWallet")
	var userWalletApi = v1.ApiGroupApp.ClothingApiGroup.UserWalletApi
	{
		userWalletRouter.POST("createUserWallet", userWalletApi.CreateUserWallet)             // 新建UserWallet
		userWalletRouter.DELETE("deleteUserWallet", userWalletApi.DeleteUserWallet)           // 删除UserWallet
		userWalletRouter.DELETE("deleteUserWalletByIds", userWalletApi.DeleteUserWalletByIds) // 批量删除UserWallet
		userWalletRouter.PUT("updateUserWallet", userWalletApi.UpdateUserWallet)              // 更新UserWallet
	}
	{
		userWalletRouterWithoutRecord.GET("findUserWallet", userWalletApi.FindUserWallet)       // 根据ID获取UserWallet
		userWalletRouterWithoutRecord.GET("getUserWalletList", userWalletApi.GetUserWalletList) // 获取UserWallet列表
	}
	{
		h5UserWalletRouterWithoutRecord.GET("getMyWalletList", userWalletApi.GetMyWallet)         // 获取UserWallet列表
		h5UserWalletRouterWithoutRecord.GET("getUserWalletList", userWalletApi.GetUserWalletList) // 获取UserWallet列表
	}
}
