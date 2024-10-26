import service from '@/utils/request'
// @Tags CliWithdraw
// @Summary 创建提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliWithdraw true "创建提币详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliwithdraw/createCliWithdraw [post]
export const createCliWithdraw = (data) => {
  return service({
    url: '/cliwithdraw/createCliWithdraw',
    method: 'post',
    data
  })
}

// @Tags CliWithdraw
// @Summary 删除提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliWithdraw true "删除提币详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliwithdraw/deleteCliWithdraw [delete]
export const deleteCliWithdraw = (params) => {
  return service({
    url: '/cliwithdraw/deleteCliWithdraw',
    method: 'delete',
    params
  })
}

// @Tags CliWithdraw
// @Summary 批量删除提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除提币详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliwithdraw/deleteCliWithdraw [delete]
export const deleteCliWithdrawByIds = (params) => {
  return service({
    url: '/cliwithdraw/deleteCliWithdrawByIds',
    method: 'delete',
    params
  })
}

// @Tags CliWithdraw
// @Summary 更新提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliWithdraw true "更新提币详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliwithdraw/updateCliWithdraw [put]
export const updateCliWithdraw = (data) => {
  return service({
    url: '/cliwithdraw/updateCliWithdraw',
    method: 'put',
    data
  })
}

// @Tags CliWithdraw
// @Summary 用id查询提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliWithdraw true "用id查询提币详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliwithdraw/findCliWithdraw [get]
export const findCliWithdraw = (params) => {
  return service({
    url: '/cliwithdraw/findCliWithdraw',
    method: 'get',
    params
  })
}

// @Tags CliWithdraw
// @Summary 分页获取提币详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取提币详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliwithdraw/getCliWithdrawList [get]
export const getCliWithdrawList = (params) => {
  return service({
    url: '/cliwithdraw/getCliWithdrawList',
    method: 'get',
    params
  })
}

// @Tags CliWithdraw
// @Summary 不需要鉴权的提币详情接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliWithdrawSearch true "分页获取提币详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliwithdraw/getCliWithdrawPublic [get]
export const getCliWithdrawPublic = () => {
  return service({
    url: '/cliwithdraw/getCliWithdrawPublic',
    method: 'get',
  })
}
