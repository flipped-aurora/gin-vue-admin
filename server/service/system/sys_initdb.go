package system

import (
	"database/sql"
	"fmt"
	"path/filepath"

	uuid "github.com/satori/go.uuid"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/source"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: writeConfig
//@description: 回写配置
//@param: viper *viper.Viper, mysql config.Mysql
//@return: error

type InitDBService struct {
}

func (initDBService *InitDBService) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.GVA_CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	viper.Set("jwt.signing-key", uuid.NewV4())
	return viper.WriteConfig()
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: createTable
//@description: 创建数据库(mysql)
//@param: dsn string, driver string, createSql
//@return: error

func (initDBService *InitDBService) createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func (initDBService *InitDBService) initDB(InitDBFunctions ...system.InitDBFunc) (err error) {
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
//@param: conf request.InitDB
//@return: error

func (initDBService *InitDBService) InitDB(conf request.InitDB) error {

	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}

	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Dbname:   conf.DBName,
		Username: conf.UserName,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if MysqlConfig.Dbname == "" {
		return nil
	}

	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.GVA_DB = db
	}

	err := global.GVA_DB.AutoMigrate(
		system.SysUser{},
		system.SysAuthority{},
		system.SysApi{},
		system.SysBaseMenu{},
		system.SysBaseMenuParameter{},
		system.JwtBlacklist{},
		system.SysDictionary{},
		system.SysDictionaryDetail{},
		example.ExaFileUploadAndDownload{},
		example.ExaFile{},
		example.ExaFileChunk{},
		example.ExaCustomer{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
	)
	if err != nil {
		global.GVA_DB = nil
		return err
	}
	err = initDBService.initDB(
		source.Admin,
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
		source.UserAuthority,
	)
	if err != nil {
		global.GVA_DB = nil
		return err
	}

	if err = initDBService.writeConfig(global.GVA_VP, MysqlConfig); err != nil {
		return err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}
