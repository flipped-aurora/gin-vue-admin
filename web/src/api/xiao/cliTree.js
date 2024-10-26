import service from '@/utils/request'
// @Tags CliTree
// @Summary 创建用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliTree true "创建用户关系表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cliTree/createCliTree [post]
export const createCliTree = (data) => {
  return service({
    url: '/cliTree/createCliTree',
    method: 'post',
    data
  })
}

// @Tags CliTree
// @Summary 删除用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliTree true "删除用户关系表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliTree/deleteCliTree [delete]
export const deleteCliTree = (params) => {
  return service({
    url: '/cliTree/deleteCliTree',
    method: 'delete',
    params
  })
}

// @Tags CliTree
// @Summary 批量删除用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户关系表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cliTree/deleteCliTree [delete]
export const deleteCliTreeByIds = (params) => {
  return service({
    url: '/cliTree/deleteCliTreeByIds',
    method: 'delete',
    params
  })
}

// @Tags CliTree
// @Summary 更新用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CliTree true "更新用户关系表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cliTree/updateCliTree [put]
export const updateCliTree = (data) => {
  return service({
    url: '/cliTree/updateCliTree',
    method: 'put',
    data
  })
}

// @Tags CliTree
// @Summary 用id查询用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CliTree true "用id查询用户关系表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cliTree/findCliTree [get]
export const findCliTree = (params) => {
  return service({
    url: '/cliTree/findCliTree',
    method: 'get',
    params
  })
}

// @Tags CliTree
// @Summary 分页获取用户关系表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户关系表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cliTree/getCliTreeList [get]
export const getCliTreeList = (params) => {
  return service({
    url: '/cliTree/getCliTreeList',
    method: 'get',
    params
  })
}

// @Tags CliTree
// @Summary 不需要鉴权的用户关系表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliTreeSearch true "分页获取用户关系表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliTree/getCliTreePublic [get]
export const getCliTreePublic = () => {
  return service({
    url: '/cliTree/getCliTreePublic',
    method: 'get',
  })
}
