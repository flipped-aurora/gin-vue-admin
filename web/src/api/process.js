import service from '@/utils/request'

// @Tags Process
// @Summary 创建Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Process true "创建Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /process/createProcess [post]
export const createProcess = (data) => {
  return service({
    url: '/process/createProcess',
    method: 'post',
    data
  })
}

// @Tags Process
// @Summary 删除Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Process true "删除Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /process/deleteProcess [delete]
export const deleteProcess = (data) => {
  return service({
    url: '/process/deleteProcess',
    method: 'delete',
    data
  })
}

// @Tags Process
// @Summary 删除Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /process/deleteProcess [delete]
export const deleteProcessByIds = (data) => {
  return service({
    url: '/process/deleteProcessByIds',
    method: 'delete',
    data
  })
}

// @Tags Process
// @Summary 更新Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Process true "更新Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /process/updateProcess [put]
export const updateProcess = (data) => {
  return service({
    url: '/process/updateProcess',
    method: 'put',
    data
  })
}

// @Tags Process
// @Summary 用id查询Process
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Process true "用id查询Process"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /process/findProcess [get]
export const findProcess = (params) => {
  return service({
    url: '/process/findProcess',
    method: 'get',
    params
  })
}

// @Tags Process
// @Summary 分页获取Process列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Process列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /process/getProcessList [get]
export const getProcessList = (params) => {
  return service({
    url: '/process/getProcessList',
    method: 'get',
    params
  })
}
