package autocode

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type AutoCodeExampleRouter struct {
}

func (s *AutoCodeExampleRouter) InitSysAutoCodeExampleRouter(Router *gin.RouterGroup) {
	autoCodeExampleRouter := Router.Group("autoCodeExample").Use(middleware.OperationRecord())
	var autoCodeExampleApi = v1.ApiGroupApp.AutoCodeApiGroup.AutoCodeExampleApi
	{
		autoCodeExampleRouter.POST("createSysAutoCodeExample", autoCodeExampleApi.CreateAutoCodeExample)   // 新建AutoCodeExample
		autoCodeExampleRouter.DELETE("deleteSysAutoCodeExample", autoCodeExampleApi.DeleteAutoCodeExample) // 删除AutoCodeExample
		autoCodeExampleRouter.PUT("updateSysAutoCodeExample", autoCodeExampleApi.UpdateAutoCodeExample)    // 更新AutoCodeExample
		autoCodeExampleRouter.GET("findSysAutoCodeExample", autoCodeExampleApi.FindAutoCodeExample)        // 根据ID获取AutoCodeExample
		autoCodeExampleRouter.GET("getSysAutoCodeExampleList", autoCodeExampleApi.GetAutoCodeExampleList)  // 获取AutoCodeExample列表
	}
}
