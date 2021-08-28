package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	notify_response "github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Api struct {
}

// @Tags Notify
// @Summary 发送文字消息接口
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body notify_response.TextNotify true "发送文字消息的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /notify/sendTextMessage [post]
func (s *Api) SendTextMessage(c *gin.Context) {
	var textNotify notify_response.TextNotify
	_ = c.ShouldBindJSON(&textNotify)
	if err := service.ServiceGroupApp.SendTextMessage(textNotify.Content, textNotify.AtMobiles, textNotify.IsAtAll); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

// @Tags Notify
// @Summary 发送图文链接消息接口
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body notify_response.LinkNotify true "发送图文链接消息的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /notify/sendLinkMessage [post]
func (s *Api) SendLinkMessage(c *gin.Context) {
	var linkNotify notify_response.LinkNotify
	_ = c.ShouldBindJSON(&linkNotify)
	if err := service.ServiceGroupApp.SendLinkMessage(linkNotify.Content, linkNotify.Title, linkNotify.PicUrl, linkNotify.MessageUrl); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

// @Tags Notify
// @Summary 发送markdown消息接口
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body notify_response.MarkdownNotify true "发送markdown消息的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /notify/sendMarkdownMessage [post]
func (s *Api) SendMarkdownMessage(c *gin.Context) {
	var markdownNotify notify_response.MarkdownNotify
	_ = c.ShouldBindJSON(&markdownNotify)
	if err := service.ServiceGroupApp.SendMarkdownMessage(markdownNotify.Content, markdownNotify.Title, markdownNotify.AtMobiles, markdownNotify.IsAtAll); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
