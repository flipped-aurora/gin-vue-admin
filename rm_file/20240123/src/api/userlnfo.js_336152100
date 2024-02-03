import service from '@/utils/request'

// @Tags UserInfo
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserInfo true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userData/createUserInfo [post]
export const createUserInfo = (data) => {
  return service({
    url: '/userData/createUserInfo',
    method: 'post',
    data
  })
}

// @Tags UserInfo
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserInfo true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userData/deleteUserInfo [delete]
export const deleteUserInfo = (params) => {
  return service({
    url: '/userData/deleteUserInfo',
    method: 'delete',
    params
  })
}

// @Tags UserInfo
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userData/deleteUserInfo [delete]
export const deleteUserInfoByIds = (params) => {
  return service({
    url: '/userData/deleteUserInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags UserInfo
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserInfo true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userData/updateUserInfo [put]
export const updateUserInfo = (data) => {
  return service({
    url: '/userData/updateUserInfo',
    method: 'put',
    data
  })
}

// @Tags UserInfo
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserInfo true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userData/findUserInfo [get]
export const findUserInfo = (params) => {
  return service({
    url: '/userData/findUserInfo',
    method: 'get',
    params
  })
}

// @Tags UserInfo
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userData/getUserInfoList [get]
export const getUserInfoList = (params) => {
  return service({
    url: '/userData/getUserInfoList',
    method: 'get',
    params
  })
}
