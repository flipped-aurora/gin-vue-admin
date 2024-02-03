package middlerware

import (
	"app/global"
	"app/response"
	"app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// jwt中间件，用于检测当前登录用户是否合法
func JWTCheck() gin.HandlerFunc {
	// 我们这里jwt鉴权取头部信息 jwt-code 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	return func(c *gin.Context) {
		y := c.Request.Header.Get("Authorization")
		if y == "" {
			response.CannotPassWithMessage("用户未登录", c)
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(y, &utils.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.JWTKey), nil
		})
		if token.Valid {
			c.Set("userID", token.Claims.(*utils.MyCustomClaims).UserId)
			c.Next()
		} else {
			fmt.Println(err)
			response.CannotPassWithMessage("用户登录过期或非法登录", c)
			c.Abort()
			return
		}
	}
}
