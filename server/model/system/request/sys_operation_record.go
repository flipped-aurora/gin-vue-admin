package request

import (
	"github.com/flipped-aurora/gin-vue-admin/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
