import service from '@/utils/request'

// @Tags SysRemotes
// @Summary 创建远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRemotes true "创建远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysRemotes/createSysRemotes [post]
export const createSysRemotes = (data) => {
  return service({
    url: '/sysRemotes/createSysRemotes',
    method: 'post',
    data
  })
}

// @Tags SysRemotes
// @Summary 删除远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRemotes true "删除远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysRemotes/deleteSysRemotes [delete]
export const deleteSysRemotes = (data) => {
  return service({
    url: '/sysRemotes/deleteSysRemotes',
    method: 'delete',
    data
  })
}

// @Tags SysRemotes
// @Summary 批量删除远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysRemotes/deleteSysRemotes [delete]
export const deleteSysRemotesByIds = (data) => {
  return service({
    url: '/sysRemotes/deleteSysRemotesByIds',
    method: 'delete',
    data
  })
}

// @Tags SysRemotes
// @Summary 更新远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRemotes true "更新远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysRemotes/updateSysRemotes [put]
export const updateSysRemotes = (data) => {
  return service({
    url: '/sysRemotes/updateSysRemotes',
    method: 'put',
    data
  })
}

// @Tags SysRemotes
// @Summary 用id查询远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysRemotes true "用id查询远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysRemotes/findSysRemotes [get]
export const findSysRemotes = (params) => {
  return service({
    url: '/sysRemotes/findSysRemotes',
    method: 'get',
    params
  })
}

// @Tags SysRemotes
// @Summary 分页获取远程服务器配置表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取远程服务器配置表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysRemotes/getSysRemotesList [get]
export const getSysRemotesList = (params) => {
  return service({
    url: '/sysRemotes/getSysRemotesList',
    method: 'get',
    params
  })
}
