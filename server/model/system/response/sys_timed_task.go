package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// SysTimedTaskRow 列表行: 任务 + 运行态(下次执行时间从调度器快照实时取, 不落库)
type SysTimedTaskRow struct {
	system.SysTimedTask
	NextRunAt *time.Time `json:"nextRunAt"` // 已调度(enabled)任务的下次执行时间; 未调度为 null
}
