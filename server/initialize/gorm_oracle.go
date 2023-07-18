package initialize

import (
	//"github.com/dzwvip/oracle"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"

	//_ "github.com/godror/godror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormOracle 初始化oracle数据库
// 如果需要Oracle库 放开import里的注释 把下方 mysql.Config 改为 oracle.Config ;  mysql.New 改为 oracle.New
func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormOracleByConfig 初始化Oracle数据库用过传入配置
func GormOracleByConfig(m config.Oracle) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
