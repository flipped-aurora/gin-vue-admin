package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyApplyRouter struct {
}

// InitCompanyApplyRouter 初始化 CompanyApply 路由信息
func (s *CompanyApplyRouter) InitCompanyApplyRouter(Router *gin.RouterGroup) {
	companyApplyRouter := Router.Group("companyApply").Use(middleware.OperationRecord())
	companyApplyRouterWithoutRecord := Router.Group("companyApply")
	h5CompanyApplyRouterWithoutRecord := Router.Group(global.GetAppApi() + "companyApply")
	var companyApplyApi = v1.ApiGroupApp.ClothingApiGroup.CompanyApplyApi
	{
		companyApplyRouter.POST("createCompanyApply", companyApplyApi.CreateCompanyApply)             // 新建CompanyApply
		companyApplyRouter.DELETE("deleteCompanyApply", companyApplyApi.DeleteCompanyApply)           // 删除CompanyApply
		companyApplyRouter.DELETE("deleteCompanyApplyByIds", companyApplyApi.DeleteCompanyApplyByIds) // 批量删除CompanyApply
		companyApplyRouter.PUT("updateCompanyApply", companyApplyApi.UpdateCompanyApply)              // 更新CompanyApply
	}
	{
		companyApplyRouterWithoutRecord.GET("findCompanyApply", companyApplyApi.FindCompanyApply)       // 根据ID获取CompanyApply
		companyApplyRouterWithoutRecord.GET("getCompanyApplyList", companyApplyApi.GetCompanyApplyList) // 获取CompanyApply列表
	}
	{
		h5CompanyApplyRouterWithoutRecord.PUT("optApply", companyApplyApi.OptApply) // 审核申请
	}
}
