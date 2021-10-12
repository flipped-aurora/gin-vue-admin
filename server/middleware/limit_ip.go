package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type LimitConfig struct {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey func(c *gin.Context) string
	// 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark func(key string, Expire int) bool
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l *LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if l.CheckOrMark(l.GenerationKey(c), l.Expire) {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": "操作频繁,请稍后再试"})
			c.Abort()
			return
		}
	}
}

// DefaultGenerationKey 默认生成key
func DefaultGenerationKey(c *gin.Context) string {
	return "GVA_Limit" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) bool {
	// 判断是否开启redis
	if global.GVA_REDIS == nil {
		return true
	}
	if err := SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.GVA_LOG.Error("limit", zap.Error(err))
		return false
	}
	return true

}

// IPLimit ip限制
func IPLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "RequestClientIPLimit===" + c.ClientIP()
		limit := global.GVA_CONFIG.System.LimitCountIP
		second := global.GVA_CONFIG.System.LimitTimeIP
		expiration := time.Duration(second) * time.Second
		if err := SetLimitWithTime(key, limit, expiration); err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		// 继续往下处理
		c.Next()
	}
}

// 设置访问次数
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.GVA_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := global.GVA_REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err := pipe.Exec(context.Background())
		return err
	} else {
		//次数
		if times, err := global.GVA_REDIS.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.GVA_REDIS.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("请求太过频繁，请稍后再试")
				} else {
					return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
				}
			} else {
				return global.GVA_REDIS.Incr(context.Background(), key).Err()
			}
		}
	}
}
