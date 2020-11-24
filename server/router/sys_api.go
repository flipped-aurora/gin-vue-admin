package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("createApi", v1.CreateApi)   // 创建Api
		ApiRouter.POST("deleteApi", v1.DeleteApi)   // 删除Api
		ApiRouter.POST("getApiList", v1.GetApiList) // 获取Api列表
		ApiRouter.POST("getApiById", v1.GetApiById) // 获取单条Api消息
		ApiRouter.POST("updateApi", v1.UpdateApi)   // 更新api
		ApiRouter.POST("getAllApis", v1.GetAllApis) // 获取所有api
	}
}
