package initialize

import (
	oracle "github.com/dzwvip/gorm-oracle"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"gorm.io/gorm"
)

// GormOracle 初始化oracle数据库
func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	return initOracleDatabase(m)
}

// GormOracleByConfig 初始化Oracle数据库用过传入配置
func GormOracleByConfig(m config.Oracle) *gorm.DB {
	return initOracleDatabase(m)
}

// initOracleDatabase 初始化Oracle数据库的辅助函数
func initOracleDatabase(m config.Oracle) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(oracle.Open(m.Dsn()), internal.Gorm.Config(general)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
