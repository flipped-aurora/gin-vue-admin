package request

import (
	"github.com/gzpz/golf-sales-system/server/model/common/request"
	"github.com/gzpz/golf-sales-system/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
