package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	//global.GVA_LOG = core.Zap() // 初始化zap日志库
	//zap.ReplaceGlobals(global.GVA_LOG)
	//以下为zao的包装方法，实现打印trace信息,使用的时候将上两行注释掉
	//同时需要将global.GVA_LOG 的数据类型改成*internal.LogWrapper
	global.GVA_LOG = core.ZapWrapper() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG.ZapLogger)

	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
