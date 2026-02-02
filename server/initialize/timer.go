package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/task"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
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

		// 其他定时任务定在这里 参考上方使用方法

		// 证书监控定时任务：每天凌晨2点更新所有域名证书状态
		_, err = global.GVA_Timer.AddTaskByFunc("UpdateAllCertificates", "0 0 2 * * *", func() {
			err := task.UpdateAllCertificates()
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时更新所有域名证书状态", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}
	}()
}
