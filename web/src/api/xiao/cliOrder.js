import service from '@/utils/request'
// @Tags CliOrder
// @Summary 创建订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliOrder true "创建订单详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliOrder/createCliOrder [post]
export const createCliOrder = (data) => {
  return service({
    url: '/cliOrder/createCliOrder',
    method: 'post',
    data
  })
}

// @Tags CliOrder
// @Summary 删除订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliOrder true "删除订单详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliOrder/deleteCliOrder [delete]
export const deleteCliOrder = (params) => {
  return service({
    url: '/cliOrder/deleteCliOrder',
    method: 'delete',
    params
  })
}

// @Tags CliOrder
// @Summary 批量删除订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliOrder/deleteCliOrder [delete]
export const deleteCliOrderByIds = (params) => {
  return service({
    url: '/cliOrder/deleteCliOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags CliOrder
// @Summary 更新订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliOrder true "更新订单详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliOrder/updateCliOrder [put]
export const updateCliOrder = (data) => {
  return service({
    url: '/cliOrder/updateCliOrder',
    method: 'put',
    data
  })
}

// @Tags CliOrder
// @Summary 用id查询订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliOrder true "用id查询订单详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliOrder/findCliOrder [get]
export const findCliOrder = (params) => {
  return service({
    url: '/cliOrder/findCliOrder',
    method: 'get',
    params
  })
}

// @Tags CliOrder
// @Summary 分页获取订单详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliOrder/getCliOrderList [get]
export const getCliOrderList = (params) => {
  return service({
    url: '/cliOrder/getCliOrderList',
    method: 'get',
    params
  })
}

// @Tags CliOrder
// @Summary 不需要鉴权的订单详情接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliOrderSearch true "分页获取订单详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliOrder/getCliOrderPublic [get]
export const getCliOrderPublic = () => {
  return service({
    url: '/cliOrder/getCliOrderPublic',
    method: 'get',
  })
}
