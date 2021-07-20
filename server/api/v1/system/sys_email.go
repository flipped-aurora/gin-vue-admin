package system

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (s *SystemApi) EmailTest(c *gin.Context) {
	if err := emailService.EmailTest(); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
