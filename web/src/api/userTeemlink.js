import service from '@/utils/request'

// @Tags UserTeemlink
// @Summary 创建UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserTeemlink true "创建UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usertl/createUserTeemlink [post]
export const createUserTeemlink = (data) => {
  return service({
    url: '/usertl/createUserTeemlink',
    method: 'post',
    data
  })
}

// @Tags UserTeemlink
// @Summary 删除UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserTeemlink true "删除UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /usertl/deleteUserTeemlink [delete]
export const deleteUserTeemlink = (data) => {
  return service({
    url: '/usertl/deleteUserTeemlink',
    method: 'delete',
    data
  })
}

// @Tags UserTeemlink
// @Summary 删除UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /usertl/deleteUserTeemlink [delete]
export const deleteUserTeemlinkByIds = (data) => {
  return service({
    url: '/usertl/deleteUserTeemlinkByIds',
    method: 'delete',
    data
  })
}

// @Tags UserTeemlink
// @Summary 更新UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserTeemlink true "更新UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /usertl/updateUserTeemlink [put]
export const updateUserTeemlink = (data) => {
  return service({
    url: '/usertl/updateUserTeemlink',
    method: 'put',
    data
  })
}

// @Tags UserTeemlink
// @Summary 用id查询UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserTeemlink true "用id查询UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /usertl/findUserTeemlink [get]
export const findUserTeemlink = (params) => {
  return service({
    url: '/usertl/findUserTeemlink',
    method: 'get',
    params
  })
}

// @Tags UserTeemlink
// @Summary 分页获取UserTeemlink列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserTeemlink列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usertl/getUserTeemlinkList [get]
export const getUserTeemlinkList = (params) => {
  return service({
    url: '/usertl/getUserTeemlinkList',
    method: 'get',
    params
  })
}
