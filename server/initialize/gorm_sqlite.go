package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
	s := global.GVA_CONFIG.Sqlite
	return initSqliteDatabase(s)
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func GormSqliteByConfig(s config.Sqlite) *gorm.DB {
	return initSqliteDatabase(s)
}

// initSqliteDatabase 初始化Sqlite数据库辅助函数
func initSqliteDatabase(s config.Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	// 数据库配置
	general := s.GeneralDB
	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(general)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
