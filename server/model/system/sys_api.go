package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"size:256;comment:api路径"`             // api路径
	Description string `json:"description" gorm:"size:256;comment:api中文描述"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"size:256;comment:api组"`          // api组
	Method      string `json:"method" gorm:"size:256;default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (SysApi) TableName() string {
	return "sys_apis"
}
