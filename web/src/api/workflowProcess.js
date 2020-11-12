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
        url: "/workflowProcess/createWorkflowProcess",
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
        url: "/workflowProcess/deleteWorkflowProcess",
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
        url: "/workflowProcess/deleteWorkflowProcessByIds",
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
        url: "/workflowProcess/updateWorkflowProcess",
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
        url: "/workflowProcess/findWorkflowProcess",
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
        url: "/workflowProcess/getWorkflowProcessList",
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
        url: "/workflowProcess/findWorkflowStep",
        method: 'get',
        params
    })
}