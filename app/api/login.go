package api

import (
	"app/global"
	"app/response"
	"app/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 登录逻辑
func Login(c *gin.Context) {
	code := c.Query("code")
	var appid = global.APPID
	var secret = global.SECRET

	res, err := utils.Request(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&secret=%s", appid, code, secret), "GET", nil)
	if err != nil {
		response.FailWithDetailed(err.Error(), "获取服务器code失败", c)
		return
	}
	var resData map[string]interface{}
	err = json.Unmarshal([]byte(res), &resData) // 进行JSON到Map的转换
	if err != nil {
		response.FailWithDetailed(err.Error(), "json反序列化失败", c)
		return
	}
	if resData["openid"] != "" {
		user, _ := UserDataService.GetUserInfoByWxOpenid(resData["openid"].(string))
		if user.ID == 0 {
			user.UserWxopenid = resData["openid"].(string)
			err := UserDataService.CreateUserInfo(&user)
			if err != nil {
				response.FailWithDetailed(err.Error(), "创建用户失败", c)
				return
			}
			user.SessionKey = resData["session_key"].(string)
			result := make(map[string]interface{})
			jwt, err := utils.GiveJWT(user.ID)
			result["jwt"] = jwt
			result["user"] = user
			response.OkWithData(result, c)
			return
		} else {
			user.UserWxopenid = resData["openid"].(string)
			user.SessionKey = resData["session_key"].(string)
			response.OkWithData(user, c)
		}
	} else {
		fmt.Println(resData)
		response.FailWithDetailed(resData, "获取openid失败", c)
		return
	}

}

func GetJwt(c *gin.Context) {
	id := c.Query("ID")
	if id == "" {
		response.FailWithMessage("参数请求错误", c)
		return
	}
	idu, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithDetailed(err.Error(), "jwt生成失败", c)
	}
	token, err := utils.GiveJWT(uint(idu))
	if err != nil {
		response.FailWithDetailed(err.Error(), "JWT生成错误", c)
		return
	}
	response.OkWithDetailed(token, "JWT生成成功", c)
}
