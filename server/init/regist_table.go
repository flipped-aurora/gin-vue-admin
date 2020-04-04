package init

import (
	"gin-vue-admin/model"
	"github.com/jinzhu/gorm"
)

//注册数据库表专用
func RegisterTable(db *gorm.DB) {
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysMenu{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
	)
	L.Debug("register table success")
}
