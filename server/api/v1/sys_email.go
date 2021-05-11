package v1

import (
	"github.com/eyotang/game-api-admin/server/global"
	"github.com/eyotang/game-api-admin/server/model/response"
	"github.com/eyotang/game-api-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func EmailTest(c *gin.Context) {
	if err := service.EmailTest(); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
