package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"os"
)

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(c *gin.Context) {
	var a model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	err := service.CreateTemp(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		os.Remove("./ginvueadmin.zip")
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		os.Remove("./ginvueadmin.zip")
	}
}
