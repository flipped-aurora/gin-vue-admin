package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/init"
	"gin-vue-admin/init/qmlog"
	"os"
	//"runtime"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

var (
	mysqlHost = os.Getenv("MYSQLHOST")
	mysqlPort = os.Getenv("MYSQLPORT")
)

func main() {
	if err := qmlog.NewLogger(); err != nil {
		panic(err)
	}
	// 可以通过环境变量来覆盖配置值
	// 未设定有效的环境变量时，使用配置值
	mysqlConfig := init.GinVueAdminconfig.MysqlAdmin
	// 链接初始化数据库
	init.RegisterMysql(mysqlConfig) // 链接初始化数据库

	// 注册数据库表
	init.RegisterTable(global.GVA_DB)
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()

	core.RunWindowsServer()
}
