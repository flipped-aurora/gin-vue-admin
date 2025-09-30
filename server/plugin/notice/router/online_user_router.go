package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	noticeApi "github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/api"
	"github.com/gin-gonic/gin"
)

type OnlineUserRouter struct{}

// InitOnlineUserRouter 初始化在线用户路由
func (r *OnlineUserRouter) InitOnlineUserRouter(Router *gin.RouterGroup) {
	onlineUserRouter := Router.Group("onlineUser")
	onlineUserRouterWithoutRecord := Router.Group("onlineUser")

	// 需要记录操作日志的路由
	onlineUserRouter.Use(middleware.OperationRecord())
	{
		// 管理员接口
		onlineUserRouter.DELETE("removeOnlineUser/:userId", noticeApi.ApiGroupApp.OnlineUserApi.RemoveOnlineUser) // 移除在线用户
		onlineUserRouter.POST("cleanOfflineUsers", noticeApi.ApiGroupApp.OnlineUserApi.CleanOfflineUsers)         // 清理离线用户
	}

	// 不需要记录操作日志的路由
	onlineUserRouterWithoutRecord.Use(middleware.JWTAuth())
	{
		// 查询接口
		onlineUserRouterWithoutRecord.GET("getOnlineUsers", noticeApi.ApiGroupApp.OnlineUserApi.GetOnlineUsers)                     // 获取在线用户列表
		onlineUserRouterWithoutRecord.POST("getOnlineUsers", noticeApi.ApiGroupApp.OnlineUserApi.GetOnlineUsersWithPagination)      // 获取在线用户列表（分页）
		onlineUserRouterWithoutRecord.GET("getOnlineUsersByRole/:roleId", noticeApi.ApiGroupApp.OnlineUserApi.GetOnlineUsersByRole) // 根据角色获取在线用户
		onlineUserRouterWithoutRecord.GET("checkUserOnline/:userId", noticeApi.ApiGroupApp.OnlineUserApi.CheckUserOnline)           // 检查用户是否在线
		onlineUserRouterWithoutRecord.GET("getOnlineUserCount", noticeApi.ApiGroupApp.OnlineUserApi.GetOnlineUserCount)             // 获取在线用户数量
		onlineUserRouterWithoutRecord.GET("getOnlineUserStats", noticeApi.ApiGroupApp.OnlineUserApi.GetOnlineUserStats)             // 获取在线用户统计数据
	}

}
