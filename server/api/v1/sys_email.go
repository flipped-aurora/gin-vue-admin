package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags system
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /email/emailTest [post]
func EmailTest(c *gin.Context) {
	err := service.EmailTest()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("发送失败，%v", err), c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
