// 自动生成模板ExaWfLeave
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type ExaWfLeave struct {
	global.GVA_MODEL
	Cause     string    `json:"cause" form:"cause" gorm:"column:cause;comment:"`
	StartTime time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:"`
	EndTime   time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:"`
}

type ExaWfLeaveWorkflow struct {
	// 工作流操作结构体
	WorkflowBase `json:"wf"`
	ExaWfLeave   `json:"business"`
}

func (e ExaWfLeave) TableName() string {
	return "exa_wf_leaves"
}
