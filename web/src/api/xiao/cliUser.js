import service from '@/utils/request'
// @Tags CliUser
// @Summary 创建cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliUser true "创建cliUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliUser/createCliUser [post]
export const createCliUser = (data) => {
  return service({
    url: '/cliUser/createCliUser',
    method: 'post',
    data
  })
}

// @Tags CliUser
// @Summary 删除cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliUser true "删除cliUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliUser/deleteCliUser [delete]
export const deleteCliUser = (params) => {
  return service({
    url: '/cliUser/deleteCliUser',
    method: 'delete',
    params
  })
}

// @Tags CliUser
// @Summary 批量删除cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cliUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliUser/deleteCliUser [delete]
export const deleteCliUserByIds = (params) => {
  return service({
    url: '/cliUser/deleteCliUserByIds',
    method: 'delete',
    params
  })
}

// @Tags CliUser
// @Summary 更新cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliUser true "更新cliUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliUser/updateCliUser [put]
export const updateCliUser = (data) => {
  return service({
    url: '/cliUser/updateCliUser',
    method: 'put',
    data
  })
}

// @Tags CliUser
// @Summary 用id查询cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliUser true "用id查询cliUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliUser/findCliUser [get]
export const findCliUser = (params) => {
  return service({
    url: '/cliUser/findCliUser',
    method: 'get',
    params
  })
}

// @Tags CliUser
// @Summary 分页获取cliUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cliUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliUser/getCliUserList [get]
export const getCliUserList = (params) => {
  return service({
    url: '/cliUser/getCliUserList',
    method: 'get',
    params
  })
}

// @Tags CliUser
// @Summary 不需要鉴权的cliUser表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliUserSearch true "分页获取cliUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliUser/getCliUserPublic [get]
export const getCliUserPublic = () => {
  return service({
    url: '/cliUser/getCliUserPublic',
    method: 'get',
  })
}
// Rigester 客户端注册方法
// @Tags CliUser
// @Summary 客户端注册方法
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /cliUser/rigester [POST]
export const rigester = () => {
  return service({
    url: '/cliUser/rigester',
    method: 'POST'
  })
}
