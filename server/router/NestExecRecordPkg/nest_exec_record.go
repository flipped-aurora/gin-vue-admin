package NestExecRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NestExecRecordRouter struct {
}

// InitNestExecRecordRouter 初始化 NestExecRecord 路由信息
func (s *NestExecRecordRouter) InitNestExecRecordRouter(Router *gin.RouterGroup) {
	NtERecordRouter := Router.Group("NtERecord").Use(middleware.OperationRecord())
	NtERecordRouterWithoutRecord := Router.Group("NtERecord")
	var NtERecordApi = v1.ApiGroupApp.NestExecRecordPkgApiGroup.NestExecRecordApi
	{
		NtERecordRouter.POST("createNestExecRecord", NtERecordApi.CreateNestExecRecord)                      // 新建NestExecRecord
		NtERecordRouter.DELETE("deleteNestExecRecord", NtERecordApi.DeleteNestExecRecord)                    // 删除NestExecRecord
		NtERecordRouter.DELETE("deleteNestExecRecordByIds", NtERecordApi.DeleteNestExecRecordByIds)          // 批量删除NestExecRecord
		NtERecordRouter.PUT("updateNestExecRecord", NtERecordApi.UpdateNestExecRecord)                       // 更新NestExecRecord
		NtERecordRouter.PUT("updateNestExecRecordWithAirline", NtERecordApi.UpdateNestExecRecordWithAirline) // 更新NestExecRecord及关联的航线数据
	}
	{
		NtERecordRouterWithoutRecord.GET("findNestExecRecord", NtERecordApi.FindNestExecRecord)       // 根据ID获取NestExecRecord
		NtERecordRouterWithoutRecord.GET("getNestExecRecordList", NtERecordApi.GetNestExecRecordList) // 获取NestExecRecord列表
	}
}
