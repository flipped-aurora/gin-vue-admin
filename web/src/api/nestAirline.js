import service from '@/utils/request'

// @Tags NestAirline
// @Summary 创建NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestAirline true "创建NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtAirline/createNestAirline [post]
export const createNestAirline = (data) => {
  return service({
    url: '/NtAirline/createNestAirline',
    method: 'post',
    data
  })
}

// @Tags NestAirline
// @Summary 删除NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestAirline true "删除NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NtAirline/deleteNestAirline [delete]
export const deleteNestAirline = (data) => {
  return service({
    url: '/NtAirline/deleteNestAirline',
    method: 'delete',
    data
  })
}

// @Tags NestAirline
// @Summary 删除NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NtAirline/deleteNestAirline [delete]
export const deleteNestAirlineByIds = (data) => {
  return service({
    url: '/NtAirline/deleteNestAirlineByIds',
    method: 'delete',
    data
  })
}

// @Tags NestAirline
// @Summary 更新NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestAirline true "更新NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NtAirline/updateNestAirline [put]
export const updateNestAirline = (data) => {
  return service({
    url: '/NtAirline/updateNestAirline',
    method: 'put',
    data
  })
}

// @Tags NestAirline
// @Summary 用id查询NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NestAirline true "用id查询NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NtAirline/findNestAirline [get]
export const findNestAirline = (params) => {
  return service({
    url: '/NtAirline/findNestAirline',
    method: 'get',
    params
  })
}

// @Tags NestAirline
// @Summary 分页获取NestAirline列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取NestAirline列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtAirline/getNestAirlineList [get]
export const getNestAirlineList = (params) => {
  return service({
    url: '/NtAirline/getNestAirlineList',
    method: 'get',
    params
  })
}
