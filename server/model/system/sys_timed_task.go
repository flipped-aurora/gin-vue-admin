package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

const (
	TimedTaskExecutorMethod = "method" // 调用已注册本体方法
	TimedTaskExecutorHTTP   = "http"   // 定时 HTTP 调用
)

// SysTimedTask 定时任务定义(任务唯一事实源, 调度状态以 enabled 为准)
// 唯一性: name 由 service 层校验(软删除下 DB 唯一索引会挡住同名重建, 故不建唯一索引)
type SysTimedTask struct {
	global.GVA_MODEL
	Name             string         `json:"name" form:"name" gorm:"index;column:name;comment:任务名"`
	Description      string         `json:"description" form:"description" gorm:"column:description;comment:任务说明(面板展示的提示)"`
	Spec             string         `json:"spec" form:"spec" gorm:"column:spec;comment:cron表达式(支持@daily等描述符)"`
	WithSeconds      bool           `json:"withSeconds" form:"withSeconds" gorm:"column:with_seconds;comment:表达式是否含秒位"`
	ExecutorType     string         `json:"executorType" form:"executorType" gorm:"column:executor_type;comment:执行器类型 method/http"`
	MethodName       string         `json:"methodName" form:"methodName" gorm:"column:method_name;comment:已注册方法名(executor=method)"`
	Params           datatypes.JSON `json:"params" form:"params" gorm:"column:params;comment:方法自由JSON入参" swaggertype:"object"`
	HttpUrl          string         `json:"httpUrl" form:"httpUrl" gorm:"column:http_url;comment:HTTP回调地址(executor=http)"`
	HttpMethod       string         `json:"httpMethod" form:"httpMethod" gorm:"column:http_method;comment:HTTP方法"`
	HttpHeader       datatypes.JSON `json:"httpHeader" form:"httpHeader" gorm:"column:http_header;comment:HTTP自定义请求头(JSON对象)" swaggertype:"object"`
	HttpBody         string         `json:"httpBody" form:"httpBody" gorm:"type:text;column:http_body;comment:HTTP请求体"`
	HttpAllowPrivate bool           `json:"httpAllowPrivate" form:"httpAllowPrivate" gorm:"column:http_allow_private;comment:允许访问内网/环回地址(默认禁止,SSRF防护)"`
	Enabled          bool           `json:"enabled" form:"enabled" gorm:"column:enabled;comment:是否启用"`
}

func (SysTimedTask) TableName() string {
	return "sys_timed_tasks"
}
