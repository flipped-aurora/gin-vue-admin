package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
