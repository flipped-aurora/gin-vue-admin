package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	//"gin-vue-admin/gva/init_data"
	"gin-vue-admin/initialize"
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
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	// case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		initialize.Mysql()
	}
	initialize.DBTables()
	//init_data.InitData() // 打开注释即可初始化数据  使用过后一定要再次注释或删除
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	core.RunWindowsServer()
}
