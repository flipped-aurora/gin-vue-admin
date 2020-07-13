package request

import "gin-vue-admin/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
