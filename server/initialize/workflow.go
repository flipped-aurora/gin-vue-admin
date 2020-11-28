package initialize

import "gin-vue-admin/model"

func initWorkflowModel() {
	model.WorkflowBusinessStruct = make(map[string]func() model.GVA_Workflow)
	model.WorkflowBusinessStruct["leave"] = func() model.GVA_Workflow {
		return new(model.ExaWfLeaveWorkflow)
	}
}

func initWorkflowTable() {
	model.WorkflowBusinessTable = make(map[string]func() interface{})
	model.WorkflowBusinessTable["leave"] = func() interface{} {
		return new(model.ExaWfLeave)
	}
}

func InitWkMode() {
	initWorkflowModel()
	initWorkflowTable()
}
