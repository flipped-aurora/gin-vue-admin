import service from '@/utils/request'
// @Tags CliMainprofit
// @Summary 创建结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainprofit true "创建结算总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /climainprofit/createCliMainprofit [post]
export const createCliMainprofit = (data) => {
  return service({
    url: '/climainprofit/createCliMainprofit',
    method: 'post',
    data
  })
}

// @Tags CliMainprofit
// @Summary 删除结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainprofit true "删除结算总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /climainprofit/deleteCliMainprofit [delete]
export const deleteCliMainprofit = (params) => {
  return service({
    url: '/climainprofit/deleteCliMainprofit',
    method: 'delete',
    params
  })
}

// @Tags CliMainprofit
// @Summary 批量删除结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除结算总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /climainprofit/deleteCliMainprofit [delete]
export const deleteCliMainprofitByIds = (params) => {
  return service({
    url: '/climainprofit/deleteCliMainprofitByIds',
    method: 'delete',
    params
  })
}

// @Tags CliMainprofit
// @Summary 更新结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainprofit true "更新结算总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /climainprofit/updateCliMainprofit [put]
export const updateCliMainprofit = (data) => {
  return service({
    url: '/climainprofit/updateCliMainprofit',
    method: 'put',
    data
  })
}

// @Tags CliMainprofit
// @Summary 用id查询结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliMainprofit true "用id查询结算总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /climainprofit/findCliMainprofit [get]
export const findCliMainprofit = (params) => {
  return service({
    url: '/climainprofit/findCliMainprofit',
    method: 'get',
    params
  })
}

// @Tags CliMainprofit
// @Summary 分页获取结算总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取结算总表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /climainprofit/getCliMainprofitList [get]
export const getCliMainprofitList = (params) => {
  return service({
    url: '/climainprofit/getCliMainprofitList',
    method: 'get',
    params
  })
}

// @Tags CliMainprofit
// @Summary 不需要鉴权的结算总表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainprofitSearch true "分页获取结算总表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /climainprofit/getCliMainprofitPublic [get]
export const getCliMainprofitPublic = () => {
  return service({
    url: '/climainprofit/getCliMainprofitPublic',
    method: 'get',
  })
}
// PullProfit 推荐结算
// @Tags CliMainprofit
// @Summary 推荐结算
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /climainprofit/pullprofit [GET]
export const pullprofit = () => {
  return service({
    url: '/climainprofit/pullprofit',
    method: 'GET'
  })
}
// TeamProfit 团队结算
// @Tags CliMainprofit
// @Summary 团队结算
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /climainprofit/teamprofit [GET]
export const teamprofit = () => {
  return service({
    url: '/climainprofit/teamprofit',
    method: 'GET'
  })
}
