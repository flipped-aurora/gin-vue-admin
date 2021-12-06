package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)             // 新建SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
