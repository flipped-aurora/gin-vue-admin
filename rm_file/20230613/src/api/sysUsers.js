import service from '@/utils/request'

// @Tags SysUsers
// @Summary 创建SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUsers true "创建SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUsers/createSysUsers [post]
export const createSysUsers = (data) => {
  return service({
    url: '/sysUsers/createSysUsers',
    method: 'post',
    data
  })
}

// @Tags SysUsers
// @Summary 删除SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUsers true "删除SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUsers/deleteSysUsers [delete]
export const deleteSysUsers = (data) => {
  return service({
    url: '/sysUsers/deleteSysUsers',
    method: 'delete',
    data
  })
}

// @Tags SysUsers
// @Summary 删除SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUsers/deleteSysUsers [delete]
export const deleteSysUsersByIds = (data) => {
  return service({
    url: '/sysUsers/deleteSysUsersByIds',
    method: 'delete',
    data
  })
}

// @Tags SysUsers
// @Summary 更新SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUsers true "更新SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUsers/updateSysUsers [put]
export const updateSysUsers = (data) => {
  return service({
    url: '/sysUsers/updateSysUsers',
    method: 'put',
    data
  })
}

// @Tags SysUsers
// @Summary 用id查询SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysUsers true "用id查询SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUsers/findSysUsers [get]
export const findSysUsers = (params) => {
  return service({
    url: '/sysUsers/findSysUsers',
    method: 'get',
    params
  })
}

// @Tags SysUsers
// @Summary 分页获取SysUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUsers/getSysUsersList [get]
export const getSysUsersList = (params) => {
  return service({
    url: '/sysUsers/getSysUsersList',
    method: 'get',
    params
  })
}
