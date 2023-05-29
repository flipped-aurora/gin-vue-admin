package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComputationRouter struct {
}

// InitComputationRouter 初始化 Computation 路由信息
func (s *ComputationRouter) InitComputationRouter(Router *gin.RouterGroup) {
	computationRouter := Router.Group("computation").Use(middleware.OperationRecord())
	computationRouterWithoutRecord := Router.Group("computation")
	h5ComputationRouterWithoutRecord := Router.Group(global.GetAppApi() + "computation")
	var computationApi = v1.ApiGroupApp.ClothingApiGroup.ComputationApi
	{
		computationRouter.POST("createComputation", computationApi.CreateComputation)             // 新建Computation
		computationRouter.DELETE("deleteComputation", computationApi.DeleteComputation)           // 删除Computation
		computationRouter.DELETE("deleteComputationByIds", computationApi.DeleteComputationByIds) // 批量删除Computation
		computationRouter.PUT("updateComputation", computationApi.UpdateComputation)              // 更新Computation
	}
	{
		computationRouterWithoutRecord.GET("findComputation", computationApi.FindComputation)       // 根据ID获取Computation
		computationRouterWithoutRecord.GET("getComputationList", computationApi.GetComputationList) // 获取Computation列表
	}
	{
		h5ComputationRouterWithoutRecord.POST("doComputation", computationApi.DoComputation) // 计算
	}
}
