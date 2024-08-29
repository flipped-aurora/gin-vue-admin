import service from '@/utils/request'

// @Tags BizCmdToolApi
// @Summary 创建后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizCmdToolApi true "创建后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bizCmdToolApi/createBizCmdToolApi [post]
export const createBizCmdToolApi = (data) => {
  return service({
    url: '/bizCmdToolApi/createBizCmdToolApi',
    method: 'post',
    data
  })
}

// @Tags BizCmdToolApi
// @Summary 删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizCmdToolApi true "删除后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizCmdToolApi/deleteBizCmdToolApi [delete]
export const deleteBizCmdToolApi = (params) => {
  return service({
    url: '/bizCmdToolApi/deleteBizCmdToolApi',
    method: 'delete',
    params
  })
}

// @Tags BizCmdToolApi
// @Summary 批量删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizCmdToolApi/deleteBizCmdToolApi [delete]
export const deleteBizCmdToolApiByIds = (params) => {
  return service({
    url: '/bizCmdToolApi/deleteBizCmdToolApiByIds',
    method: 'delete',
    params
  })
}

// @Tags BizCmdToolApi
// @Summary 更新后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizCmdToolApi true "更新后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bizCmdToolApi/updateBizCmdToolApi [put]
export const updateBizCmdToolApi = (data) => {
  return service({
    url: '/bizCmdToolApi/updateBizCmdToolApi',
    method: 'put',
    data
  })
}

// @Tags BizCmdToolApi
// @Summary 用id查询后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BizCmdToolApi true "用id查询后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bizCmdToolApi/findBizCmdToolApi [get]
export const findBizCmdToolApi = (params) => {
  return service({
    url: '/bizCmdToolApi/findBizCmdToolApi',
    method: 'get',
    params
  })
}

// @Tags BizCmdToolApi
// @Summary 分页获取后端工具指令api列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取后端工具指令api列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bizCmdToolApi/getBizCmdToolApiList [get]
export const getBizCmdToolApiList = (params) => {
  return service({
    url: '/bizCmdToolApi/getBizCmdToolApiList',
    method: 'get',
    params
  })
}
