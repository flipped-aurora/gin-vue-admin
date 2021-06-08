import service from '@/utils/request'

// @Tags WorkflowProcess
// @Summary 创建WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "创建WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/createWorkflowProcess [post]
export const createWorkflowProcess = (data) => {
  return service({
    url: '/workflowProcess/createWorkflowProcess',
    method: 'post',
    data
  })
}

// @Tags WorkflowProcess
// @Summary 删除WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "删除WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workflowProcess/deleteWorkflowProcess [delete]
export const deleteWorkflowProcess = (data) => {
  return service({
    url: '/workflowProcess/deleteWorkflowProcess',
    method: 'delete',
    data
  })
}

// @Tags WorkflowProcess
// @Summary 删除WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workflowProcess/deleteWorkflowProcess [delete]
export const deleteWorkflowProcessByIds = (data) => {
  return service({
    url: '/workflowProcess/deleteWorkflowProcessByIds',
    method: 'delete',
    data
  })
}

// @Tags WorkflowProcess
// @Summary 更新WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "更新WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /workflowProcess/updateWorkflowProcess [put]
export const updateWorkflowProcess = (data) => {
  return service({
    url: '/workflowProcess/updateWorkflowProcess',
    method: 'put',
    data
  })
}

// @Tags WorkflowProcess
// @Summary 用id查询WorkflowProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "用id查询WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workflowProcess/findWorkflowProcess [get]
export const findWorkflowProcess = (params) => {
  return service({
    url: '/workflowProcess/findWorkflowProcess',
    method: 'get',
    params
  })
}

// @Tags WorkflowProcess
// @Summary 分页获取WorkflowProcess列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取WorkflowProcess列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getWorkflowProcessList [get]
export const getWorkflowProcessList = (params) => {
  return service({
    url: '/workflowProcess/getWorkflowProcessList',
    method: 'get',
    params
  })
}

// @Tags WorkflowProcess
// @Summary 用id查询工作流步骤
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WorkflowProcess true "用id查询WorkflowProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workflowProcess/findWorkflowStep [get]
export const findWorkflowStep = (params) => {
  return service({
    url: '/workflowProcess/findWorkflowStep',
    method: 'get',
    params
  })
}

// @Tags ExaWfLeave
// @Summary 创建ExaWfLeave
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/startWorkflow [post]
export const startWorkflow = (data, params = { businessType: data.wf.businessType }) => {
  return service({
    url: '/workflowProcess/startWorkflow',
    method: 'post',
    data,
    params
  })
}

// @Tags ExaWfLeave
// @Summary 创建ExaWfLeave
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/completeWorkflowMove [post]
export const completeWorkflowMove = (data, params = { businessType: data.wf.businessType }) => {
  return service({
    url: '/workflowProcess/completeWorkflowMove',
    method: 'post',
    data,
    params
  })
}

// @Tags WorkflowProcess
// @Summary 我发起的工作流
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getMyStated [get]
export const getMyStated = () => {
  return service({
    url: '/workflowProcess/getMyStated',
    method: 'get'
  })
}

// @Tags WorkflowProcess
// @Summary 我发起的工作流
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getMyNeed [get]
export const getMyNeed = () => {
  return service({
    url: '/workflowProcess/getMyNeed',
    method: 'get'
  })
}

// @Tags WorkflowProcess
// @Summary 根据id获取当前节点详情和历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取当前节点详情和过往"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workflowProcess/getWorkflowMoveByID [get]
export const getWorkflowMoveByID = (params) => {
  return service({
    url: '/workflowProcess/getWorkflowMoveByID',
    method: 'get',
    params
  })
}
