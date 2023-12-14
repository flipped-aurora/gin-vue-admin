package request

import (
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
