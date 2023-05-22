import service from '@/utils/request'

// @Tags Agent
// @Summary 创建Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agent true "创建Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agent/createAgent [post]
export const createAgent = (data) => {
  return service({
    url: '/agent/createAgent',
    method: 'post',
    data
  })
}

// @Tags Agent
// @Summary 删除Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agent true "删除Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agent/deleteAgent [delete]
export const deleteAgent = (data) => {
  return service({
    url: '/agent/deleteAgent',
    method: 'delete',
    data
  })
}

// @Tags Agent
// @Summary 删除Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agent/deleteAgent [delete]
export const deleteAgentByIds = (data) => {
  return service({
    url: '/agent/deleteAgentByIds',
    method: 'delete',
    data
  })
}

// @Tags Agent
// @Summary 更新Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agent true "更新Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /agent/updateAgent [put]
export const updateAgent = (data) => {
  return service({
    url: '/agent/updateAgent',
    method: 'put',
    data
  })
}

// @Tags Agent
// @Summary 用id查询Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Agent true "用id查询Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /agent/findAgent [get]
export const findAgent = (params) => {
  return service({
    url: '/agent/findAgent',
    method: 'get',
    params
  })
}

// @Tags Agent
// @Summary 分页获取Agent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Agent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agent/getAgentList [get]
export const getAgentList = (params) => {
  return service({
    url: '/agent/getAgentList',
    method: 'get',
    params
  })
}
