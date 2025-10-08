package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
	"time"
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
	欢迎使用 gin-vue-admin
	当前版本:%s
	加群方式:微信号：shouzi_1994 QQ群：470239250
	项目地址：https://github.com/flipped-aurora/gin-vue-admin
	插件市场:https://plugin.gin-vue-admin.com
	GVA讨论社区:https://support.qq.com/products/371961
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认MCP SSE地址:http://127.0.0.1%s%s
	默认MCP Message地址:http://127.0.0.1%s%s
	默认前端文件运行地址:http://127.0.0.1:8080
	--------------------------------------版权声明--------------------------------------
	** 版权所有方：flipped-aurora开源团队 **
	** 版权持有公司：北京翻转极光科技有限责任公司 **
	** 剔除授权标识需购买商用授权：https://gin-vue-admin.com/empower/index.html **
	** 感谢您对Gin-Vue-Admin的支持与关注 合法授权使用更有利于项目的长久发展**
`, global.Version, address, address, global.GVA_CONFIG.MCP.SSEPath, address, global.GVA_CONFIG.MCP.MessagePath)
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
