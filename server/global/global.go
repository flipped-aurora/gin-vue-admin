package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/tools"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_REDIS               *tools.Redis
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	IdWorker                *tools.Worker // 雪花算法实例
	GVA_Timer               timer.Timer   = timer.NewTimerTask()
	GVA_Concurrency_Control               = &singleflight.Group{}

	BlackCache local_cache.Cache
)
