package api

import (
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
// @Param data body modelInterface.PageInfo true "生成验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	id := captcha.NewLen(6)
	captcha.Server(captcha.StdWidth,captcha.StdHeight)
	servers.ReportFormat(c,true,"test",gin.H{
		"id":id,
		"picPath":"/base/captcha/"+id+".png",
	})
}

// @Tags base
// @Summary 生成验证码图片路径
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "生成验证码图片路径"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha/:id [get]
func CaptchaImg(c *gin.Context) {
	servers.GinCapthcaServeHTTP(c.Writer, c.Request)
}







