import service from '@/utils/request'
// @Tags CliLoad
// @Summary 创建cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliLoad true "创建cliLoad表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliLoad/createCliLoad [post]
export const createCliLoad = (data) => {
  return service({
    url: '/cliLoad/createCliLoad',
    method: 'post',
    data
  })
}

// @Tags CliLoad
// @Summary 删除cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliLoad true "删除cliLoad表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliLoad/deleteCliLoad [delete]
export const deleteCliLoad = (params) => {
  return service({
    url: '/cliLoad/deleteCliLoad',
    method: 'delete',
    params
  })
}

// @Tags CliLoad
// @Summary 批量删除cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cliLoad表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliLoad/deleteCliLoad [delete]
export const deleteCliLoadByIds = (params) => {
  return service({
    url: '/cliLoad/deleteCliLoadByIds',
    method: 'delete',
    params
  })
}

// @Tags CliLoad
// @Summary 更新cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliLoad true "更新cliLoad表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliLoad/updateCliLoad [put]
export const updateCliLoad = (data) => {
  return service({
    url: '/cliLoad/updateCliLoad',
    method: 'put',
    data
  })
}

// @Tags CliLoad
// @Summary 用id查询cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliLoad true "用id查询cliLoad表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliLoad/findCliLoad [get]
export const findCliLoad = (params) => {
  return service({
    url: '/cliLoad/findCliLoad',
    method: 'get',
    params
  })
}

// @Tags CliLoad
// @Summary 分页获取cliLoad表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cliLoad表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliLoad/getCliLoadList [get]
export const getCliLoadList = (params) => {
  return service({
    url: '/cliLoad/getCliLoadList',
    method: 'get',
    params
  })
}

// @Tags CliLoad
// @Summary 不需要鉴权的cliLoad表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliLoadSearch true "分页获取cliLoad表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliLoad/getCliLoadPublic [get]
export const getCliLoadPublic = () => {
  return service({
    url: '/cliLoad/getCliLoadPublic',
    method: 'get',
  })
}
// Login 用户登录方法
// @Tags CliLoad
// @Summary 用户登录方法
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /cliLoad/login [POST]
export const login = () => {
  return service({
    url: '/cliLoad/login',
    method: 'POST'
  })
}
