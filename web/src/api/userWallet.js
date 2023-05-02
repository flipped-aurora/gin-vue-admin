import service from '@/utils/request'

// @Tags UserWallet
// @Summary 创建UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserWallet true "创建UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userWallet/createUserWallet [post]
export const createUserWallet = (data) => {
  return service({
    url: '/userWallet/createUserWallet',
    method: 'post',
    data
  })
}

// @Tags UserWallet
// @Summary 删除UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserWallet true "删除UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userWallet/deleteUserWallet [delete]
export const deleteUserWallet = (data) => {
  return service({
    url: '/userWallet/deleteUserWallet',
    method: 'delete',
    data
  })
}

// @Tags UserWallet
// @Summary 删除UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userWallet/deleteUserWallet [delete]
export const deleteUserWalletByIds = (data) => {
  return service({
    url: '/userWallet/deleteUserWalletByIds',
    method: 'delete',
    data
  })
}

// @Tags UserWallet
// @Summary 更新UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserWallet true "更新UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userWallet/updateUserWallet [put]
export const updateUserWallet = (data) => {
  return service({
    url: '/userWallet/updateUserWallet',
    method: 'put',
    data
  })
}

// @Tags UserWallet
// @Summary 用id查询UserWallet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserWallet true "用id查询UserWallet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userWallet/findUserWallet [get]
export const findUserWallet = (params) => {
  return service({
    url: '/userWallet/findUserWallet',
    method: 'get',
    params
  })
}

// @Tags UserWallet
// @Summary 分页获取UserWallet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserWallet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userWallet/getUserWalletList [get]
export const getUserWalletList = (params) => {
  return service({
    url: '/userWallet/getUserWalletList',
    method: 'get',
    params
  })
}
