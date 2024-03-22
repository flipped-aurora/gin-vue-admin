package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MessageRouter struct {
}

// InitMessageRouter 初始化 Message 路由信息
func (s *MessageRouter) InitMessageRouter(Router *gin.RouterGroup) {
	messageRouter := Router.Group("message").Use(middleware.OperationRecord())
	messageRouterWithoutRecord := Router.Group("message")
	var messageApi = v1.ApiGroupApp.WebcmsApiGroup.MessageApi
	{
		messageRouter.POST("createMessage", messageApi.CreateMessage)             // 新建Message
		messageRouter.DELETE("deleteMessage", messageApi.DeleteMessage)           // 删除Message
		messageRouter.DELETE("deleteMessageByIds", messageApi.DeleteMessageByIds) // 批量删除Message
		messageRouter.PUT("updateMessage", messageApi.UpdateMessage)              // 更新Message
	}
	{
		messageRouterWithoutRecord.GET("findMessage", messageApi.FindMessage)       // 根据ID获取Message
		messageRouterWithoutRecord.GET("getMessageList", messageApi.GetMessageList) // 获取Message列表
	}
}
