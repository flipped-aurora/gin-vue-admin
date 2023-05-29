import service from '@/utils/request'

// @Tags Computation
// @Summary 创建Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Computation true "创建Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /computation/createComputation [post]
export const createComputation = (data) => {
  return service({
    url: '/computation/createComputation',
    method: 'post',
    data
  })
}

// @Tags Computation
// @Summary 删除Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Computation true "删除Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /computation/deleteComputation [delete]
export const deleteComputation = (data) => {
  return service({
    url: '/computation/deleteComputation',
    method: 'delete',
    data
  })
}

// @Tags Computation
// @Summary 删除Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /computation/deleteComputation [delete]
export const deleteComputationByIds = (data) => {
  return service({
    url: '/computation/deleteComputationByIds',
    method: 'delete',
    data
  })
}

// @Tags Computation
// @Summary 更新Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Computation true "更新Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /computation/updateComputation [put]
export const updateComputation = (data) => {
  return service({
    url: '/computation/updateComputation',
    method: 'put',
    data
  })
}

// @Tags Computation
// @Summary 用id查询Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Computation true "用id查询Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /computation/findComputation [get]
export const findComputation = (params) => {
  return service({
    url: '/computation/findComputation',
    method: 'get',
    params
  })
}

// @Tags Computation
// @Summary 分页获取Computation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Computation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /computation/getComputationList [get]
export const getComputationList = (params) => {
  return service({
    url: '/computation/getComputationList',
    method: 'get',
    params
  })
}
