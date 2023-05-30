package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RechargeOptionRouter struct {
}

// InitRechargeOptionRouter 初始化 RechargeOption 路由信息
func (s *RechargeOptionRouter) InitRechargeOptionRouter(Router *gin.RouterGroup) {
	rechargeOptionRouter := Router.Group("rechargeOption").Use(middleware.OperationRecord())
	rechargeOptionRouterWithoutRecord := Router.Group("rechargeOption")
	h5RechargeOptionRouterWithoutRecord := Router.Group(global.GetAppApi() + "rechargeOption")
	var rechargeOptionApi = v1.ApiGroupApp.ClothingApiGroup.RechargeOptionApi
	{
		rechargeOptionRouter.POST("createRechargeOption", rechargeOptionApi.CreateRechargeOption)             // 新建RechargeOption
		rechargeOptionRouter.DELETE("deleteRechargeOption", rechargeOptionApi.DeleteRechargeOption)           // 删除RechargeOption
		rechargeOptionRouter.DELETE("deleteRechargeOptionByIds", rechargeOptionApi.DeleteRechargeOptionByIds) // 批量删除RechargeOption
		rechargeOptionRouter.PUT("updateRechargeOption", rechargeOptionApi.UpdateRechargeOption)              // 更新RechargeOption
	}
	{
		rechargeOptionRouterWithoutRecord.GET("findRechargeOption", rechargeOptionApi.FindRechargeOption)       // 根据ID获取RechargeOption
		rechargeOptionRouterWithoutRecord.GET("getRechargeOptionList", rechargeOptionApi.GetRechargeOptionList) // 获取RechargeOption列表
	}
	{
		h5RechargeOptionRouterWithoutRecord.GET("getRechargeOptionList", rechargeOptionApi.GetRechargeOptionList) // 获取RechargeOption列表
	}
}
