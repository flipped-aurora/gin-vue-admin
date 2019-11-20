package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
)

//工作流属性表
type Workflow struct {
	gorm.Model
	WorkflowNickName    string             // 工作流名称
	WorkflowName        string             // 工作流英文id
	WorkflowDescription string             // 工作流描述
	WorkflowStep        []WorkflowStepInfo // 工作流步骤
}

// 工作流状态表
type WorkflowStepInfo struct {
	gorm.Model
	WorkflowID      uint    `json:"-"` // 所属工作流ID
	IsStrat         bool    // 是否是开始流节点
	StepName        string  // 工作流名称
	StepNo          float64 // 步骤id （第几步）
	StepAuthorityID string  // 操作者级别id
	IsEnd           bool    // 是否是完结流节点
}

//创建工作流
func (wk *Workflow) Create() error {
	err := qmsql.DEFAULTDB.Create(&wk).Error
	return err
}
