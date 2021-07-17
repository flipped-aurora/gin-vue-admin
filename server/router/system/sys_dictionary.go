package system

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct {
}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	var sysDictionaryApi = v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.POST("createSysDictionary", sysDictionaryApi.CreateSysDictionary)   // 新建SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary) // 删除SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", sysDictionaryApi.UpdateSysDictionary)    // 更新SysDictionary
		sysDictionaryRouter.GET("findSysDictionary", sysDictionaryApi.FindSysDictionary)        // 根据ID获取SysDictionary
		sysDictionaryRouter.GET("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}
