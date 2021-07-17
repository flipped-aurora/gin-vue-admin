package system

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryDetailRouter struct {
}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	var sysDictionaryDetailApi = v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", sysDictionaryDetailApi.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", sysDictionaryDetailApi.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", sysDictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
		dictionaryDetailRouter.GET("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)        // 根据ID获取SysDictionaryDetail
		dictionaryDetailRouter.GET("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList)  // 获取SysDictionaryDetail列表
	}
}
