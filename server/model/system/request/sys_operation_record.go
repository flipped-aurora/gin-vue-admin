package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
