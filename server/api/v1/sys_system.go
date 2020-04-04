package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags system
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	err, config := new(model.System).GetSystemConfig()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"config": config}, "获取成功", c)
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
	var sys model.System
	_ = c.ShouldBindJSON(&sys)
	err := sys.SetSystemConfig()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("设置失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "设置成功", c)
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
	var sys model.System
	_ = c.ShouldBindJSON(&sys)
	err := sys.SetSystemConfig()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("设置失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "设置成功", c)
	}
}
