import service from '@/utils/request'
// @Tags CliSetvip
// @Summary 创建团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliSetvip true "创建团队设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /clisetvip/createCliSetvip [post]
export const createCliSetvip = (data) => {
  return service({
    url: '/clisetvip/createCliSetvip',
    method: 'post',
    data
  })
}

// @Tags CliSetvip
// @Summary 删除团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliSetvip true "删除团队设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clisetvip/deleteCliSetvip [delete]
export const deleteCliSetvip = (params) => {
  return service({
    url: '/clisetvip/deleteCliSetvip',
    method: 'delete',
    params
  })
}

// @Tags CliSetvip
// @Summary 批量删除团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除团队设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clisetvip/deleteCliSetvip [delete]
export const deleteCliSetvipByIds = (params) => {
  return service({
    url: '/clisetvip/deleteCliSetvipByIds',
    method: 'delete',
    params
  })
}

// @Tags CliSetvip
// @Summary 更新团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliSetvip true "更新团队设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clisetvip/updateCliSetvip [put]
export const updateCliSetvip = (data) => {
  return service({
    url: '/clisetvip/updateCliSetvip',
    method: 'put',
    data
  })
}

// @Tags CliSetvip
// @Summary 用id查询团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliSetvip true "用id查询团队设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clisetvip/findCliSetvip [get]
export const findCliSetvip = (params) => {
  return service({
    url: '/clisetvip/findCliSetvip',
    method: 'get',
    params
  })
}

// @Tags CliSetvip
// @Summary 分页获取团队设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取团队设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clisetvip/getCliSetvipList [get]
export const getCliSetvipList = (params) => {
  return service({
    url: '/clisetvip/getCliSetvipList',
    method: 'get',
    params
  })
}

// @Tags CliSetvip
// @Summary 不需要鉴权的团队设置接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliSetvipSearch true "分页获取团队设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /clisetvip/getCliSetvipPublic [get]
export const getCliSetvipPublic = () => {
  return service({
    url: '/clisetvip/getCliSetvipPublic',
    method: 'get',
  })
}
