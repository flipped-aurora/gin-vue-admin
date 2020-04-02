package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/autoCodeModel"
	"github.com/gin-gonic/gin"
	"os"
)

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autoCodeModel.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(c *gin.Context) {
	var a autoCodeModel.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	err := a.CreateTemp()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		os.Remove("./ginvueadmin.zip")
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		os.Remove("./ginvueadmin.zip")
	}
}
