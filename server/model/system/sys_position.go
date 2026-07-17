package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// SysPosition 岗位(平表,不进 Casbin,仅作组织身份/职务,与角色正交)
type SysPosition struct {
	global.GVA_MODEL
	Name   string `json:"name" gorm:"index;comment:岗位名称"`          // 岗位名称
	Code   string `json:"code" gorm:"comment:岗位编码"`                // 岗位编码
	Sort   int    `json:"sort" gorm:"default:0;comment:排序"`        // 排序
	Status *bool  `json:"status" gorm:"default:true;comment:是否启用"` // 是否启用
	Remark string `json:"remark" gorm:"comment:备注"`                // 备注
}

func (SysPosition) TableName() string {
	return "sys_positions"
}
