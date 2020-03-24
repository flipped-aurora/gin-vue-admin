package main

import (
	"fmt"
	"gin-vue-admin/cmd"
	"gin-vue-admin/config"
	"gin-vue-admin/init/initRedis"
	"gin-vue-admin/init/initRouter"
	"gin-vue-admin/init/log/qmlog"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/init/registTable"
	//"runtime"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

func main() {
	var err error
	logger, err := qmlog.NewLogger()
	if err != nil {
		fmt.Println(err)
	}
	// 链接初始化数据库
	db := qmsql.InitMysql(config.GinVueAdminconfig.MysqlAdmin, logger)
	if config.GinVueAdminconfig.System.UseMultipoint {
		// 初始化redis服务
		_ = initRedis.InitRedis(logger)
	}
	// 注册数据库表
	registTable.RegistTable(db, logger)
	// 程序结束前关闭数据库链接
	defer qmsql.DEFAULTDB.Close()
	// 注册路由
	Router := initRouter.InitRouter(logger)
	//Router.RunTLS(":443","ssl.pem", "ssl.key")  // https支持 需要添加中间件
	//sysType := runtime.GOOS
	//
	//if sysType == "linux" {
	//	// LINUX系统
	//	//	借助endless开发无感知重启后台 以及前端接口重启后台功能
	//}
	//if sysType == "windows" {
	// WIN系统
	cmd.RunWindowsServer(Router, logger)
	//}
}
