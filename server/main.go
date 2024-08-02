package main

import (
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/translate"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.0
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // Initializing viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // Initializing the zap log library
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // Conneting to database using gorm
	initialize.Timer()
	initialize.DBList()
	// added by mohamed hassan to support multilanguage
	global.GVA_TRANSLATOR = translate.Translator{} // create translator inestance  here
	global.GVA_TRANSLATOR.InitTranslator(global.GVA_CONFIG.Language.Language, global.GVA_CONFIG.Language.Dir)
	// end of adding
	if global.GVA_DB != nil {
		initialize.RegisterTables() // Initializing database tables
		// defer close the database connection before the end of the program
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
