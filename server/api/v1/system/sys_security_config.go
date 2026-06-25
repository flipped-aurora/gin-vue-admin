package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SecurityConfigApi struct{}

// GetSecurityConfig
// @Tags      SecurityConfig
// @Summary   获取安全配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=system.SysSecurityConfig,msg=string}  "获取安全配置"
// @Router    /securityConfig/getSecurityConfig [get]
func (s *SecurityConfigApi) GetSecurityConfig(c *gin.Context) {
	cfg, err := securityConfigService.Get()
	if err != nil {
		global.GVA_LOG.Error("获取安全配置失败!", zap.Error(err))
		response.FailWithMessage("获取安全配置失败", c)
		return
	}
	response.OkWithDetailed(cfg, "获取成功", c)
}

// SetSecurityConfig
// @Tags      SecurityConfig
// @Summary   设置安全配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysSecurityConfig                                     true  "安全配置"
// @Success   200   {object}  response.Response{data=system.SysSecurityConfig,msg=string}  "设置安全配置"
// @Router    /securityConfig/setSecurityConfig [post]
func (s *SecurityConfigApi) SetSecurityConfig(c *gin.Context) {
	var cfg system.SysSecurityConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := securityConfigService.Set(cfg); err != nil {
		global.GVA_LOG.Error("设置安全配置失败!", zap.Error(err))
		response.FailWithMessage("设置安全配置失败", c)
		return
	}
	saved := securityConfigService.Current()
	response.OkWithDetailed(saved, "设置成功", c)
}
