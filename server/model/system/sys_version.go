// 自动生成模板SysVersion
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 版本管理 结构体  SysVersion
type SysVersion struct {
	global.GVA_MODEL
	VersionName *string `json:"versionName" form:"versionName" gorm:"comment:版本名称;column:version_name;size:255;" binding:"required"` //版本名称
	VersionCode *string `json:"versionCode" form:"versionCode" gorm:"comment:版本号;column:version_code;size:100;" binding:"required"`   //版本号
	Description *string `json:"description" form:"description" gorm:"comment:版本描述;column:description;size:500;"`                     //版本描述
	VersionData *string `json:"versionData" form:"versionData" gorm:"comment:版本数据JSON;column:version_data;type:text;"`               //版本数据
	Status      *int    `json:"status" form:"status" gorm:"default:1;comment:状态 1:草稿 2:已发布;column:status;" binding:"required"`    //状态
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 版本管理 SysVersion自定义表名 sys_versions
func (SysVersion) TableName() string {
	return "sys_versions"
}
