// server/initialize/timed_task.go
package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/datascope"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

// LoadTimedTasks 从 sys_timed_tasks 恢复调度(仅 enabled; 幂等, 可在启动/重载/initdb 后重复调用)。
// 必须在 RegisterTables(建表)之后调用。
func LoadTimedTasks() {
	if global.GVA_DB == nil {
		return
	}
	ctx := datascope.WithSystem(context.Background())
	if err := system.TimedTaskServiceApp.LoadAll(ctx); err != nil {
		logger.Bg().Mod("timedTask").Err(err).Error("定时任务启动加载失败")
	}
}
