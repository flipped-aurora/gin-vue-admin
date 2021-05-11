package request

import "github.com/eyotang/gin-vue-admin/server/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
