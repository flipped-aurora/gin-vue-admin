package request

import (
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
