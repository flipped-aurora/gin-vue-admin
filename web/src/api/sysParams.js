import service from '@/utils/request'
// @Tags SysParams
// @Summary 创建参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysParams true "创建参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysParams/createSysParams [post]
export const createSysParams = (data) => {
  return service({
    url: '/sysParams/createSysParams',
    method: 'post',
    data
  })
}

// @Tags SysParams
// @Summary 删除参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysParams true "删除参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysParams/deleteSysParams [delete]
export const deleteSysParams = (params) => {
  return service({
    url: '/sysParams/deleteSysParams',
    method: 'delete',
    params
  })
}

// @Tags SysParams
// @Summary 批量删除参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysParams/deleteSysParams [delete]
export const deleteSysParamsByIds = (params) => {
  return service({
    url: '/sysParams/deleteSysParamsByIds',
    method: 'delete',
    params
  })
}

// @Tags SysParams
// @Summary 更新参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysParams true "更新参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysParams/updateSysParams [put]
export const updateSysParams = (data) => {
  return service({
    url: '/sysParams/updateSysParams',
    method: 'put',
    data
  })
}

// @Tags SysParams
// @Summary 用id查询参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysParams true "用id查询参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysParams/findSysParams [get]
export const findSysParams = (params) => {
  return service({
    url: '/sysParams/findSysParams',
    method: 'get',
    params
  })
}

// @Tags SysParams
// @Summary 分页获取参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取参数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysParams/getSysParamsList [get]
export const getSysParamsList = (params) => {
  return service({
    url: '/sysParams/getSysParamsList',
    method: 'get',
    params
  })
}

// @Tags SysParams
// @Summary 不需要鉴权的参数接口
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysParamsSearch true "分页获取参数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysParams/getSysParam [get]
export const getSysParam = (params) => {
  return service({
    url: '/sysParams/getSysParam',
    method: 'get',
    params
  })
}
