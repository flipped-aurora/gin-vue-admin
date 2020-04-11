package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	//"runtime"
)

func main() {
	switch global.GVA_CONFIG.System.Db  {
	case "mysql":
		initialize.Mysql()
	case "sqlite":
		initialize.Sqlite()
	}
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()

	core.RunWindowsServer()
}
