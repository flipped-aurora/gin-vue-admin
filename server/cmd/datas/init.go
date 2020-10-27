package datas

import (
	"gin-vue-admin/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

func InitMysqlData(db *gorm.DB) {
	var err error
	err = InitSysApi(db)
	err = InitSysUser(db)
	err = InitExaCustomer(db)
	err = InitCasbinModel(db)
	err = InitSysAuthority(db)
	err = InitSysBaseMenus(db)
	err = InitAuthorityMenu(db)
	err = InitSysDictionary(db)
	err = InitSysAuthorityMenus(db)
	err = InitSysDataAuthorityId(db)
	err = InitSysDictionaryDetail(db)
	err = InitExaFileUploadAndDownload(db)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql]-->初始化数据成功")
}

func InitMysqlTables(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(
		model.SysApi{},
		model.SysUser{},
		model.ExaFile{},
		model.ExaCustomer{},
		model.SysBaseMenu{},
		model.SysWorkflow{},
		model.SysAuthority{},
		model.JwtBlacklist{},
		model.ExaFileChunk{},
		model.SysDictionary{},
		model.ExaSimpleUploader{},
		model.SysOperationRecord{},
		model.SysWorkflowStepInfo{},
		model.SysDictionaryDetail{},
		model.SysBaseMenuParameter{},
		model.ExaFileUploadAndDownload{},
	)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql]-->初始化数据表成功")
}

func InitPostgresqlData(db *gorm.DB) {
	var err error
	err = InitSysApi(db)
	err = InitSysUser(db)
	err = InitExaCustomer(db)
	err = InitCasbinModel(db)
	err = InitSysAuthority(db)
	err = InitSysBaseMenus(db)
	err = InitAuthorityMenu(db)
	err = InitSysAuthorityMenus(db)
	err = InitSysDataAuthorityId(db)
	err = InitSysDictionaryDetail(db)
	err = InitExaFileUploadAndDownload(db)
	err = InitSysDictionaryToPostgresql(db)
	if err != nil {
		color.Error.Printf("[Postgresql]-->初始化数据失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Postgresql]-->初始化数据成功")
}

func InitPostgresqlTables(db *gorm.DB) {
	var err error
	if !db.Migrator().HasTable("casbin_rule") {
		err = db.Migrator().CreateTable(&gormadapter.CasbinRule{})
	}
	err = db.AutoMigrate(
		model.SysApi{},
		model.SysUser{},
		model.ExaFile{},
		model.ExaCustomer{},
		model.SysBaseMenu{},
		model.SysWorkflow{},
		model.SysAuthority{},
		model.JwtBlacklist{},
		model.ExaFileChunk{},
		model.ExaSimpleUploader{},
		model.SysOperationRecord{},
		model.SysWorkflowStepInfo{},
		model.SysDictionaryDetail{},
		model.SysBaseMenuParameter{},
		model.ExaFileUploadAndDownload{},
		SysDictionaryToPostgresql{},
	)
	if err != nil {
		color.Error.Printf("[Postgresql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Postgresql]-->初始化数据表成功")
}
