import service from '@/utils/request'

// @Tags BizCloudFunction
// @Summary 创建云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizCloudFunction true "创建云函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bizCloudFunction/createBizCloudFunction [post]
export const createBizCloudFunction = (data) => {
  return service({
    url: '/bizCloudFunction/createBizCloudFunction',
    method: 'post',
    data
  })
}

// @Tags BizCloudFunction
// @Summary 删除云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizCloudFunction true "删除云函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizCloudFunction/deleteBizCloudFunction [delete]
export const deleteBizCloudFunction = (params) => {
  return service({
    url: '/bizCloudFunction/deleteBizCloudFunction',
    method: 'delete',
    params
  })
}

// @Tags BizCloudFunction
// @Summary 批量删除云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除云函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizCloudFunction/deleteBizCloudFunction [delete]
export const deleteBizCloudFunctionByIds = (params) => {
  return service({
    url: '/bizCloudFunction/deleteBizCloudFunctionByIds',
    method: 'delete',
    params
  })
}

// @Tags BizCloudFunction
// @Summary 更新云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizCloudFunction true "更新云函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bizCloudFunction/updateBizCloudFunction [put]
export const updateBizCloudFunction = (data) => {
  return service({
    url: '/bizCloudFunction/updateBizCloudFunction',
    method: 'put',
    data
  })
}

// @Tags BizCloudFunction
// @Summary 用id查询云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BizCloudFunction true "用id查询云函数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bizCloudFunction/findBizCloudFunction [get]
export const findBizCloudFunction = (params) => {
  return service({
    url: '/bizCloudFunction/findBizCloudFunction',
    method: 'get',
    params
  })
}

// @Tags BizCloudFunction
// @Summary 分页获取云函数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取云函数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bizCloudFunction/getBizCloudFunctionList [get]
export const getBizCloudFunctionList = (params) => {
  return service({
    url: '/bizCloudFunction/getBizCloudFunctionList',
    method: 'get',
    params
  })
}
