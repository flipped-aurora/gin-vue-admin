package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	// 初始化通知中心路由
	noticeGroup := private.Group("notice")
	{
		// 初始化通知管理路由
		router.RouterGroupApp.NotificationRouter.InitNotificationRouter(noticeGroup)

		// 初始化在线用户管理路由
		router.RouterGroupApp.OnlineUserRouter.InitOnlineUserRouter(noticeGroup)

		// 初始化Socket.IO路由
		router.RouterGroupApp.SocketRouter.InitSocketRouter(public)
	}
}
