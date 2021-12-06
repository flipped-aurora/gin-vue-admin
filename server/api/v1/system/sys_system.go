package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	if err, config := systemConfigService.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
}

// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	_ = c.ShouldBindJSON(&sys)
	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithData("设置成功", c)
	}
}

// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"重启系统成功"}"
// @Router /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GVA_LOG.Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
	} else {
		response.OkWithMessage("重启系统成功", c)
	}
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
