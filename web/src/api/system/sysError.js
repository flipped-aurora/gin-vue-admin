import service from '@/utils/request'
// @Tags SysError
// @Summary 创建错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysError true "创建错误日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysErrpr/createSysError [post]
export const createSysError = (data) => {
  return service({
    url: '/sysErrpr/createSysError',
    method: 'post',
    data
  })
}

// @Tags SysError
// @Summary 删除错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysError true "删除错误日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysErrpr/deleteSysError [delete]
export const deleteSysError = (params) => {
  return service({
    url: '/sysErrpr/deleteSysError',
    method: 'delete',
    params
  })
}

// @Tags SysError
// @Summary 批量删除错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除错误日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysErrpr/deleteSysError [delete]
export const deleteSysErrorByIds = (params) => {
  return service({
    url: '/sysErrpr/deleteSysErrorByIds',
    method: 'delete',
    params
  })
}

// @Tags SysError
// @Summary 更新错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysError true "更新错误日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysErrpr/updateSysError [put]
export const updateSysError = (data) => {
  return service({
    url: '/sysErrpr/updateSysError',
    method: 'put',
    data
  })
}

// @Tags SysError
// @Summary 用id查询错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysError true "用id查询错误日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysErrpr/findSysError [get]
export const findSysError = (params) => {
  return service({
    url: '/sysErrpr/findSysError',
    method: 'get',
    params
  })
}

// @Tags SysError
// @Summary 分页获取错误日志列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取错误日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysErrpr/getSysErrorList [get]
export const getSysErrorList = (params) => {
  return service({
    url: '/sysErrpr/getSysErrorList',
    method: 'get',
    params
  })
}

// @Tags SysError
// @Summary 不需要鉴权的错误日志接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysErrorSearch true "分页获取错误日志列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysErrpr/getSysErrorPublic [get]
export const getSysErrorPublic = () => {
  return service({
    url: '/sysErrpr/getSysErrorPublic',
    method: 'get',
  })
}
