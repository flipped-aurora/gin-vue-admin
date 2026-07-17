// 自动生成模板SysOperationRecord
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysOperationRecord struct {
	global.GVA_MODEL
	Ip           string  `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                                    // 请求ip
	Method       string  `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                                        // 请求方法
	Path         string  `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                                              // 请求路径
	Status       int     `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                                        // 请求状态
	LatencyMs    int64   `json:"latency_ms" form:"latency_ms" gorm:"column:latency_ms;comment:延迟(毫秒)"`                          // 延迟(毫秒),便于 SQL 直接聚合
	Agent        string  `json:"agent" form:"agent" gorm:"type:text;column:agent;comment:代理"`                                   // 代理
	ErrorMessage string  `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`                   // 错误信息
	Body         string  `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                                  // 请求Body
	Resp         string  `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                                  // 响应Body
	UserID       int     `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`                                     // 用户id
	RequestID    string  `json:"request_id" form:"request_id" gorm:"index;column:request_id;type:varchar(64);comment:请求ID"`     // 请求ID
	TraceID      string  `json:"trace_id" form:"trace_id" gorm:"index;column:trace_id;type:varchar(64);comment:链路ID"`           // 链路ID,索引支撑按 trace 反查
	DeviceID     string  `json:"device_id" form:"device_id" gorm:"column:device_id;type:varchar(64);comment:设备ID"`              // 设备ID
	User         SysUser `json:"user" form:"-"` // 操作人(Preload 填充);form:"-" 阻止 gin 查询绑定递归进 SysUser
}
