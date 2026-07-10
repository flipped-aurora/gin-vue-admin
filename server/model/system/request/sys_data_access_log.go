package request

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// SysDataAccessLogSearch 数据权限审计日志分页查询
type SysDataAccessLogSearch struct {
	common.PageInfo
	EventType   string `json:"eventType" form:"eventType"`     // 事件类型 no_identity/blocked_write
	TargetTable string `json:"targetTable" form:"targetTable"` // 受控业务表名(模糊)
}
