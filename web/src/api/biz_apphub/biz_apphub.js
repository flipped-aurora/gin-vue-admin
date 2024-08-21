import service from '@/utils/request'

// @Tags BizAppHub
// @Summary 创建biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizAppHub true "创建biz_apphub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bizAppHub/createBizAppHub [post]
export const createBizAppHub = (data) => {
  return service({
    url: '/bizAppHub/createBizAppHub',
    method: 'post',
    data
  })
}

// @Tags BizAppHub
// @Summary 删除biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizAppHub true "删除biz_apphub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizAppHub/deleteBizAppHub [delete]
export const deleteBizAppHub = (params) => {
  return service({
    url: '/bizAppHub/deleteBizAppHub',
    method: 'delete',
    params
  })
}

// @Tags BizAppHub
// @Summary 批量删除biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除biz_apphub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizAppHub/deleteBizAppHub [delete]
export const deleteBizAppHubByIds = (params) => {
  return service({
    url: '/bizAppHub/deleteBizAppHubByIds',
    method: 'delete',
    params
  })
}

// @Tags BizAppHub
// @Summary 更新biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizAppHub true "更新biz_apphub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bizAppHub/updateBizAppHub [put]
export const updateBizAppHub = (data) => {
  return service({
    url: '/bizAppHub/updateBizAppHub',
    method: 'put',
    data
  })
}

// @Tags BizAppHub
// @Summary 用id查询biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BizAppHub true "用id查询biz_apphub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bizAppHub/findBizAppHub [get]
export const findBizAppHub = (params) => {
  return service({
    url: '/bizAppHub/findBizAppHub',
    method: 'get',
    params
  })
}

// @Tags BizAppHub
// @Summary 分页获取biz_apphub列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取biz_apphub列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bizAppHub/getBizAppHubList [get]
export const getBizAppHubList = (params) => {
  return service({
    url: '/bizAppHub/getBizAppHubList',
    method: 'get',
    params
  })
}
