// 自动生成模板SysDictionaryDetail
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 如果含有time.Time 请自行import time包
type SysDictionaryDetail struct {
	global.GVA_MODEL
	Label           string                `json:"label" form:"label" gorm:"column:label;comment:展示值"`                                  // 展示值
	Value           string                `json:"value" form:"value" gorm:"column:value;comment:字典值"`                                  // 字典值
	Extend          string                `json:"extend" form:"extend" gorm:"column:extend;comment:扩展值"`                               // 扩展值
	Status          *bool                 `json:"status" form:"status" gorm:"column:status;comment:启用状态"`                              // 启用状态
	Sort            int                   `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`                                    // 排序标记
	SysDictionaryID int                   `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:关联标记"` // 关联标记
	ParentID        *uint                 `json:"parentID" form:"parentID" gorm:"column:parent_id;comment:父级字典详情ID"`                   // 父级字典详情ID
	Children        []SysDictionaryDetail `json:"children" gorm:"foreignKey:ParentID"`                                                 // 子字典详情
	Level           int                   `json:"level" form:"level" gorm:"column:level;comment:层级深度"`                                 // 层级深度，从0开始
	Path            string                `json:"path" form:"path" gorm:"column:path;comment:层级路径"`                                    // 层级路径，如 "1,2,3"
	Disabled        bool                  `json:"disabled" gorm:"-"`                                                                   // 禁用状态，根据status字段动态计算
}

func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
