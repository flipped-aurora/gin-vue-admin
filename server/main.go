package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/translate"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.8.4
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化系统
	initializeSystem()
	// 运行服务器
	core.RunServer()
}

// initializeSystem 初始化系统所有组件
// 提取为单独函数以便于系统重载时调用
func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // Initializing the zap log library
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // Conneting to database using gorm
	initialize.Timer()
	initialize.DBList()

	// added by mohamed hassan to support multilanguage
	global.GVA_TRANSLATOR = translate.Translator{} // create translator inestance  here
	//global.GVA_TRANSLATOR.InitTranslator(global.GVA_CONFIG.Language.Language, global.GVA_CONFIG.Language.Dir)
	global.GVA_TRANSLATOR.InitTranslatorEx(global.GVA_CONFIG.Language.Language, global.GVA_CONFIG.Language.DefaultLanguage, global.GVA_CONFIG.Language.Dir)
	// end of adding

	initialize.SetupHandlers() // 注册全局函数
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
