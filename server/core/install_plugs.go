package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/piexlmax/gvaplug"
)

func InstallPlugs(route *gin.Engine) {
	GvaPlugsGroupBase := route.Group("") // 纯净路由 插件可通用 传递给插件的0号位置 为了不让插件的自动路由出问题 这里建议为 ""
	/**安装一个插件需要的步骤 start**/
	GvaPlugsGroupMiddle := route.Group("") // 携带中间件的路由 传递给插件的1号位置 为了不让插件的自动路由出问题 这里建议为 ""
	GvaPlugsGroupMiddle.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// 插件安装 暂时只是后台功能 添加model 添加路由 添加对数据库的操作  详细插件测试模板可看https://github.com/piexlmax/gvaplug  此处不建议投入生产
	err := initialize.InstallPlug(global.GVA_DB, [2]*gin.RouterGroup{
		GvaPlugsGroupBase,
		GvaPlugsGroupMiddle,
	}, gvaplug.GvaPlug{SomeConfig: "插件给用户提供的配置区域"})
	if err != nil {
		panic(fmt.Sprintf("插件安装失败： %v", err))
	}
	/**安装一个插件需要的步骤 end**/
}
