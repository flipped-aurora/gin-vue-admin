import service from '@/utils/request'

// @Tags BizToolCmdSrvApi
// @Summary 创建后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizToolCmdSrvApi true "创建后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bizToolCmdSrvApi/createBizToolCmdSrvApi [post]
export const createBizToolCmdSrvApi = (data) => {
  return service({
    url: '/bizToolCmdSrvApi/createBizToolCmdSrvApi',
    method: 'post',
    data
  })
}

// @Tags BizToolCmdSrvApi
// @Summary 删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizToolCmdSrvApi true "删除后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizToolCmdSrvApi/deleteBizToolCmdSrvApi [delete]
export const deleteBizToolCmdSrvApi = (params) => {
  return service({
    url: '/bizToolCmdSrvApi/deleteBizToolCmdSrvApi',
    method: 'delete',
    params
  })
}

// @Tags BizToolCmdSrvApi
// @Summary 批量删除后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bizToolCmdSrvApi/deleteBizToolCmdSrvApi [delete]
export const deleteBizToolCmdSrvApiByIds = (params) => {
  return service({
    url: '/bizToolCmdSrvApi/deleteBizToolCmdSrvApiByIds',
    method: 'delete',
    params
  })
}

// @Tags BizToolCmdSrvApi
// @Summary 更新后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BizToolCmdSrvApi true "更新后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bizToolCmdSrvApi/updateBizToolCmdSrvApi [put]
export const updateBizToolCmdSrvApi = (data) => {
  return service({
    url: '/bizToolCmdSrvApi/updateBizToolCmdSrvApi',
    method: 'put',
    data
  })
}

// @Tags BizToolCmdSrvApi
// @Summary 用id查询后端工具指令api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BizToolCmdSrvApi true "用id查询后端工具指令api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bizToolCmdSrvApi/findBizToolCmdSrvApi [get]
export const findBizToolCmdSrvApi = (params) => {
  return service({
    url: '/bizToolCmdSrvApi/findBizToolCmdSrvApi',
    method: 'get',
    params
  })
}

// @Tags BizToolCmdSrvApi
// @Summary 分页获取后端工具指令api列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取后端工具指令api列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bizToolCmdSrvApi/getBizToolCmdSrvApiList [get]
export const getBizToolCmdSrvApiList = (params) => {
  return service({
    url: '/bizToolCmdSrvApi/getBizToolCmdSrvApiList',
    method: 'get',
    params
  })
}
