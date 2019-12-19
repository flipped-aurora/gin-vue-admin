package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

// @Tags workflow
// @Summary 注册工作流
// @Produce  application/json
// @Param data body sysModel.SysWorkflow true "注册工作流接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /workflow/createWorkFlow [post]
func CreateWorkFlow(c *gin.Context) {
	var wk sysModel.SysWorkflow
	_ = c.ShouldBind(&wk)
	err := wk.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{})
	}
}
