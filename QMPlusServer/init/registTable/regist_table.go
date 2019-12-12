package registTable

import (
	"github.com/jinzhu/gorm"
	"main/model/dbModel"
	"main/model/sysModel"
)

//注册数据库表专用
func RegistTable(db *gorm.DB) {
	db.AutoMigrate(sysModel.SysUser{},
		sysModel.SysAuthority{},
		sysModel.SysMenu{},
		sysModel.SysApi{},
		sysModel.SysApiAuthority{},
		sysModel.SysBaseMenu{},
		dbModel.ExaFileUploadAndDownload{},
		sysModel.SysWorkflow{},
		sysModel.SysWorkflowStepInfo{},
	)
}
