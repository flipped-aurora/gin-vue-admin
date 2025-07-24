import service from '@/utils/request'

// @Tags SysVersion
// @Summary 删除版本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysVersion true "删除版本管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysVersion/deleteSysVersion [delete]
export const deleteSysVersion = (params) => {
  return service({
    url: '/sysVersion/deleteSysVersion',
    method: 'delete',
    params
  })
}

// @Tags SysVersion
// @Summary 批量删除版本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除版本管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysVersion/deleteSysVersion [delete]
export const deleteSysVersionByIds = (params) => {
  return service({
    url: '/sysVersion/deleteSysVersionByIds',
    method: 'delete',
    params
  })
}

// @Tags SysVersion
// @Summary 用id查询版本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysVersion true "用id查询版本管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysVersion/findSysVersion [get]
export const findSysVersion = (params) => {
  return service({
    url: '/sysVersion/findSysVersion',
    method: 'get',
    params
  })
}

// @Tags SysVersion
// @Summary 分页获取版本管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取版本管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysVersion/getSysVersionList [get]
export const getSysVersionList = (params) => {
  return service({
    url: '/sysVersion/getSysVersionList',
    method: 'get',
    params
  })
}

// @Tags SysVersion
// @Summary 导出版本数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "导出版本数据"
// @Success 200 {string} string "{\"success\":true,\"data\":{},\"msg\":\"导出成功\"}"
// @Router /sysVersion/exportVersion [post]
export const exportVersion = (data) => {
  return service({
    url: '/sysVersion/exportVersion',
    method: 'post',
    data
  })
}

// @Tags SysVersion
// @Summary 下载版本JSON数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query string true "版本ID"
// @Success 200 {string} string "{\"success\":true,\"data\":{},\"msg\":\"下载成功\"}"
// @Router /sysVersion/downloadVersionJson [get]
export const downloadVersionJson = (params) => {
  return service({
    url: '/sysVersion/downloadVersionJson',
    method: 'get',
    params,
    responseType: 'blob'
  })
}

// @Tags SysVersion
// @Summary 导入版本数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "版本JSON数据"
// @Success 200 {string} string "{\"success\":true,\"data\":{},\"msg\":\"导入成功\"}"
// @Router /sysVersion/importVersion [post]
export const importVersion = (data) => {
  return service({
    url: '/sysVersion/importVersion',
    method: 'post',
    data
  })
}
