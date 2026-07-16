// server/router/system/sys_timed_task.go
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TimedTaskRouter struct{}

func (s *TimedTaskRouter) InitTimedTaskRouter(Router *gin.RouterGroup) {
	timedTaskRouter := Router.Group("timedTask").Use(middleware.OperationRecord())
	timedTaskRouterWithoutRecord := Router.Group("timedTask")
	{
		timedTaskRouter.POST("createTimedTask", timedTaskApi.CreateTimedTask)   // 创建定时任务
		timedTaskRouter.PUT("updateTimedTask", timedTaskApi.UpdateTimedTask)    // 更新定时任务
		timedTaskRouter.DELETE("deleteTimedTask", timedTaskApi.DeleteTimedTask) // 删除定时任务
		timedTaskRouter.POST("toggleTimedTask", timedTaskApi.ToggleTimedTask)   // 启用/停用
		timedTaskRouter.POST("triggerTimedTask", timedTaskApi.TriggerTimedTask) // 手动触发
	}
	{
		timedTaskRouterWithoutRecord.GET("getTimedTaskList", timedTaskApi.GetTimedTaskList)         // 任务列表
		timedTaskRouterWithoutRecord.GET("getTimedTaskLogList", timedTaskApi.GetTimedTaskLogList)   // 执行日志
		timedTaskRouterWithoutRecord.GET("getRegisteredMethods", timedTaskApi.GetRegisteredMethods) // 已注册方法
		// SSE 长连接: 绝不能套 TimeoutMiddleware(utils/sse/hub.go:168 约束)
		timedTaskRouterWithoutRecord.GET("alertStream", timedTaskApi.AlertStream) // 失败告警订阅
	}
}
