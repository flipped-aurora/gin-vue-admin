package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram"
	"github.com/go-pay/gopay/alipay"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"golang.org/x/sync/singleflight"
	"sync"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                 *gorm.DB
	GVA_DBList             map[string]*gorm.DB
	GVA_REDIS              *redis.Client
	GVA_WECHAT_MINI        *miniprogram.MiniProgram
	GVA_DOUYIN_OPEN        *open.API
	GVA_WECHAT_PAY         *jsapi.JsapiApiService
	GVA_WECHAT_PAY_APP     *app.AppApiService
	GVA_WECHAT_PAY_H5      *h5.H5ApiService
	GVA_ALI_PAY            *alipay.Client
	GVA_WECHAT_REFUND      *refunddomestic.RefundsApiService
	GVA_WECHAT_PAY_HANDLER *notify.Handler

	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache    local_cache.Cache
	lock          sync.RWMutex
	AppBlackCache local_cache.Cache
	GVA_KEYLOCK   = NewKeyLock()
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetAppApi() string {
	return "api/"
}
