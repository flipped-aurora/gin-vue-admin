package request

import "gin-vue-admin/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
