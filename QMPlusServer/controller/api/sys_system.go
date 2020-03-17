package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

// @Tags system
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	err, config := new(sysModel.System).GetSystemConfig()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"config": config})
	}
}

// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sysModel.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/setSystemConfig [post]
func SetSystemConfig(c *gin.Context) {
	var sys sysModel.System
	_ = c.ShouldBindJSON(&sys)
	err := sys.SetSystemConfig()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("设置失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "设置成功", gin.H{})
	}
}


//本方法开发中 开发者windows系统 缺少linux系统所需的包 因此搁置
// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sysModel.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/ReloadSystem [post]
func ReloadSystem(c *gin.Context) {
	var sys sysModel.System
	_ = c.ShouldBindJSON(&sys)
	err := sys.SetSystemConfig()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("设置失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "设置成功", gin.H{})
	}
}
