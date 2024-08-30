import service from '@/utils/request'

// @Tags BizToolCmdApi
// @Summary 创建后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizToolCmdApi true "创建后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bizToolCmdApi/createBizToolCmdApi [post]
export const createBizToolCmdApi = (data) => {
  return service({
    url: '/bizToolCmdApi/createBizToolCmdApi',
    method: 'post',
    data
  })
}

// @Tags BizToolCmdApi
// @Summary 删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizToolCmdApi true "删除后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizToolCmdApi/deleteBizToolCmdApi [delete]
export const deleteBizToolCmdApi = (params) => {
  return service({
    url: '/bizToolCmdApi/deleteBizToolCmdApi',
    method: 'delete',
    params
  })
}

// @Tags BizToolCmdApi
// @Summary 批量删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizToolCmdApi/deleteBizToolCmdApi [delete]
export const deleteBizToolCmdApiByIds = (params) => {
  return service({
    url: '/bizToolCmdApi/deleteBizToolCmdApiByIds',
    method: 'delete',
    params
  })
}

// @Tags BizToolCmdApi
// @Summary 更新后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizToolCmdApi true "更新后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bizToolCmdApi/updateBizToolCmdApi [put]
export const updateBizToolCmdApi = (data) => {
  return service({
    url: '/bizToolCmdApi/updateBizToolCmdApi',
    method: 'put',
    data
  })
}

// @Tags BizToolCmdApi
// @Summary 用id查询后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BizToolCmdApi true "用id查询后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bizToolCmdApi/findBizToolCmdApi [get]
export const findBizToolCmdApi = (params) => {
  return service({
    url: '/bizToolCmdApi/findBizToolCmdApi',
    method: 'get',
    params
  })
}

// @Tags BizToolCmdApi
// @Summary 分页获取后端工具指令api列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取后端工具指令api列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bizToolCmdApi/getBizToolCmdApiList [get]
export const getBizToolCmdApiList = (params) => {
  return service({
    url: '/bizToolCmdApi/getBizToolCmdApiList',
    method: 'get',
    params
  })
}
