// Package core 提供核心服务器启动功能
//
// 功能说明：
// 本包负责启动HTTP服务器，包括：
// 1. 初始化Redis（如果启用）
// 2. 初始化MongoDB（如果启用）
// 3. 加载JWT数据到内存
// 4. 初始化路由
// 5. 启动HTTP服务器
package core

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

// RunServer 启动HTTP服务器
//
// 功能说明：
// 这是服务器启动的入口函数，在main.go中调用。
// 按照配置初始化各种服务，然后启动HTTP服务器监听请求。
//
// 启动流程：
// 1. 根据配置初始化Redis（如果启用）
// 2. 根据配置初始化MongoDB（如果启用）
// 3. 从数据库加载JWT黑名单等数据到内存
// 4. 初始化路由（注册所有API路由）
// 5. 打印启动信息和访问地址
// 6. 启动HTTP服务器
//
// 使用示例：
//
//	// 在main.go中调用
//	func main() {
//	    initializeSystem()
//	    core.RunServer() // 启动服务器
//	}
//
// 注意事项：
// - 此函数会阻塞，直到服务器关闭
// - 服务器启动前需要确保所有初始化已完成
func RunServer() {
	// 步骤1: 初始化Redis服务（如果配置启用）
	if global.GVA_CONFIG.System.UseRedis {
		// 初始化默认Redis连接
		initialize.Redis()
		// 如果配置了多Redis实例，初始化所有Redis连接
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}

	// 步骤2: 初始化MongoDB服务（如果配置启用）
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			// MongoDB初始化失败只记录日志，不中断启动
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}

	// 步骤3: 从数据库加载JWT相关数据到内存
	// 包括JWT黑名单、权限数据等，提高查询性能
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	// 步骤4: 初始化路由
	// 注册所有API路由、中间件等
	Router := initialize.Routers()

	// 步骤5: 构建服务器地址
	// 格式：:端口号，如 :8888
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	// 步骤6: 打印启动信息和访问地址
	// 包括版本信息、文档地址、MCP地址等
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

	// 步骤7: 启动HTTP服务器
	// 参数说明：
	// - address: 监听地址
	// - Router: Gin路由引擎
	// - 10*time.Minute: 读取超时时间
	// - 10*time.Minute: 写入超时时间
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
