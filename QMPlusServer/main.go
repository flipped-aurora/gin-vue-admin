package main

import (
	"gin-vue-admin/cmd"
	"gin-vue-admin/config"
	"gin-vue-admin/init/initRedis"
	"gin-vue-admin/init/initRouter"
	"gin-vue-admin/init/qmlog"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/init/registTable"
	"runtime"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

func main() {

	qmlog.InitLog()                                            // 初始化日志
	db := qmsql.InitMysql(config.GinVueAdminconfig.MysqlAdmin) // 链接初始化数据库
	if config.GinVueAdminconfig.System.UseMultipoint {
		_ = initRedis.InitRedis() // 初始化redis服务
	}
	registTable.RegistTable(db)       // 注册数据库表
	defer qmsql.DEFAULTDB.Close()     // 程序结束前关闭数据库链接
	Router := initRouter.InitRouter() // 注册路由
	qmlog.QMLog.Info("服务器开启")         // 日志测试代码
	//Router.RunTLS(":443","ssl.pem", "ssl.key")  // https支持 需要添加中间件
	sysType := runtime.GOOS

	if sysType == "linux" {
		// LINUX系统
		//	借助endless开发无感知重启后台 以及前端接口重启后台功能
	}
	if sysType == "windows" {
		// WIN系统
		cmd.RunWindowsServer(Router)
	}
}
