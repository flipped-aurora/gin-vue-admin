package middleware

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"time"
)

// ip限制
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
	if err != nil || count == 0 {
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
