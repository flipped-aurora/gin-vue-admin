package source

import (
	"database/sql"
	"github.com/spf13/viper"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: CreateTable
//@description: 创建数据库(mysql)
//@param: dsn string, driver string, createSql
//@return: error

func CreateTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: WriteConfig
//@description: 回写配置
//@param:
//@return: error

func WriteConfig(viper *viper.Viper, conf map[string]interface{}) error {
	for k, v := range conf {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@author: [Songzhibin97](https://github.com/Songzhibin97)
//@function: InitDB
//@description: 初始化db
//@param:
//@return: error

func InitDB() {
	Mysql.CheckDatabase()
	Mysql.CheckUtf8mb4()
	Mysql.Info()
	Mysql.Init()
	Mysql.AutoMigrateTables()
	Mysql.InitData()
}
