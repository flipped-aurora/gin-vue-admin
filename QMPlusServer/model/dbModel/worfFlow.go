package dbModel

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Application struct {
	gorm.Model
	WorkFlowID	string  // 所属工作流
	WorkFlowStepInfoID string // 当前节点ID
	ApplicationName string  // 申请人姓名
	ApplicationCause	string // 请假原因
	ApplicationStartData time.Time // 请假开始日期
	ApplicationEndData time.Time // 请假开始日期

}

type ApplicationWorkFlowProcess struct {
	gorm.Model
	ApplicationId uint  // 当前工作流所属申请的ID
	CurrentNode string   // 当前进度节点
	HistoricalNode string  //上一个进度节点
	CurrentUser string   // 当前进度操作人
	HistoricalUser string // 上一个进度的操作人
	State bool  // 状态 是否是正在进行的状态
}

type Workflow struct {
	gorm.Model
	WorkflowName string // 工作流名称
	WorkflowDescription string //工作流描述
	WorkflowStep []WorkflowStepInfo //工作流步骤
}

type WorkflowStepInfo struct {
	gorm.Model
	WorkflowID uint
	IsStrat	bool
	StepName string
	StepNo	float64
	StepAuthorityId string
	IsEnd bool
}