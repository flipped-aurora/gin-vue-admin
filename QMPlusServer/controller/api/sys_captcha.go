package api

import (
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

// 获取图片验证码id
// @Tags base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	captchaId := captcha.NewLen(config.GinVueAdminconfig.Captcha.KeyLong)
	servers.ReportFormat(c, true, "验证码获取成功", gin.H{
		"captchaId": captchaId,
		"picPath":   "/base/captcha/" + captchaId + ".png",
	})
}

// @Tags base
// @Summary 生成验证码图片路径
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha/:captchaId [get]
func CaptchaImg(c *gin.Context) {
	servers.GinCapthcaServeHTTP(c.Writer, c.Request)
}
