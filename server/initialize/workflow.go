package initialize

import "gin-vue-admin/model"

func initWkModel() {
	model.WorkflowBusinessStruct = make(map[string]func() model.GVA_Workflow)
	model.WorkflowBusinessStruct["leave"] = func() model.GVA_Workflow {
		return new(model.ExaWfLeaveWorkflow)
	}
}

func initWkTable() {

	model.WorkflowBusinessTable = make(map[string]string)
	model.WorkflowBusinessTable["leave"] = "exa_wf_leaves"
}

func InitWkMode() {
	initWkModel()
	initWkTable()

}
