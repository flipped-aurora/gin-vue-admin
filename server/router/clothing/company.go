package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyRouter struct {
}

// InitCompanyRouter 初始化 Company 路由信息
func (s *CompanyRouter) InitCompanyRouter(Router *gin.RouterGroup) {
	companyRouter := Router.Group("company").Use(middleware.OperationRecord())
	companyRouterWithoutRecord := Router.Group("company")
	var companyApi = v1.ApiGroupApp.ClothingApiGroup.CompanyApi
	{
		companyRouter.POST("createCompany", companyApi.CreateCompany)   // 新建Company
		companyRouter.DELETE("deleteCompany", companyApi.DeleteCompany) // 删除Company
		companyRouter.DELETE("deleteCompanyByIds", companyApi.DeleteCompanyByIds) // 批量删除Company
		companyRouter.PUT("updateCompany", companyApi.UpdateCompany)    // 更新Company
	}
	{
		companyRouterWithoutRecord.GET("findCompany", companyApi.FindCompany)        // 根据ID获取Company
		companyRouterWithoutRecord.GET("getCompanyList", companyApi.GetCompanyList)  // 获取Company列表
	}
}
