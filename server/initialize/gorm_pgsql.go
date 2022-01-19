package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// GormPgSql 初始化 Postgresql 数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormPgSql() *gorm.DB {
	p := global.GVA_CONFIG.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		if global.GVA_CONFIG.OpenReadWriteSeparation {
			var writeDsn []gorm.Dialector
			var readDsn []gorm.Dialector
			switch p.Duty {
			case "read":
				readDsn = append(readDsn, postgres.Open(p.Dsn()))
			case "write":
				writeDsn = append(writeDsn, postgres.Open(p.Dsn()))
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
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func GormPgSqlByConfig(p config.DB) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		if global.GVA_CONFIG.OpenReadWriteSeparation {
			var writeDsn []gorm.Dialector
			var readDsn []gorm.Dialector
			switch p.Duty {
			case "read":
				readDsn = append(readDsn, postgres.Open(p.Dsn()))
			case "write":
				writeDsn = append(writeDsn, postgres.Open(p.Dsn()))
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
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
