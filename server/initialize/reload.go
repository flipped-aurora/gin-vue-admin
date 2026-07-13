package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

// Reload 优雅地重新加载系统配置
func Reload() error {
	logger.Bg().Mod("system").Info("正在重新加载系统配置...")

	// 重新加载配置文件
	if err := global.GVA_VP.ReadInConfig(); err != nil {
		logger.Bg().Mod("system").Err(err).Error("重新读取配置文件失败!")
		return err
	}

	// 重新初始化数据库连接
	if global.GVA_DB != nil {
		db, _ := global.GVA_DB.DB()
		err := db.Close()
		if err != nil {
			logger.Bg().Mod("system").Err(err).Error("关闭原数据库连接失败!")
			return err
		}
	}

	// 重新建立数据库连接
	global.GVA_DB = Gorm()

	// 重新初始化其他配置
	OtherInit()
	DBList()

	if global.GVA_DB != nil {
		// 重新注册数据权限 GORM 回调
		RegisterDataScopeCallbacks()
		// 确保数据库表结构是最新的
		RegisterTables()
	}

	// 重新初始化定时任务
	Timer()
	LoadTimedTasks()

	logger.Bg().Mod("system").Info("系统配置重新加载完成")
	return nil
}
