package service

import (
	"database/sql"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/source"
	"github.com/spf13/viper"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: writeConfig
//@description: 回写配置
//@param:
//@return: error

func writeConfig(viper *viper.Viper, conf map[string]interface{}) error {
	for k, v := range conf {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: createTable
//@description: 创建数据库(mysql)
//@param: dsn string, driver string, createSql
//@return: error

func createTable(dsn string, driver string, createSql string) error {
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

func initDB(InitDBFunctions ...model.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: InitDB
//@description: 创建数据库并初始化
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

func InitDB(conf request.InitDB) error {
	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}
	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	fmt.Println(dsn)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := createTable(dsn, "mysql", createSql); err != nil {
		return err
	}
	setting := map[string]interface{}{
		"mysql.path":     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		"mysql.db-name":  conf.DBName,
		"mysql.username": conf.UserName,
		"mysql.password": conf.Password,
	}
	if err := writeConfig(global.GVA_VP, setting); err != nil {
		return err
	}
	err := initDB(source.Admin,
		source.Api,
		source.AuthorityMenu,
		source.Authority,
		source.AuthoritiesMenus,
		source.Casbin,
		source.DataAuthorities,
		source.Dictionary,
		source.DictionaryDetail,
		source.File,
		source.BaseMenu,
		source.Workflow)
	if err!=nil {
		return err
	}
	return nil
}
