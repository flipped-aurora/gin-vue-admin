package qmsql

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-vue-admin/config"
)

var DEFAULTDB *gorm.DB

//初始化数据库并产生数据库全局变量
func InitMysql(admin config.MysqlAdmin) *gorm.DB {
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		// 数据库初始化失败时，直接退出程序
		log.Fatalf("DEFAULTDB数据库启动异常: %s", err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		DEFAULTDB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		DEFAULTDB.LogMode(admin.LogMode)
	}
	return DEFAULTDB
}
