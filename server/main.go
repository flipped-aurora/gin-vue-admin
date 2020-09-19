package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/gva/init_data"
	"gin-vue-admin/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.Gorm()
	if global.GVA_CONFIG.System.NeedInitData {
		init_data.InitData() // 通过配置文件初始化数据 默认为 false 首次运行需要将 ./config.yaml中 system下的 need-init-data 修改为true
	}
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	core.RunWindowsServer()
}
