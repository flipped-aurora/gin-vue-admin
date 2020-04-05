package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
)

//工作流属性表
type SysWorkflow struct {
	gorm.Model
	WorkflowNickName    string                `json:"workflowNickName"`    // 工作流名称
	WorkflowName        string                `json:"workflowName"`        // 工作流英文id
	WorkflowDescription string                `json:"workflowDescription"` // 工作流描述
	WorkflowStepInfo    []SysWorkflowStepInfo `json:"workflowStep"`        // 工作流步骤
}

// 工作流状态表
type SysWorkflowStepInfo struct {
	gorm.Model
	SysWorkflowID   uint    `json:"workflowID"`      // 所属工作流ID
	IsStrat         bool    `json:"isStrat"`         // 是否是开始流节点
	StepName        string  `json:"stepName"`        // 工作流名称
	StepNo          float64 `json:"stepNo"`          // 步骤id （第几步）
	StepAuthorityID string  `json:"stepAuthorityID"` // 操作者级别id
	IsEnd           bool    `json:"isEnd"`           // 是否是完结流节点
}

// @title    Create
// @description   create a workflow, 创建工作流
// @auth                     （2020/04/05  20:22 ）
// @return                     error
func (wk *SysWorkflow) Create() error {
	err := global.GVA_DB.Create(&wk).Error
	return err
}
