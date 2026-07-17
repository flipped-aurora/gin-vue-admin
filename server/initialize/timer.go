// server/initialize/timer.go
package initialize

import (
	"context"
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mediaService "github.com/flipped-aurora/gin-vue-admin/server/service/media"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
)

// Timer 注册可供定时任务面板选用的命名方法。
// 不再硬编码调度: 调度由 sys_timed_tasks 表驱动(种子见 source/system/timed_task.go),
// 启动加载见 LoadTimedTasks。ctx 由统一 Runner 注入(已带 datascope.WithSystem 与超时)。
// 二开注册自己的任务: 在此追加 task.Register, 然后在面板新建任务选择该方法即可。
func Timer() {
	task.Register("ClearDB", "清理数据库过期日志(操作记录/JWT黑名单/定时任务执行日志)", func(ctx context.Context, _ json.RawMessage) error {
		return task.ClearTable(global.GVA_DB.WithContext(ctx))
	})
	task.Register("CleanStaleUploads", "清理过期大文件上传会话", func(ctx context.Context, _ json.RawMessage) error {
		svc := mediaService.MediaUploadService{}
		return svc.CleanupStale(ctx, global.GVA_CONFIG.Media.SessionTTL)
	})
}
