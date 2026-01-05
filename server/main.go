// Package main gin-vue-admin 后端服务入口文件
//
// 功能说明：
// 这是gin-vue-admin后端服务的程序入口，负责启动整个应用。
// 主要功能包括：
// 1. 初始化系统配置（Viper）
// 2. 初始化日志系统（Zap）
// 3. 初始化数据库连接（GORM）
// 4. 初始化定时任务
// 5. 初始化路由和中间件
// 6. 启动HTTP服务器
//
// 启动流程：
// main() -> initializeSystem() -> core.RunServer()
package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// Swagger文档配置
// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.8.7
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 步骤1: 初始化系统所有组件
	// 包括配置、日志、数据库、定时任务等
	initializeSystem()

	// 步骤2: 运行HTTP服务器
	// 启动Gin服务器，监听配置的端口
	core.RunServer()
}

// initializeSystem 初始化系统所有组件
//
// 功能说明：
// 按照依赖顺序初始化系统的各个组件，确保组件之间的依赖关系正确。
// 提取为单独函数以便于系统重载时调用（如配置热更新后重新初始化）。
//
// 初始化顺序：
// 1. Viper配置管理器（读取config.yaml）
// 2. 其他初始化（如验证器、翻译器等）
// 3. Zap日志库（用于记录日志）
// 4. GORM数据库连接（连接MySQL/PostgreSQL等）
// 5. 定时任务管理器
// 6. 多数据库列表初始化
// 7. 全局函数注册（如模板函数）
// 8. 数据库表初始化（自动创建表结构）
//
// 使用示例：
//
//	// 在main函数中调用
//	initializeSystem()
//
//	// 配置热更新后重新初始化
//	initializeSystem()
//
// 注意事项：
// - 初始化顺序很重要，不能随意调整
// - 如果某个组件初始化失败，后续组件可能无法正常工作
func initializeSystem() {
	// 步骤1: 初始化Viper配置管理器
	// 读取config.yaml配置文件，支持环境变量覆盖
	global.GVA_VP = core.Viper()

	// 步骤2: 初始化其他组件（验证器、翻译器等）
	initialize.OtherInit()

	// 步骤3: 初始化zap日志库
	// 配置日志格式、输出位置、日志级别等
	global.GVA_LOG = core.Zap()
	// 替换zap的全局logger，使zap.L()等全局函数可用
	zap.ReplaceGlobals(global.GVA_LOG)

	// 步骤4: 初始化GORM数据库连接
	// 根据配置连接MySQL/PostgreSQL/SQLite等数据库
	global.GVA_DB = initialize.Gorm()

	// 步骤5: 初始化定时任务管理器
	// 用于执行定时任务，如清理过期数据等
	initialize.Timer()

	// 步骤6: 初始化多数据库列表
	// 如果配置了多个数据库，初始化所有数据库连接
	initialize.DBList()

	// 步骤7: 注册全局函数
	// 注册模板函数、工具函数等，供全局使用
	initialize.SetupHandlers()

	// 步骤8: 初始化数据库表
	// 如果数据库连接成功，自动创建或更新表结构
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
