import service from '@/utils/request'

// @Tags AppRole
// @Summary 创建AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRole true "创建AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRole/createAppRole [post]
export const createAppRole = (data) => {
  return service({
    url: '/appRole/createAppRole',
    method: 'post',
    data
  })
}

// @Tags AppRole
// @Summary 删除AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRole true "删除AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRole/deleteAppRole [delete]
export const deleteAppRole = (data) => {
  return service({
    url: '/appRole/deleteAppRole',
    method: 'delete',
    data
  })
}

// @Tags AppRole
// @Summary 删除AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRole/deleteAppRole [delete]
export const deleteAppRoleByIds = (data) => {
  return service({
    url: '/appRole/deleteAppRoleByIds',
    method: 'delete',
    data
  })
}

// @Tags AppRole
// @Summary 更新AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRole true "更新AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appRole/updateAppRole [put]
export const updateAppRole = (data) => {
  return service({
    url: '/appRole/updateAppRole',
    method: 'put',
    data
  })
}

// @Tags AppRole
// @Summary 用id查询AppRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppRole true "用id查询AppRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appRole/findAppRole [get]
export const findAppRole = (params) => {
  return service({
    url: '/appRole/findAppRole',
    method: 'get',
    params
  })
}

// @Tags AppRole
// @Summary 分页获取AppRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRole/getAppRoleList [get]
export const getAppRoleList = (params) => {
  return service({
    url: '/appRole/getAppRoleList',
    method: 'get',
    params
  })
}
