package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// SysDataAccessLog 数据权限审计日志(sys_ 前缀 → 数据权限回调自动跳过, 无递归风险)
// 记录两类事件: no_identity(无身份访问受控表, 旁路盲区待补) / blocked_write(过滤后写 0 行, 疑似越权尝试)
type SysDataAccessLog struct {
	global.GVA_MODEL
	EventType   string `json:"eventType" gorm:"index;comment:事件类型 no_identity/blocked_write"` // 事件类型
	TargetTable string `json:"targetTable" gorm:"index;comment:受控业务表名"`                       // 受控业务表名
	Operation   string `json:"operation" gorm:"comment:操作 query/create/update/delete"`        // 操作类型
	UserID      uint   `json:"userId" gorm:"comment:事发用户ID(无身份事件为0)"`                         // 用户ID
	AuthorityID uint   `json:"authorityId" gorm:"comment:事发角色ID"`                             // 角色ID
	Scope       int    `json:"scope" gorm:"comment:事发时数据权限档位"`                                // 数据权限档位
	RequestID   string `json:"requestId" gorm:"comment:请求ID"`                                 // 请求链路ID
	Method      string `json:"method" gorm:"comment:HTTP方法"`                                  // HTTP方法
	Path        string `json:"path" gorm:"comment:请求路径"`                                      // 请求路径
	Detail      string `json:"detail" gorm:"comment:详情"`                                      // 详情
}

func (SysDataAccessLog) TableName() string {
	return "sys_data_access_logs"
}
