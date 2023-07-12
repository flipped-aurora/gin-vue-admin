import service from '@/utils/request'

// @Tags NestRole
// @Summary 创建NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestRole true "创建NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestrole/createNestRole [post]
export const createNestRole = (data) => {
  return service({
    url: '/nestrole/createNestRole',
    method: 'post',
    data
  })
}

// @Tags NestRole
// @Summary 删除NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestRole true "删除NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nestrole/deleteNestRole [delete]
export const deleteNestRole = (data) => {
  return service({
    url: '/nestrole/deleteNestRole',
    method: 'delete',
    data
  })
}

// @Tags NestRole
// @Summary 删除NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nestrole/deleteNestRole [delete]
export const deleteNestRoleByIds = (data) => {
  return service({
    url: '/nestrole/deleteNestRoleByIds',
    method: 'delete',
    data
  })
}

// @Tags NestRole
// @Summary 更新NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestRole true "更新NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /nestrole/updateNestRole [put]
export const updateNestRole = (data) => {
  return service({
    url: '/nestrole/updateNestRole',
    method: 'put',
    data
  })
}

// @Tags NestRole
// @Summary 用id查询NestRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NestRole true "用id查询NestRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /nestrole/findNestRole [get]
export const findNestRole = (params) => {
  return service({
    url: '/nestrole/findNestRole',
    method: 'get',
    params
  })
}

// @Tags NestRole
// @Summary 分页获取NestRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取NestRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestrole/getNestRoleList [get]
export const getNestRoleList = (params) => {
  return service({
    url: '/nestrole/getNestRoleList',
    method: 'get',
    params
  })
}
