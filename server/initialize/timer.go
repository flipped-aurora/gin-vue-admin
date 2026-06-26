package initialize

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mediaService "github.com/flipped-aurora/gin-vue-admin/server/service/media"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 清理过期大文件上传会话
		_, err = global.GVA_Timer.AddTaskByFunc("CleanStaleUploads", "@hourly", func() {
			svc := mediaService.MediaUploadService{}
			if err := svc.CleanupStale(global.GVA_CONFIG.Media.SessionTTL); err != nil {
				global.GVA_LOG.Error("CleanStaleUploads error", zap.Error(err))
			}
		}, "定时清理过期上传会话", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 其他定时任务定在这里 参考上方使用方法

		//_, err := global.GVA_Timer.AddTaskByFunc("定时任务标识", "corn表达式", func() {
		//	具体执行内容...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}
