package request

import "github.com/eyotang/game-api-admin/server/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
