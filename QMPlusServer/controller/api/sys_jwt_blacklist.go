package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context){
	token := c.Request.Header.Get("x-token")
	ModelJwt := sysModel.JwtBlacklist{
		Jwt:token,
	}
	err := ModelJwt.JsonInBlacklist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("jwt作废失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "jwt作废成功", gin.H{})
	}
}
