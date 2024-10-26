import service from '@/utils/request'
// @Tags CliMainwith
// @Summary 创建提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainwith true "创建提币总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /climainwith/createCliMainwith [post]
export const createCliMainwith = (data) => {
  return service({
    url: '/climainwith/createCliMainwith',
    method: 'post',
    data
  })
}

// @Tags CliMainwith
// @Summary 删除提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainwith true "删除提币总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /climainwith/deleteCliMainwith [delete]
export const deleteCliMainwith = (params) => {
  return service({
    url: '/climainwith/deleteCliMainwith',
    method: 'delete',
    params
  })
}

// @Tags CliMainwith
// @Summary 批量删除提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除提币总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /climainwith/deleteCliMainwith [delete]
export const deleteCliMainwithByIds = (params) => {
  return service({
    url: '/climainwith/deleteCliMainwithByIds',
    method: 'delete',
    params
  })
}

// @Tags CliMainwith
// @Summary 更新提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliMainwith true "更新提币总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /climainwith/updateCliMainwith [put]
export const updateCliMainwith = (data) => {
  return service({
    url: '/climainwith/updateCliMainwith',
    method: 'put',
    data
  })
}

// @Tags CliMainwith
// @Summary 用id查询提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliMainwith true "用id查询提币总表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /climainwith/findCliMainwith [get]
export const findCliMainwith = (params) => {
  return service({
    url: '/climainwith/findCliMainwith',
    method: 'get',
    params
  })
}

// @Tags CliMainwith
// @Summary 分页获取提币总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取提币总表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /climainwith/getCliMainwithList [get]
export const getCliMainwithList = (params) => {
  return service({
    url: '/climainwith/getCliMainwithList',
    method: 'get',
    params
  })
}

// @Tags CliMainwith
// @Summary 不需要鉴权的提币总表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainwithSearch true "分页获取提币总表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /climainwith/getCliMainwithPublic [get]
export const getCliMainwithPublic = () => {
  return service({
    url: '/climainwith/getCliMainwithPublic',
    method: 'get',
  })
}
