package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysDataAccessLogRouter struct{}

func (s *SysDataAccessLogRouter) InitSysDataAccessLogRouter(Router *gin.RouterGroup) {
	logRouter := Router.Group("dataAccessLog").Use(middleware.OperationRecord())
	logRouterWithoutRecord := Router.Group("dataAccessLog")
	{
		logRouter.DELETE("deleteDataAccessLogByIds", dataAccessLogApi.DeleteDataAccessLogByIds) // 批量删除审计日志
	}
	{
		logRouterWithoutRecord.POST("getDataAccessLogList", dataAccessLogApi.GetDataAccessLogList) // 分页获取审计日志
	}
}
