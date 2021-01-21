package information

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var Workflow = new(workflow)

type workflow struct{}

var WorkflowProcess = []model.WorkflowProcess{
	{ID: "leaveFlow", CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "leaveFlow", Clazz: "process", Label: "请假流程（演示）", HideIcon: false, Description: "请假流程演示", View: "view/iconList/index.vue"},
}

var WorkflowNodes = []model.WorkflowNode{
	{ID: "end1603681358043", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "end", Label: "请假失败", Type: "end-node", Shape: "end-node", Description: "", View: "view/exa_wf_leave/exa_wf_leaveFrom.vue", X: 302, Y: 545.5, HideIcon: false, AssignType: "", AssignValue: "", Success: false},
	{ID: "end1603681360882", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "end", Label: "请假成功", Type: "end-node", Shape: "end-node", Description: "请假完成，具体结果等待提交", View: "view/exa_wf_leave/exa_wf_leaveFrom.vue", X: 83.5, Y: 546, HideIcon: false, AssignType: "", AssignValue: "", Success: true},
	{ID: "start1603681292875", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "start", Label: "发起请假", Type: "start-node", Shape: "start-node", Description: "发起一个请假流程", View: "view/exa_wf_leave/exa_wf_leaveFrom.vue", X: 201, Y: 109, HideIcon: false, AssignType: "", AssignValue: "", Success: false},
	{ID: "userTask1603681299962", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "userTask", Label: "审批", Type: "user-task-node", Shape: "user-task-node", Description: "审批会签", View: "view/exa_wf_leave/exa_wf_leaveFrom.vue", X: 202, Y: 320.5, HideIcon: false, AssignType: "user", AssignValue: ",1,2,", Success: false},
}
var WorkflowEdge = []model.WorkflowEdge{
	{ID: "flow1604985849039", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "flow", Source: "start1603681292875", Target: "userTask1603681299962", SourceAnchor: 1, TargetAnchor: 3, Shape: "flow-polyline-round", Label: "", HideIcon: false, ConditionExpression: "", Reverse: false},
	{ID: "flow1604985879574", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "flow", Source: "userTask1603681299962", Target: "end1603681360882", SourceAnchor: 0, TargetAnchor: 2, Shape: "flow-polyline-round", Label: "同意", HideIcon: false, ConditionExpression: "yes", Reverse: false},
	{ID: "flow1604985881207", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "flow", Source: "userTask1603681299962", Target: "end1603681358043", SourceAnchor: 2, TargetAnchor: 2, Shape: "flow-polyline-round", Label: "不同意", HideIcon: false, ConditionExpression: "no", Reverse: false},
}
var WorkflowStartPoint = []model.WorkflowStartPoint{
	{WorkflowEdgeID: "flow1604985849039", GVA_MODEL: global.GVA_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, X: 137, Y: 201, Index: 1},
	{WorkflowEdgeID: "flow1604985879574", GVA_MODEL: global.GVA_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, X: 320.5, Y: 174, Index: 0},
	{WorkflowEdgeID: "flow1604985881207", GVA_MODEL: global.GVA_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, X: 320.5, Y: 230, Index: 2},
}

var WorkflowEndPoint = []model.WorkflowEndPoint{
	{WorkflowEdgeID: "flow1604985849039", GVA_MODEL: global.GVA_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, X: 270, Y: 202, Index: 3},
	{WorkflowEdgeID: "flow1604985879574", GVA_MODEL: global.GVA_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, X: 518, Y: 83.5, Index: 2},
	{WorkflowEdgeID: "flow1604985881207", GVA_MODEL: global.GVA_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, X: 517.5, Y: 302, Index: 2},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 工作流相关 表数据初始化
func (w *workflow) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&WorkflowProcess).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&WorkflowNodes).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&WorkflowEdge).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&WorkflowStartPoint).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&WorkflowEndPoint).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> 工作流相关 表初始数据成功!")
		return nil
	})
}
