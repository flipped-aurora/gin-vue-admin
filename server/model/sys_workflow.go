package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
	"time"
)

var WorkflowBusinessStruct map[string]func() GVA_Workflow
var WorkflowBusinessTable map[string]func() interface{}

type GVA_Workflow interface {
	CreateWorkflowMove() *WorkflowMove
	GetBusinessType() string
	GetBusinessID() uint
	GetWorkflowBase() WorkflowBase
}

type WorkflowBase struct {
	WorkflowMoveID    uint   `json:"workflowMoveID" gorm:"-"`
	BusinessID        uint   `json:"businessID" gorm:"<-:false;column:id"` // 业务对应ID（businessID）的返回
	BusinessType      string `json:"businessType" gorm:"-"`
	PromoterID        uint   `json:"promoterID" gorm:"-"`
	OperatorID        uint   `json:"operatorID" gorm:"-"`
	WorkflowProcessID string `json:"workflowProcessID" gorm:"-"`
	WorkflowNodeID    string `json:"workflowNodeID" gorm:"-"`
	Param             string `json:"param" gorm:"-"`
	Action            string `json:"action" gorm:"-"`
}

func (w WorkflowBase) CreateWorkflowMove() (businessModel *WorkflowMove) {
	return &WorkflowMove{
		GVA_MODEL:         global.GVA_MODEL{ID: w.WorkflowMoveID},
		BusinessType:      w.BusinessType,
		PromoterID:        w.PromoterID,
		OperatorID:        w.OperatorID,
		Param:             w.Param,
		WorkflowProcessID: w.WorkflowProcessID,
		WorkflowNodeID:    w.WorkflowNodeID,
		BusinessID:        w.BusinessID,
		Action:            w.Action,
		IsActive:          true,
	}
}

func (w WorkflowBase) GetBusinessType() (businessType string) {
	return w.BusinessType
}

func (w WorkflowBase) GetBusinessID() (businessID uint) {
	return w.BusinessID
}

func (w WorkflowBase) GetWorkflowBase() (workflowBase WorkflowBase) {
	return w
}

type WorkflowMove struct {
	global.GVA_MODEL
	WorkflowProcessID string          `json:"workflowProcessID" gorm:"comment:工作流模板ID"`
	WorkflowProcess   WorkflowProcess `gorm:"<-:false" json:"workflowProcess" gorm:"comment:工作流模板具体信息"`
	WorkflowNodeID    string          `json:"workflowNodeID" gorm:"comment:工作流节点ID"`
	WorkflowNode      WorkflowNode    `gorm:"<-:false" json:"workflowNode" gorm:"comment:工作流节点具体信息"`
	BusinessType      string          `json:"businessType" gorm:"comment:业务标记"`
	BusinessID        uint            `json:"businessID" gorm:"comment:业务ID"`
	PromoterID        uint            `json:"promoterID" gorm:"comment:当前流转发起人"`
	Promoter          SysUser         `gorm:"<-:false" json:"promoter" gorm:"comment:当前流转发起人信息"`
	OperatorID        uint            `json:"operatorID" gorm:"comment:当前流转操作人"`
	Operator          SysUser         `gorm:"<-:false" json:"operator" gorm:"comment:当前流转操作人信息"`
	Action            string          `json:"action" gorm:"comment:工作流驱动事件"`
	Param             string          `json:"param" gorm:"comment:工作流驱动参数"`
	IsActive          bool            `json:"isActive" gorm:"comment:是否是活跃节点 "`
}

type WorkflowProcess struct {
	ID          string `json:"id" form:"id" gorm:"comment:流程标识;primaryKey;unique;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name" gorm:"comment:流程名称"`
	Category    string         `json:"category" gorm:"comment:分类"`
	Clazz       string         `json:"clazz" gorm:"comment:类型"`
	Label       string         `json:"label" gorm:"comment:流程标题"`
	HideIcon    bool           `json:"hideIcon" gorm:"comment:是否隐藏图标"`
	Description string         `json:"description" gorm:"comment:详细介绍"`
	View        string         `json:"view" gorm:"comment:前端视图文件"`
	Nodes       []WorkflowNode `json:"nodes"` // 流程节点数据
	Edges       []WorkflowEdge `json:"edges"` // 流程链接数据
}

type WorkflowNode struct {
	ID                string `json:"id" form:"id" gorm:"comment:节点id;primaryKey;unique;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
	WorkflowProcessID string         `json:"workflowProcessID" gorm:"comment:流程标识"`
	Clazz             string         `json:"clazz" gorm:"comment:节点类型"`
	Label             string         `json:"label" gorm:"comment:节点名称"`
	Type              string         `json:"type" gorm:"comment:图标类型"`
	Shape             string         `json:"shape" gorm:"comment:形状"`
	Description       string         `json:"description" gorm:"comment:详细介绍"`
	View              string         `json:"view" gorm:"comment:前端视图文件"`
	X                 float64        `json:"y" gorm:"comment:x位置"`
	Y                 float64        `json:"x" gorm:"comment:y位置"`
	WaitState         string         `json:"waitState" gorm:"comment:等待属性"`
	StateValue        string         `json:"stateValue" gorm:"comment:等待值"`
	To                string         `json:"to" gorm:"comment:收件人"`
	Subject           string         `json:"subject" gorm:"comment:标题"`
	Content           string         `json:"content" gorm:"comment:内容"`
	Cycle             string         `json:"cycle" gorm:"comment:循环时间"`
	Duration          string         `json:"duration" gorm:"comment:持续时间"`
	HideIcon          bool           `json:"hideIcon" gorm:"comment:是否隐藏图标"`
	DueDate           *time.Time     `json:"dueDate" gorm:"comment:到期时间"`
	AssignType        string         `json:"assignType" gorm:"comment:审批类型"`
	AssignValue       string         `json:"assignValue" gorm:"comment:审批类型值"`
	Success           bool           `json:"success" gorm:"comment:是否成功"`
}

type WorkflowEdge struct {
	ID                  string `json:"id" form:"id" gorm:"comment:唯一标识;primaryKey;unique;not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt     `json:"-" gorm:"index"`
	WorkflowProcessID   string             `json:"-" gorm:"comment:流程标识"`
	Clazz               string             `json:"clazz" gorm:"comment:类型（线）"`
	Source              string             `json:"source" gorm:"comment:起点节点"`
	Target              string             `json:"target" gorm:"comment:目标节点"`
	SourceAnchor        int                `json:"sourceAnchor" gorm:"comment:起点"`
	TargetAnchor        int                `json:"targetAnchor" gorm:"comment:目标点"`
	Description         string             `json:"description" gorm:"comment:详细介绍"`
	Shape               string             `json:"shape" gorm:"comment:形状"`
	StartPoint          WorkflowStartPoint `json:"startPoint"` // 起点信息
	EndPoint            WorkflowEndPoint   `json:"endPoint"`   // 终点信息
	Label               string             `json:"label" gorm:"comment:标题"`
	HideIcon            bool               `json:"hideIcon" gorm:"comment:隐藏图标"`
	ConditionExpression string             `json:"conditionExpression" gorm:"comment:条件标识"`
	Seq                 string             `json:"seq" gorm:"comment:序号"`
	Reverse             bool               `json:"reverse" gorm:"comment:是否反向"`
}

type WorkflowStartPoint struct {
	WorkflowEdgeID string
	global.GVA_MODEL
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Index int     `json:"index"`
}

type WorkflowEndPoint struct {
	WorkflowEdgeID string
	global.GVA_MODEL
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Index int     `json:"index"`
}
