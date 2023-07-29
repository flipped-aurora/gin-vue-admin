package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg"

	"github.com/flipped-aurora/gin-vue-admin/server/model/AerialPhotographyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/UserTeemlinkPkg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, NestInfo.NestInfo{}, Nestrolepkg.NestRole{}, NestExecRecordPkg.NestExecRecord{}, FlyResultPkg.FlyResult{}, UserTeemlinkPkg.UserTeemlink{}, NestAirlinePkg.NestAirline{}, AerialPhotographyResultPkg.AerialPhotographyResult{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
