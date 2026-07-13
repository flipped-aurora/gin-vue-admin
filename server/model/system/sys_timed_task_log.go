package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const (
	TimedTaskTriggerAuto   = "auto"   // 调度触发
	TimedTaskTriggerManual = "manual" // 面板手动触发

	TimedTaskStatusSuccess = "success"
	TimedTaskStatusFail    = "fail"
	TimedTaskStatusTimeout = "timeout"
)

// SysTimedTaskLog 定时任务执行日志
type SysTimedTaskLog struct {
	global.GVA_MODEL
	TaskId      uint      `json:"taskId" form:"taskId" gorm:"index;column:task_id;comment:任务ID"`
	TaskName    string    `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名快照(任务删除后日志仍可读)"`
	TriggerType string    `json:"triggerType" form:"triggerType" gorm:"column:trigger_type;comment:触发方式 auto/manual"`
	StartedAt   time.Time `json:"startedAt" gorm:"column:started_at;comment:开始时间"`
	FinishedAt  time.Time `json:"finishedAt" gorm:"column:finished_at;comment:结束时间"`
	DurationMs  int64     `json:"durationMs" gorm:"column:duration_ms;comment:耗时毫秒"`
	Status      string    `json:"status" form:"status" gorm:"column:status;comment:结果 success/fail/timeout"`
	ErrorMsg    string    `json:"errorMsg" gorm:"type:text;column:error_msg;comment:错误信息(截断)"`
	Output      string    `json:"output" gorm:"type:text;column:output;comment:输出摘要(截断)"`
}

func (SysTimedTaskLog) TableName() string {
	return "sys_timed_task_logs"
}
