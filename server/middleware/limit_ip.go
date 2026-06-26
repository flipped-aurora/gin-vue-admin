package middleware

import (
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type LimitConfig struct {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey func(c *gin.Context) string
	// 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err.Error()})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey 默认生成key
func DefaultGenerationKey(c *gin.Context) string {
	return "GVA_Limit" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	// 无缓存句柄（极端启动期）时 fail-open，避免误伤
	if global.GVA_CACHE == nil {
		return nil
	}
	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.GVA_LOG.Error("limit", zap.Error(err))
	}
	return err
}

func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.GVA_CONFIG.System.LimitTimeIP,
		Limit:         global.GVA_CONFIG.System.LimitCountIP,
	}.LimitWithTime()
}

// SetLimitWithTime 设置访问次数：窗口内计数到达 limit 即拒绝。
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.GVA_CACHE.IncrementWithExpire(key, 1, expiration)
	if err != nil {
		// 运行时缓存异常：记录日志并 fail-open 放行
		global.GVA_LOG.Error("limit increment", zap.Error(err))
		return nil
	}
	if count > int64(limit) {
		return errors.New("请求太过频繁，请稍后再试")
	}
	return nil
}

// CacheCheckOrMark 基于 GVA_CACHE 的限流计数 超限返回错误 cache 异常 fail-open
func CacheCheckOrMark(key string, expire int, limit int) error {
	if global.GVA_CACHE == nil {
		return nil
	}
	n, err := global.GVA_CACHE.IncrementWithExpire(key, 1, time.Duration(expire)*time.Second)
	if err != nil {
		global.GVA_LOG.Error("limit", zap.Error(err))
		return nil // fail-open
	}
	if int(n) > limit {
		return errors.New("请求太过频繁，请稍后再试")
	}
	return nil
}

// SecurityLimit 按安全配置对登录/敏感接口限流 未开启则放行
func SecurityLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		enable, window, count := service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService.CurrentLimit()
		if !enable {
			c.Next()
			return
		}
		key := "GVA_SecLimit" + c.ClientIP() + c.FullPath()
		if err := CacheCheckOrMark(key, window, count); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
