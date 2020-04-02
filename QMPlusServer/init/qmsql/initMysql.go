package qmsql

import (
	"gin-vue-admin/config"
	"gin-vue-admin/init/initlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DEFAULTDB *gorm.DB

//初始化数据库并产生数据库全局变量
func InitMysql(admin config.MysqlAdmin) *gorm.DB {
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		log.L.Error("DEFAULTDB数据库启动异常",err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		DEFAULTDB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		DEFAULTDB.LogMode(admin.LogMode)
	}
	return DEFAULTDB
}
