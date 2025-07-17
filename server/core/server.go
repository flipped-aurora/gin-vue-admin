package core

import (
	"fmt"

	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

func RunServer() {
	if global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}

	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	fmt.Printf(`

		%s gin-vue-admin
		%s:v2.8.2
		%s
		%s：https://github.com/flipped-aurora/gin-vue-admin
		%s:https://plugin.gin-vue-admin.com
		%s:https://support.qq.com/products/371961
		%s:http://127.0.0.1%s/swagger/index.html
		默认MCP SSE地址:http://127.0.0.1%s%s
		默认MCP Message地址:http://127.0.0.1%s%s
		%s:http://127.0.0.1:8080
		%s
		%s
		%s
		%s
`, global.Translate("core.server.welcomeTo"),
		global.Translate("core.server.currentVersion"),
		global.Translate("core.server.joinGroup"),
		global.Translate("core.server.website"),
		global.Translate("core.server.pluginMarket"),
		global.Translate("core.server.community"),
		global.Translate("core.server.swagger"),
		address, address, global.GVA_CONFIG.MCP.SSEPath, address, global.GVA_CONFIG.MCP.MessagePath,
		global.Translate("core.server.frontend"),
		global.Translate("core.server.copyright1"),
		global.Translate("core.server.copyright2"),
		global.Translate("core.server.copyright3"),
		global.Translate("core.server.copyright4"))

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
