package initialize

import (
	"github.com/dzwvip/oracle"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	_ "github.com/godror/godror"
	"gorm.io/gorm"
)

// GormOracle 初始化oracle数据库

func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := oracle.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度

	}
	if db, err := gorm.Open(oracle.New(oracleConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormOracleByConfig 初始化Mysql数据库用过传入配置
func GormOracleByConfig(m config.Oracle) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := oracle.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度

	}
	if db, err := gorm.Open(oracle.New(oracleConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
