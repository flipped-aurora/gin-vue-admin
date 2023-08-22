import service from '@/utils/request'

// @Tags OrderInfo
// @Summary 创建OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderInfo true "创建OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderInfo/createOrderInfo [post]
export const createOrderInfo = (data) => {
  return service({
    url: '/orderInfo/createOrderInfo',
    method: 'post',
    data
  })
}

// @Tags OrderInfo
// @Summary 删除OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderInfo true "删除OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderInfo/deleteOrderInfo [delete]
export const deleteOrderInfo = (data) => {
  return service({
    url: '/orderInfo/deleteOrderInfo',
    method: 'delete',
    data
  })
}

// @Tags OrderInfo
// @Summary 删除OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderInfo/deleteOrderInfo [delete]
export const deleteOrderInfoByIds = (data) => {
  return service({
    url: '/orderInfo/deleteOrderInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderInfo
// @Summary 更新OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderInfo true "更新OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderInfo/updateOrderInfo [put]
export const updateOrderInfo = (data) => {
  return service({
    url: '/orderInfo/updateOrderInfo',
    method: 'put',
    data
  })
}

// @Tags OrderInfo
// @Summary 用id查询OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderInfo true "用id查询OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderInfo/findOrderInfo [get]
export const findOrderInfo = (params) => {
  return service({
    url: '/orderInfo/findOrderInfo',
    method: 'get',
    params
  })
}

// @Tags OrderInfo
// @Summary 分页获取OrderInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderInfo/getOrderInfoList [get]
export const getOrderInfoList = (params) => {
  return service({
    url: '/orderInfo/getOrderInfoList',
    method: 'get',
    params
  })
}
