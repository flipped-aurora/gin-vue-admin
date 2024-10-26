import service from '@/utils/request'
// @Tags CliProfit
// @Summary 创建结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliProfit true "创建结算详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliprofit/createCliProfit [post]
export const createCliProfit = (data) => {
  return service({
    url: '/cliprofit/createCliProfit',
    method: 'post',
    data
  })
}

// @Tags CliProfit
// @Summary 删除结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliProfit true "删除结算详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliprofit/deleteCliProfit [delete]
export const deleteCliProfit = (params) => {
  return service({
    url: '/cliprofit/deleteCliProfit',
    method: 'delete',
    params
  })
}

// @Tags CliProfit
// @Summary 批量删除结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除结算详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliprofit/deleteCliProfit [delete]
export const deleteCliProfitByIds = (params) => {
  return service({
    url: '/cliprofit/deleteCliProfitByIds',
    method: 'delete',
    params
  })
}

// @Tags CliProfit
// @Summary 更新结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliProfit true "更新结算详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliprofit/updateCliProfit [put]
export const updateCliProfit = (data) => {
  return service({
    url: '/cliprofit/updateCliProfit',
    method: 'put',
    data
  })
}

// @Tags CliProfit
// @Summary 用id查询结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliProfit true "用id查询结算详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliprofit/findCliProfit [get]
export const findCliProfit = (params) => {
  return service({
    url: '/cliprofit/findCliProfit',
    method: 'get',
    params
  })
}

// @Tags CliProfit
// @Summary 分页获取结算详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取结算详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliprofit/getCliProfitList [get]
export const getCliProfitList = (params) => {
  return service({
    url: '/cliprofit/getCliProfitList',
    method: 'get',
    params
  })
}

// @Tags CliProfit
// @Summary 不需要鉴权的结算详情接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliProfitSearch true "分页获取结算详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliprofit/getCliProfitPublic [get]
export const getCliProfitPublic = () => {
  return service({
    url: '/cliprofit/getCliProfitPublic',
    method: 'get',
  })
}
