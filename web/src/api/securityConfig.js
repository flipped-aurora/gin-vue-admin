import service from '@/utils/request'

// @Tags SecurityConfig
// @Summary 获取安全配置
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=system.SysSecurityConfig,msg=string} "获取安全配置"
// @Router /securityConfig/getSecurityConfig [get]
export const getSecurityConfig = () => {
  return service({
    url: '/securityConfig/getSecurityConfig',
    method: 'get'
  })
}

// @Tags SecurityConfig
// @Summary 设置安全配置
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysSecurityConfig true "安全配置"
// @Success 200 {object} response.Response{data=system.SysSecurityConfig,msg=string} "设置安全配置"
// @Router /securityConfig/setSecurityConfig [post]
export const setSecurityConfig = (data) => {
  return service({
    url: '/securityConfig/setSecurityConfig',
    method: 'post',
    data
  })
}
