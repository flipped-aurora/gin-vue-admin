package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		if global.GVA_CONFIG.OpenReadWriteSeparation {
			var writeDsn []gorm.Dialector
			var readDsn []gorm.Dialector
			switch m.Duty {
			case "read":
				readDsn = append(readDsn, mysql.Open(m.Dsn()))
			case "write":
				writeDsn = append(writeDsn, mysql.Open(m.Dsn()))
			default:
				return nil
			}
			dbResolverCfg := dbresolver.Config{
				Sources:  writeDsn,
				Replicas: readDsn,
				Policy:   dbresolver.RandomPolicy{}}
			readWritePlugin := dbresolver.Register(dbResolverCfg)
			err = db.Use(readWritePlugin)
			if err != nil {
				return nil
			}
		}
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.DB) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		if global.GVA_CONFIG.OpenReadWriteSeparation {
			var writeDsn []gorm.Dialector
			var readDsn []gorm.Dialector
			switch m.Duty {
			case "read":
				readDsn = append(readDsn, mysql.Open(m.Dsn()))
			case "write":
				writeDsn = append(writeDsn, mysql.Open(m.Dsn()))
			default:
				return nil
			}
			dbResolverCfg := dbresolver.Config{
				Sources:  writeDsn,
				Replicas: readDsn,
				Policy:   dbresolver.RandomPolicy{}}
			readWritePlugin := dbresolver.Register(dbResolverCfg)
			err = db.Use(readWritePlugin)
			if err != nil {
				return nil
			}
		}
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
