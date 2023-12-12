package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 用户被拉黑后 可在此处拦截 启用redis可加速判断过程减少资源消耗
		if global.GVA_CONFIG.JWT.UserEnableVerify {
			var userEnable int
			userInfo := &system.SysUser{}
			getUserEnable := func(uuid string) int {
				userInfo, err = userService.FindUserByUuid(uuid)
				if err != nil {
					return 0
				}
				return userInfo.Enable
			}
			if global.GVA_REDIS != nil {
				userEnable, err = global.GVA_REDIS.Get(context.Background(), strconv.Itoa(int(claims.BaseClaims.ID))+".user.enable").Int()
				if err != nil {
					// 若没有缓存则从数据库中获取
					userEnable = getUserEnable(claims.UUID.String())
					// 缓存到redis
					if dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime); err == nil {
						global.GVA_REDIS.Set(context.Background(), strconv.Itoa(int(claims.BaseClaims.ID))+".user.enable", userEnable, dr)
					}
				}
			} else {
				userEnable = getUserEnable(claims.UUID.String())
			}
			if userEnable != 1 {
				response.FailWithDetailed(gin.H{"reload": true}, "您的帐户已被删除或禁用", c)
				c.Abort()
				return
			}
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			if global.GVA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
