import service from '@/utils/request'

// @Tags ProjectHub
// @Summary 创建ProjectHub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectHub true "创建ProjectHub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Project/createProjectHub [post]
export const createProjectHub = (data) => {
  return service({
    url: '/Project/createProjectHub',
    method: 'post',
    data
  })
}

// @Tags ProjectHub
// @Summary 删除ProjectHub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectHub true "删除ProjectHub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Project/deleteProjectHub [delete]
export const deleteProjectHub = (data) => {
  return service({
    url: '/Project/deleteProjectHub',
    method: 'delete',
    data
  })
}

// @Tags ProjectHub
// @Summary 删除ProjectHub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectHub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Project/deleteProjectHub [delete]
export const deleteProjectHubByIds = (data) => {
  return service({
    url: '/Project/deleteProjectHubByIds',
    method: 'delete',
    data
  })
}

// @Tags ProjectHub
// @Summary 更新ProjectHub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectHub true "更新ProjectHub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Project/updateProjectHub [put]
export const updateProjectHub = (data) => {
  return service({
    url: '/Project/updateProjectHub',
    method: 'put',
    data
  })
}

// @Tags ProjectHub
// @Summary 用id查询ProjectHub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectHub true "用id查询ProjectHub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Project/findProjectHub [get]
export const findProjectHub = (params) => {
  return service({
    url: '/Project/findProjectHub',
    method: 'get',
    params
  })
}

// @Tags ProjectHub
// @Summary 分页获取ProjectHub列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProjectHub列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Project/getProjectHubList [get]
export const getProjectHubList = (params) => {
  return service({
    url: '/Project/getProjectHubList',
    method: 'get',
    params
  })
}
