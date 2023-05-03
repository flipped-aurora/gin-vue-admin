package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProcessRouter struct {
}

// InitProcessRouter 初始化 Process 路由信息
func (s *ProcessRouter) InitProcessRouter(Router *gin.RouterGroup) {
	processRouter := Router.Group("process").Use(middleware.OperationRecord())
	processRouterWithoutRecord := Router.Group("process")
	var processApi = v1.ApiGroupApp.ClothingApiGroup.ProcessApi
	{
		processRouter.POST("createProcess", processApi.CreateProcess)   // 新建Process
		processRouter.DELETE("deleteProcess", processApi.DeleteProcess) // 删除Process
		processRouter.DELETE("deleteProcessByIds", processApi.DeleteProcessByIds) // 批量删除Process
		processRouter.PUT("updateProcess", processApi.UpdateProcess)    // 更新Process
	}
	{
		processRouterWithoutRecord.GET("findProcess", processApi.FindProcess)        // 根据ID获取Process
		processRouterWithoutRecord.GET("getProcessList", processApi.GetProcessList)  // 获取Process列表
	}
}
