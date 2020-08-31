package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
}

type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string
	DataAuthorityIdAuthorityId string
}

func InitSysApi() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.SysApiData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysUser() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.SysUserData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitExaCustomer() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.ExaCustomerData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitCasbinModel() (err error) {
	if !global.GVA_DB.Migrator().HasTable("casbin_rule") {
		if err := global.GVA_DB.Migrator().CreateTable(&gormadapter.CasbinRule{}); err != nil {
			return err
		}
	}
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.CasbinModelData()
	if tx.Table("casbin_rule").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysAuthority() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.SysAuthorityData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysBaseMenus() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.SysBaseMenusData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitAuthorityMenu() (err error) {
	return global.GVA_DB.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`, `sys_base_menus`.`updated_at` AS `updated_at`, `sys_base_menus`.`deleted_at` AS `deleted_at`, `sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`, `sys_base_menus`.`title`  AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on ((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)))").Error
}

func InitSysDictionary() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.SysDictionaryData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysAuthorityMenus() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []SysAuthorityMenus{
		{"888", 1},
		{"888", 2},
		{"888", 3},
		{"888", 4},
		{"888", 5},
		{"888", 6},
		{"888", 7},
		{"888", 8},
		{"888", 9},
		{"888", 10},
		{"888", 11},
		{"888", 12},
		{"888", 13},
		{"888", 14},
		{"888", 15},
		{"888", 16},
		{"888", 17},
		{"888", 18},
		{"888", 19},
		{"888", 20},
		{"888", 21},
		{"888", 22},
		{"888", 23},
		{"888", 24},
		{"888", 25},
		{"8881", 1},
		{"8881", 2},
		{"8881", 8},
		{"8881", 17},
		{"8881", 18},
		{"8881", 19},
		{"8881", 20},
		{"9528", 1},
		{"9528", 2},
		{"9528", 3},
		{"9528", 4},
		{"9528", 5},
		{"9528", 6},
		{"9528", 7},
		{"9528", 8},
		{"9528", 9},
		{"9528", 10},
		{"9528", 11},
		{"9528", 12},
		{"9528", 13},
		{"9528", 14},
		{"9528", 15},
		{"9528", 17},
		{"9528", 18},
		{"9528", 19},
		{"9528", 20},
	}
	if tx.Table("sys_authority_menus").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDataAuthorityId() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []SysDataAuthorityId{
		{"888", "888"},
		{"888", "8881"},
		{"888", "9528"},
		{"9528", "8881"},
		{"9528", "9528"},
	}
	if tx.Table("sys_data_authority_ids").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDictionaryDetail() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.SysDictionaryDetailData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitExaFileUploadAndDownload() (err error) {
	tx := global.GVA_DB.Begin() // 开始事务
	insert := model.ExaFileUploadAndDownloadData()
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}
