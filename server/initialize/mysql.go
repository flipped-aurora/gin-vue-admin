package initialize

import (
	"gin-vue-admin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.GVA_CONFIG.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       admin.Username + ":" + admin.Password + "@(" + admin.Path + ")/" + admin.Dbname + "?" + admin.Config, // DSN data source name
		DefaultStringSize:         256,                                                                                                  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                // 根据版本自动配置
	}
	var gormConfig *gorm.Config
	if admin.LogMode { //根据配置决定是否开启日志
		gormConfig = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		gormConfig = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.GVA_DB = db
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(admin.MaxIdleConns)
		sqlDB.SetMaxOpenConns(admin.MaxOpenConns)
	}
}
