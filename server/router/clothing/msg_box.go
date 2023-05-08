package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsgBoxRouter struct {
}

// InitMsgBoxRouter 初始化 MsgBox 路由信息
func (s *MsgBoxRouter) InitMsgBoxRouter(Router *gin.RouterGroup) {
	msgBoxRouter := Router.Group("msgBox").Use(middleware.OperationRecord())
	msgBoxRouterWithoutRecord := Router.Group("msgBox")
	h5MsgBoxRouterWithoutRecord := Router.Group(global.GetAppApi() + "msgBox")
	var msgBoxApi = v1.ApiGroupApp.ClothingApiGroup.MsgBoxApi
	{
		msgBoxRouter.POST("createMsgBox", msgBoxApi.CreateMsgBox)             // 新建MsgBox
		msgBoxRouter.DELETE("deleteMsgBox", msgBoxApi.DeleteMsgBox)           // 删除MsgBox
		msgBoxRouter.DELETE("deleteMsgBoxByIds", msgBoxApi.DeleteMsgBoxByIds) // 批量删除MsgBox
		msgBoxRouter.PUT("updateMsgBox", msgBoxApi.UpdateMsgBox)              // 更新MsgBox
	}
	{
		msgBoxRouterWithoutRecord.GET("findMsgBox", msgBoxApi.FindMsgBox)       // 根据ID获取MsgBox
		msgBoxRouterWithoutRecord.GET("getMsgBoxList", msgBoxApi.GetMsgBoxList) // 获取MsgBox列表
	}
	{
		h5MsgBoxRouterWithoutRecord.GET("setRead", msgBoxApi.SetRead)                   // 设置已读
		h5MsgBoxRouterWithoutRecord.GET("getMyMsgBoxList", msgBoxApi.GetMyMsgList)      // 分页获取收到的信息
		h5MsgBoxRouterWithoutRecord.GET("getMySendMsgList", msgBoxApi.GetMySendMsgList) // 分页获取发送的信息
	}
}
