package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AliSmsApi struct{}

// @Tags System
// @Summary 发送（阿里）短信
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /aliSms/sendAliSms[post]
func (s *AliSmsApi) SendAliSms(c *gin.Context) {
	var ali model.AliModel
	_ = c.ShouldBindJSON(&ali)
	if err := service.ServiceGroupApp.SendAliSms(ali.Phones, ali.TemplateCode, ali.TemplateParam); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
