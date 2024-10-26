import service from '@/utils/request'
// @Tags CliMainorder
// @Summary 创建订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainorder true "创建订单总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /climainorder/createCliMainorder [post]
export const createCliMainorder = (data) => {
  return service({
    url: '/climainorder/createCliMainorder',
    method: 'post',
    data
  })
}

// @Tags CliMainorder
// @Summary 删除订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainorder true "删除订单总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /climainorder/deleteCliMainorder [delete]
export const deleteCliMainorder = (params) => {
  return service({
    url: '/climainorder/deleteCliMainorder',
    method: 'delete',
    params
  })
}

// @Tags CliMainorder
// @Summary 批量删除订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /climainorder/deleteCliMainorder [delete]
export const deleteCliMainorderByIds = (params) => {
  return service({
    url: '/climainorder/deleteCliMainorderByIds',
    method: 'delete',
    params
  })
}

// @Tags CliMainorder
// @Summary 更新订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainorder true "更新订单总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /climainorder/updateCliMainorder [put]
export const updateCliMainorder = (data) => {
  return service({
    url: '/climainorder/updateCliMainorder',
    method: 'put',
    data
  })
}

// @Tags CliMainorder
// @Summary 用id查询订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliMainorder true "用id查询订单总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /climainorder/findCliMainorder [get]
export const findCliMainorder = (params) => {
  return service({
    url: '/climainorder/findCliMainorder',
    method: 'get',
    params
  })
}

// @Tags CliMainorder
// @Summary 分页获取订单总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单总表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /climainorder/getCliMainorderList [get]
export const getCliMainorderList = (params) => {
  return service({
    url: '/climainorder/getCliMainorderList',
    method: 'get',
    params
  })
}

// @Tags CliMainorder
// @Summary 不需要鉴权的订单总表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainorderSearch true "分页获取订单总表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /climainorder/getCliMainorderPublic [get]
export const getCliMainorderPublic = () => {
  return service({
    url: '/climainorder/getCliMainorderPublic',
    method: 'get',
  })
}
// Buy 客户端购买方法
// @Tags CliMainorder
// @Summary 客户端购买方法
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /climainorder/buy [POST]
export const buy = () => {
  return service({
    url: '/climainorder/buy',
    method: 'POST'
  })
}
