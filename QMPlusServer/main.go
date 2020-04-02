package main

import (
	"gin-vue-admin/cmd"
	"gin-vue-admin/config"
	"gin-vue-admin/init/initRedis"
	"gin-vue-admin/init/initRouter"
	"gin-vue-admin/init/initlog/qmlog"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/init/registTerable"
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
	mysqlConfig := config.GinVueAdminconfig.MysqlAdmin
	if mysqlHost != "" && mysqlPort != "" {
		mysqlConfig.Path = mysqlHost + ":" + mysqlPort
	}
	// 链接初始化数据库
	db := qmsql.InitMysql(mysqlConfig) // 链接初始化数据库
	if config.GinVueAdminconfig.System.UseMultipoint {
		// 初始化redis服务
		_ = initRedis.InitRedis()
	}
	// 注册数据库表
	registerTable.RegisterTable(db)
	// 程序结束前关闭数据库链接
	defer qmsql.DEFAULTDB.Close()
	// 注册路由
	Router := initRouter.InitRouter()

	Router.Static("/form-generator", "./static/form-generator")
	//Router.RunTLS(":443","ssl.pem", "ssl.key")  // https支持 需要添加中间件
	//sysType := runtime.GOOS
	//
	//if sysType == "linux" {
	//	// LINUX系统
	//	//	借助endless开发无感知重启后台 以及前端接口重启后台功能
	//}
	//if sysType == "windows" {
	// WIN系统
	cmd.RunWindowsServer(Router)
	//}
}
