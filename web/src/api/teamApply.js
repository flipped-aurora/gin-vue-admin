import service from '@/utils/request'

// @Tags TeamApply
// @Summary 创建TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamApply true "创建TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamApply/createTeamApply [post]
export const createTeamApply = (data) => {
  return service({
    url: '/teamApply/createTeamApply',
    method: 'post',
    data
  })
}

// @Tags TeamApply
// @Summary 删除TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamApply true "删除TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamApply/deleteTeamApply [delete]
export const deleteTeamApply = (data) => {
  return service({
    url: '/teamApply/deleteTeamApply',
    method: 'delete',
    data
  })
}

// @Tags TeamApply
// @Summary 删除TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamApply/deleteTeamApply [delete]
export const deleteTeamApplyByIds = (data) => {
  return service({
    url: '/teamApply/deleteTeamApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags TeamApply
// @Summary 更新TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamApply true "更新TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamApply/updateTeamApply [put]
export const updateTeamApply = (data) => {
  return service({
    url: '/teamApply/updateTeamApply',
    method: 'put',
    data
  })
}

// @Tags TeamApply
// @Summary 用id查询TeamApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeamApply true "用id查询TeamApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamApply/findTeamApply [get]
export const findTeamApply = (params) => {
  return service({
    url: '/teamApply/findTeamApply',
    method: 'get',
    params
  })
}

// @Tags TeamApply
// @Summary 分页获取TeamApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeamApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamApply/getTeamApplyList [get]
export const getTeamApplyList = (params) => {
  return service({
    url: '/teamApply/getTeamApplyList',
    method: 'get',
    params
  })
}
