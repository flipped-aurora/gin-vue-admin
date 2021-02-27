package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: InitDB
//@description: 创建数据库并初始化
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

func InitDB(conf model.InitDB) error {
	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}
	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	fmt.Println(dsn)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := utils.CreateTable(dsn, "mysql", createSql); err != nil {
		return err
	}
	setting := map[string]interface{}{
		"mysql.path":     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		"mysql.db-name":  conf.DBName,
		"mysql.username": conf.UserName,
		"mysql.password": conf.Password,
	}
	if err := utils.WriteConfig(global.GVA_VP, setting); err != nil {
		return err
	}
	utils.InitDB()
	return nil
}
