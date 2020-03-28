package qmsql

import (
	"gin-vue-admin/config"
	"gin-vue-admin/init/initlog"

	"github.com/cenkalti/backoff/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DEFAULTDB *gorm.DB

//初始化数据库并产生数据库全局变量
func InitMysql(admin config.MysqlAdmin) *gorm.DB {

	// 当DB无法连接时，进行适当的Retry处理
	var err error
	openDB := func() error {
		DEFAULTDB, err = gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config)
		if err != nil {
			return err
		}
		return nil
	}
	policy := backoff.WithMaxRetries(backoff.NewConstantBackOff(admin.Retry.Wait), uint64(admin.Retry.Count))

	err = backoff.Retry(openDB, policy)
	if err != nil {
		// Retry之后仍然无法连接，数据库初始化失败，直接退出程序
		log.L.Fatal("DEFAULTDB数据库启动异常:", err.Error())
	}

	DEFAULTDB.DB().SetMaxIdleConns(admin.MaxIdleConns)
	DEFAULTDB.DB().SetMaxOpenConns(admin.MaxOpenConns)
	DEFAULTDB.LogMode(admin.LogMode)
	return DEFAULTDB
}
