import service from '@/utils/request'

// @Tags ShopOrders
// @Summary 创建shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopOrders true "创建shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopOrders/createShopOrders [post]
export const createShopOrders = (data) => {
  return service({
    url: '/shopOrders/createShopOrders',
    method: 'post',
    data
  })
}

// @Tags ShopOrders
// @Summary 删除shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopOrders true "删除shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopOrders/deleteShopOrders [delete]
export const deleteShopOrders = (data) => {
  return service({
    url: '/shopOrders/deleteShopOrders',
    method: 'delete',
    data
  })
}

// @Tags ShopOrders
// @Summary 批量删除shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopOrders/deleteShopOrders [delete]
export const deleteShopOrdersByIds = (data) => {
  return service({
    url: '/shopOrders/deleteShopOrdersByIds',
    method: 'delete',
    data
  })
}

// @Tags ShopOrders
// @Summary 更新shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopOrders true "更新shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopOrders/updateShopOrders [put]
export const updateShopOrders = (data) => {
  return service({
    url: '/shopOrders/updateShopOrders',
    method: 'put',
    data
  })
}

// @Tags ShopOrders
// @Summary 用id查询shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopOrders true "用id查询shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopOrders/findShopOrders [get]
export const findShopOrders = (params) => {
  return service({
    url: '/shopOrders/findShopOrders',
    method: 'get',
    params
  })
}

// @Tags OrderRefund
// @Summary 根据ID查订单退款
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderRefund true "根据ID查订单退款"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopOrders/RefundOrders [get]
export const RefundShopOrders = (params) => {
  return service({
    url: '/shopOrders/refundOrder',
    method: 'put',
    params
  })
}

// @Tags ShopOrders
// @Summary 分页获取shopOrders表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取shopOrders表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopOrders/getShopOrdersList [get]
export const getShopOrdersList = (params) => {
  return service({
    url: '/shopOrders/getShopOrdersList',
    method: 'get',
    params
  })
}
