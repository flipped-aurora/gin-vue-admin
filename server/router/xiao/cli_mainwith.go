package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliMainwithRouter struct{}

// InitCliMainwithRouter 初始化 提币总表 路由信息
func (s *CliMainwithRouter) InitCliMainwithRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	climainwithRouter := Router.Group("climainwith").Use(middleware.OperationRecord())
	climainwithRouterWithoutRecord := Router.Group("climainwith")
	climainwithRouterWithoutAuth := PublicRouter.Group("climainwith")
	{
		climainwithRouter.POST("createCliMainwith", climainwithApi.CreateCliMainwith)             // 新建提币总表
		climainwithRouter.DELETE("deleteCliMainwith", climainwithApi.DeleteCliMainwith)           // 删除提币总表
		climainwithRouter.DELETE("deleteCliMainwithByIds", climainwithApi.DeleteCliMainwithByIds) // 批量删除提币总表
		climainwithRouter.PUT("updateCliMainwith", climainwithApi.UpdateCliMainwith)              // 更新提币总表
	}
	{
		climainwithRouterWithoutRecord.GET("findCliMainwith", climainwithApi.FindCliMainwith)       // 根据ID获取提币总表
		climainwithRouterWithoutRecord.GET("getCliMainwithList", climainwithApi.GetCliMainwithList) // 获取提币总表列表
	}
	{
		climainwithRouterWithoutAuth.GET("getCliMainwithPublic", climainwithApi.GetCliMainwithPublic) // 提币总表开放接口
	}
}
