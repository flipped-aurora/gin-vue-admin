package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CroppingRecordRouter struct {
}

// InitCroppingRecordRouter 初始化 CroppingRecord 路由信息
func (s *CroppingRecordRouter) InitCroppingRecordRouter(Router *gin.RouterGroup) {
	croppingRecordRouter := Router.Group("croppingRecord").Use(middleware.OperationRecord())
	croppingRecordRouterWithoutRecord := Router.Group("croppingRecord")
	h5CroppingRecordRouterWithoutRecord := Router.Group(global.GetAppApi() + "croppingRecord")
	var croppingRecordApi = v1.ApiGroupApp.ClothingApiGroup.CroppingRecordApi
	{
		croppingRecordRouter.POST("createCroppingRecord", croppingRecordApi.CreateCroppingRecord)             // 新建CroppingRecord
		croppingRecordRouter.DELETE("deleteCroppingRecord", croppingRecordApi.DeleteCroppingRecord)           // 删除CroppingRecord
		croppingRecordRouter.DELETE("deleteCroppingRecordByIds", croppingRecordApi.DeleteCroppingRecordByIds) // 批量删除CroppingRecord
		croppingRecordRouter.PUT("updateCroppingRecord", croppingRecordApi.UpdateCroppingRecord)              // 更新CroppingRecord
	}
	{
		croppingRecordRouterWithoutRecord.GET("findCroppingRecord", croppingRecordApi.FindCroppingRecord)       // 根据ID获取CroppingRecord
		croppingRecordRouterWithoutRecord.GET("getCroppingRecordList", croppingRecordApi.GetCroppingRecordList) // 获取CroppingRecord列表
	}
	{
		h5CroppingRecordRouterWithoutRecord.POST("createCroppingRecord", croppingRecordApi.CreateCroppingRecord)  // 新建CroppingRecord
		h5CroppingRecordRouterWithoutRecord.PUT("updateCroppingRecord", croppingRecordApi.UpdateCroppingRecord)   // 更新CroppingRecord
		h5CroppingRecordRouterWithoutRecord.GET("findCroppingRecord", croppingRecordApi.FindCroppingRecord)       // 根据ID获取CroppingRecord
		h5CroppingRecordRouterWithoutRecord.GET("getCroppingRecordList", croppingRecordApi.GetCroppingRecordList) // 获取CroppingRecord列表
	}
}
