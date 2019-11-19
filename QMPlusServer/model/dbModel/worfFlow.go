package dbModel

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Application struct {
	gorm.Model
	WorkFlowID           string    // 所属工作流ID
	WorkFlowStepInfoID   string    // 当前节点ID
	ApplicationName      string    // 申请人姓名
	ApplicationCause     string    // 请假原因
	ApplicationStartData time.Time // 请假开始日期
	ApplicationEndData   time.Time // 请假开始日期

}

// 流转表
type ApplicationWorkFlowProcess struct {
	gorm.Model
	ApplicationId  uint // 当前工作流所属申请的ID
	WorkflowID     uint
	CurrentNode    string // 当前进度节点
	HistoricalNode string //上一个进度节点
	CurrentUser    string // 当前进度操作人
	HistoricalUser string // 上一个进度的操作人
	State          bool   // 状态 是否是正在进行的状态
}

//工作流属性表
type Workflow struct {
	gorm.Model
	WorkflowNickName    string             // 工作流名称
	WorkflowName        string             // 工作流英文id
	WorkflowDescription string             //工作流描述
	WorkflowStep        []WorkflowStepInfo //工作流步骤
}

// 工作流状态表
type WorkflowStepInfo struct {
	gorm.Model
	WorkflowID      uint    // 所属工作流ID
	IsStrat         bool    // 是否是开始流节点
	StepName        string  // 工作流名称
	StepNo          float64 // 步骤id （第几步）
	StepAuthorityId string  // 操作者级别id
	IsEnd           bool    // 是否是完结流节点
}
