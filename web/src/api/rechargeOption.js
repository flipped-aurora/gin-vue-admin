import service from '@/utils/request'

// @Tags RechargeOption
// @Summary 创建RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeOption true "创建RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rechargeOption/createRechargeOption [post]
export const createRechargeOption = (data) => {
  return service({
    url: '/rechargeOption/createRechargeOption',
    method: 'post',
    data
  })
}

// @Tags RechargeOption
// @Summary 删除RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeOption true "删除RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeOption/deleteRechargeOption [delete]
export const deleteRechargeOption = (data) => {
  return service({
    url: '/rechargeOption/deleteRechargeOption',
    method: 'delete',
    data
  })
}

// @Tags RechargeOption
// @Summary 删除RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeOption/deleteRechargeOption [delete]
export const deleteRechargeOptionByIds = (data) => {
  return service({
    url: '/rechargeOption/deleteRechargeOptionByIds',
    method: 'delete',
    data
  })
}

// @Tags RechargeOption
// @Summary 更新RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeOption true "更新RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rechargeOption/updateRechargeOption [put]
export const updateRechargeOption = (data) => {
  return service({
    url: '/rechargeOption/updateRechargeOption',
    method: 'put',
    data
  })
}

// @Tags RechargeOption
// @Summary 用id查询RechargeOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RechargeOption true "用id查询RechargeOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rechargeOption/findRechargeOption [get]
export const findRechargeOption = (params) => {
  return service({
    url: '/rechargeOption/findRechargeOption',
    method: 'get',
    params
  })
}

// @Tags RechargeOption
// @Summary 分页获取RechargeOption列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RechargeOption列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rechargeOption/getRechargeOptionList [get]
export const getRechargeOptionList = (params) => {
  return service({
    url: '/rechargeOption/getRechargeOptionList',
    method: 'get',
    params
  })
}
