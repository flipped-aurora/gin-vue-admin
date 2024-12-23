package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"log"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	gormConfig := internal.Gorm.Config(p.Prefix, p.Singular)
	if global.GVA_CONFIG.Pgsql.SlowThreshold > 0 {
		// 定义慢查询阈值
		slowThreshold := time.Duration(global.GVA_CONFIG.Pgsql.SlowThreshold) * time.Millisecond

		// 配置自定义日志器
		newLogger := logger.New(
			log.New(&SqlLogWriter{}, "", log.LstdFlags), // 标准日志输出
			logger.Config{
				SlowThreshold: slowThreshold, // 慢 SQL 阈值
				LogLevel:      logger.Warn,   // 日志级别：Warn 或 Info
				Colorful:      false,         // 是否启用彩色输出
			},
		)
		gormConfig.Logger = newLogger
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), gormConfig); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func GormPgSqlByConfig(p config.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

type SqlLogWriter struct{}

func (w *SqlLogWriter) Write(p []byte) (n int, err error) {
	if strings.Contains(string(p), "SLOW SQL") {
		global.GVA_LOG.Error("SLOW-SQL", zap.String("sql", string(p)))
	} else {
		global.GVA_LOG.Info("NORMAL-SQL", zap.String("sql", string(p)))
	}
	return len(p), nil
}
