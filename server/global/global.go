package global

import (
	"gin-vue-admin/init"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_LOG    init.Logger
	GVA_CONFIG init.Config
	GVA_VP     *viper.Viper
)
