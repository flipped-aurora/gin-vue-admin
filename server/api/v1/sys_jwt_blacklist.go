package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	ModelJwt := model.JwtBlacklist{
		Jwt: token,
	}
	err := ModelJwt.JsonInBlacklist()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("jwt作废失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "jwt作废成功", c)
	}
}
