package request

import "github.com/eyotang/game-proxy/server/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
