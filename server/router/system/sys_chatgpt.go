package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChatGptRouter struct{}

func (s *ChatGptRouter) InitChatGptRouter(Router *gin.RouterGroup) {
	chatGptRouter := Router.Group("chatGpt").Use(middleware.OperationRecord())
	chatGptRouterWithoutRecord := Router.Group("chatGpt")
	chatGptApi := v1.ApiGroupApp.SystemApiGroup.ChatGptApi
	{
		chatGptRouter.POST("createToken", chatGptApi.CreateToken)
	}
	{
		chatGptRouterWithoutRecord.POST("getTable", chatGptApi.GetTable)
	}
}
