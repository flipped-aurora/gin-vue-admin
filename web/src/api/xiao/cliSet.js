import service from '@/utils/request'
// @Tags CliSet
// @Summary 创建结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliSet true "创建结算设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliset/createCliSet [post]
export const createCliSet = (data) => {
  return service({
    url: '/cliset/createCliSet',
    method: 'post',
    data
  })
}

// @Tags CliSet
// @Summary 删除结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliSet true "删除结算设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliset/deleteCliSet [delete]
export const deleteCliSet = (params) => {
  return service({
    url: '/cliset/deleteCliSet',
    method: 'delete',
    params
  })
}

// @Tags CliSet
// @Summary 批量删除结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除结算设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliset/deleteCliSet [delete]
export const deleteCliSetByIds = (params) => {
  return service({
    url: '/cliset/deleteCliSetByIds',
    method: 'delete',
    params
  })
}

// @Tags CliSet
// @Summary 更新结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliSet true "更新结算设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliset/updateCliSet [put]
export const updateCliSet = (data) => {
  return service({
    url: '/cliset/updateCliSet',
    method: 'put',
    data
  })
}

// @Tags CliSet
// @Summary 用id查询结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliSet true "用id查询结算设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliset/findCliSet [get]
export const findCliSet = (params) => {
  return service({
    url: '/cliset/findCliSet',
    method: 'get',
    params
  })
}

// @Tags CliSet
// @Summary 分页获取结算设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取结算设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliset/getCliSetList [get]
export const getCliSetList = (params) => {
  return service({
    url: '/cliset/getCliSetList',
    method: 'get',
    params
  })
}

// @Tags CliSet
// @Summary 不需要鉴权的结算设置接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliSetSearch true "分页获取结算设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliset/getCliSetPublic [get]
export const getCliSetPublic = () => {
  return service({
    url: '/cliset/getCliSetPublic',
    method: 'get',
  })
}
