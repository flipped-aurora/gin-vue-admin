package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", dictionaryDetailApi.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", dictionaryDetailApi.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", dictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", dictionaryDetailApi.FindSysDictionaryDetail)           // 根据ID获取SysDictionaryDetail
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", dictionaryDetailApi.GetSysDictionaryDetailList)     // 获取SysDictionaryDetail列表
		dictionaryDetailRouterWithoutRecord.GET("getDictionaryTreeList", dictionaryDetailApi.GetDictionaryTreeList)               // 获取字典详情树形结构
		dictionaryDetailRouterWithoutRecord.GET("getDictionaryTreeListByType", dictionaryDetailApi.GetDictionaryTreeListByType)   // 根据字典类型获取字典详情树形结构
		dictionaryDetailRouterWithoutRecord.GET("getDictionaryDetailsByParent", dictionaryDetailApi.GetDictionaryDetailsByParent) // 根据父级ID获取字典详情
		dictionaryDetailRouterWithoutRecord.GET("getDictionaryPath", dictionaryDetailApi.GetDictionaryPath)                       // 获取字典详情的完整路径
	}
}
