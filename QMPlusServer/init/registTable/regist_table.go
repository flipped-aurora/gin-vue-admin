package registTable

import (
	"gin-vue-admin/init/initlog"
	"gin-vue-admin/model/dbModel"
	"gin-vue-admin/model/sysModel"
	"github.com/jinzhu/gorm"
)

//注册数据库表专用
func RegistTable(db *gorm.DB, logger log.Logger) {
	db.AutoMigrate(sysModel.SysUser{},
		sysModel.SysAuthority{},
		sysModel.SysMenu{},
		sysModel.SysApi{},
		sysModel.SysBaseMenu{},
		sysModel.JwtBlacklist{},
		sysModel.SysWorkflow{},
		sysModel.SysWorkflowStepInfo{},
		dbModel.ExaFileUploadAndDownload{},
		dbModel.ExaFile{},
		dbModel.ExaFileChunk{},
		dbModel.ExaCustomer{},
	)
	logger.Debug("register table success")
}
