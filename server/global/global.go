package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"kirer.cn/server/config"
	"kirer.cn/server/utils/timer"
)

var (
	DB                  *gorm.DB
	REDIS               *redis.Client
	CONFIG              config.Server
	VP                  *viper.Viper
	LOG                 *zap.Logger
	TIMER               timer.Timer = timer.NewTimerTask()
	Concurrency_Control             = &singleflight.Group{}
	BlackCache          local_cache.Cache
)
