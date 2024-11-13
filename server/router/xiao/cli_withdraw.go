package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliWithdrawRouter struct{}

// InitCliWithdrawRouter 初始化 提币详情 路由信息
func (s *CliWithdrawRouter) InitCliWithdrawRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cliwithdrawRouter := Router.Group("cliwithdraw").Use(middleware.OperationRecord())
	cliwithdrawRouterWithoutRecord := Router.Group("cliwithdraw")
	cliwithdrawRouterWithoutAuth := PublicRouter.Group("cliwithdraw")
	myjwtRouter := PublicRouter.Group("cliwithdraw").Use(middleware.JWTAuth())
	{
		cliwithdrawRouter.POST("createCliWithdraw", cliwithdrawApi.CreateCliWithdraw)             // 新建提币详情
		cliwithdrawRouter.DELETE("deleteCliWithdraw", cliwithdrawApi.DeleteCliWithdraw)           // 删除提币详情
		cliwithdrawRouter.DELETE("deleteCliWithdrawByIds", cliwithdrawApi.DeleteCliWithdrawByIds) // 批量删除提币详情
		cliwithdrawRouter.PUT("updateCliWithdraw", cliwithdrawApi.UpdateCliWithdraw)              // 更新提币详情
	}
	{
		cliwithdrawRouterWithoutRecord.GET("findCliWithdraw", cliwithdrawApi.FindCliWithdraw)       // 根据ID获取提币详情
		cliwithdrawRouterWithoutRecord.GET("getCliWithdrawList", cliwithdrawApi.GetCliWithdrawList) // 获取提币详情列表
	}
	{
		cliwithdrawRouterWithoutAuth.GET("getCliWithdrawPublic", cliwithdrawApi.GetCliWithdrawPublic) // 提币详情开放接口
	}
	{
		myjwtRouter.POST("dowithdraw", cliwithdrawApi.CliDoWithdraw)
	}
}
