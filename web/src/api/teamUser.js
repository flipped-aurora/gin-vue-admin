import service from '@/utils/request'

// @Tags TeamUser
// @Summary 创建TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamUser true "创建TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamUser/createTeamUser [post]
export const createTeamUser = (data) => {
  return service({
    url: '/teamUser/createTeamUser',
    method: 'post',
    data
  })
}

// @Tags TeamUser
// @Summary 删除TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamUser true "删除TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamUser/deleteTeamUser [delete]
export const deleteTeamUser = (data) => {
  return service({
    url: '/teamUser/deleteTeamUser',
    method: 'delete',
    data
  })
}

// @Tags TeamUser
// @Summary 删除TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamUser/deleteTeamUser [delete]
export const deleteTeamUserByIds = (data) => {
  return service({
    url: '/teamUser/deleteTeamUserByIds',
    method: 'delete',
    data
  })
}

// @Tags TeamUser
// @Summary 更新TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamUser true "更新TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamUser/updateTeamUser [put]
export const updateTeamUser = (data) => {
  return service({
    url: '/teamUser/updateTeamUser',
    method: 'put',
    data
  })
}

// @Tags TeamUser
// @Summary 用id查询TeamUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeamUser true "用id查询TeamUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamUser/findTeamUser [get]
export const findTeamUser = (params) => {
  return service({
    url: '/teamUser/findTeamUser',
    method: 'get',
    params
  })
}

// @Tags TeamUser
// @Summary 分页获取TeamUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeamUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamUser/getTeamUserList [get]
export const getTeamUserList = (params) => {
  return service({
    url: '/teamUser/getTeamUserList',
    method: 'get',
    params
  })
}
