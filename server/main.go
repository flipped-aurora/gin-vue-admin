package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"log"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.6.2
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	// 屎山代码临时用 start 莫介意
	defer global.RecordDB.Close()
	rootPath := global.GVA_CONFIG.AutoCode.Root
	rmFilePathRecord := filepath.Join(rootPath, "rm_file", "rm_record.db")
	record_db, err := sql.Open("sqlite3", rmFilePathRecord)
	if err != nil {
		log.Fatal(err)
	}

	_, err = record_db.Exec("CREATE TABLE IF NOT EXISTS records (path TEXT, file TEXT, UPDATE_TIME DATETIME)")
	if err != nil {
		log.Fatal(err)
	}

	global.RecordDB = record_db
	// 屎山代码临时用 end 莫介意

	core.RunWindowsServer()
}
