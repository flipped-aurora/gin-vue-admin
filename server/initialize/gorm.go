package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var err error

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		GormMysql()
	case "postgresql":
		GormPostgreSql()
	//case "sqlite": // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	//	GormSqlite()
	case "sqlserver":
		GormSqlServer()
	}
}

// GormDBTables 注册数据库表专用
func GormDBTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaSimpleUploader{},
		model.ExaCustomer{},
		model.SysOperationRecord{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}

// GormMysql 初始化Mysql数据库
func GormMysql() {
	m := global.GVA_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := config(m.LogMode)
	if global.GVA_DB, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
	} else {
		GormDBTables(global.GVA_DB)
		sqlDB, _ := global.GVA_DB.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	}
}

// GormPostgreSql 初始化PostgreSql数据库
func GormPostgreSql() {
	p := global.GVA_CONFIG.Postgresql
	dsn := "host="+ p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
	postgresConfig := postgres.Config{
		DSN:                  dsn,                    // DSN data source name
		PreferSimpleProtocol: p.PreferSimpleProtocol, // 禁用隐式 prepared statement
	}
	gormConfig := config(p.Logger)
	if global.GVA_DB, err = gorm.Open(postgres.New(postgresConfig), gormConfig); err != nil {
		global.GVA_LOG.Error("PostgreSql启动异常", zap.Any("err", err))
		os.Exit(0)
	} else {
		GormDBTables(global.GVA_DB)
		sqlDB, _ := global.GVA_DB.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	}
}

// GormSqlite 初始化Sqlite数据库 sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
//func GormSqlite() {
//	s := global.GVA_CONFIG.Sqlite
//	gormConfig := config(s.Logger)
//	if global.GVA_DB, err = gorm.Open(sqlite.Open(s.Path), gormConfig); err != nil {
//		global.GVA_LOG.Error("Sqlite启动异常", zap.Any("err", err))
//		os.Exit(0)
//	} else {
//      GormDBTables(global.GVA_DB)
//		sqlDB, _ := global.GVA_DB.DB()
//		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
//	}
//}

// GormSqlServer 初始化SqlServer数据库
func GormSqlServer() {
	ss := global.GVA_CONFIG.Sqlserver
	dsn := "sqlserver://" + ss.Username + ":" + ss.Password + "@" + ss.Path + "?database=gorm"
	if global.GVA_DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		global.GVA_LOG.Error("SqlServer启动异常", zap.Any("err", err))
		os.Exit(0)
	} else {
		GormDBTables(global.GVA_DB)
		sqlDB, _ := global.GVA_DB.DB()
		sqlDB.SetMaxIdleConns(ss.MaxIdleConns)
		sqlDB.SetMaxOpenConns(ss.MaxOpenConns)
	}
}

// config 根据配置决定是否开启日志
func config(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}
