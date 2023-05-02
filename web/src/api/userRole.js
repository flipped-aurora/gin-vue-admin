import service from '@/utils/request'

// @Tags UserRole
// @Summary 创建UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserRole true "创建UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userRole/createUserRole [post]
export const createUserRole = (data) => {
  return service({
    url: '/userRole/createUserRole',
    method: 'post',
    data
  })
}

// @Tags UserRole
// @Summary 删除UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserRole true "删除UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userRole/deleteUserRole [delete]
export const deleteUserRole = (data) => {
  return service({
    url: '/userRole/deleteUserRole',
    method: 'delete',
    data
  })
}

// @Tags UserRole
// @Summary 删除UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userRole/deleteUserRole [delete]
export const deleteUserRoleByIds = (data) => {
  return service({
    url: '/userRole/deleteUserRoleByIds',
    method: 'delete',
    data
  })
}

// @Tags UserRole
// @Summary 更新UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserRole true "更新UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userRole/updateUserRole [put]
export const updateUserRole = (data) => {
  return service({
    url: '/userRole/updateUserRole',
    method: 'put',
    data
  })
}

// @Tags UserRole
// @Summary 用id查询UserRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserRole true "用id查询UserRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userRole/findUserRole [get]
export const findUserRole = (params) => {
  return service({
    url: '/userRole/findUserRole',
    method: 'get',
    params
  })
}

// @Tags UserRole
// @Summary 分页获取UserRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userRole/getUserRoleList [get]
export const getUserRoleList = (params) => {
  return service({
    url: '/userRole/getUserRoleList',
    method: 'get',
    params
  })
}
