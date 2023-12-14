package global

import (
	"sync"

	"github.com/qiniu/qmgo"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"kirer.cn/server/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"kirer.cn/server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	REDIS  *redis.Client
	MONGO  *qmgo.QmgoClient
	CONFIG config.Server
	VP     *viper.Viper
	// LOG    *oplogging.Logger
	LOG                 *zap.Logger
	TIMER               timer.Timer = timer.NewTimerTask()
	Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
