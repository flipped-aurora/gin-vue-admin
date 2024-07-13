package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/router"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRouters(group *gin.Engine) {
	PublicGroup := group.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := group.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		router.TodoListRouter.InitRouter(PrivateGroup, PublicGroup)
	}
}
