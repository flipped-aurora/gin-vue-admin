import service from '@/utils/request'

// @Tags NestInfo
// @Summary 创建NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestInfo true "创建NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestinfo/createNestInfo [post]
export const createNestInfo = (data) => {
  return service({
    url: '/nestinfo/createNestInfo',
    method: 'post',
    data
  })
}

// @Tags NestInfo
// @Summary 删除NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestInfo true "删除NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nestinfo/deleteNestInfo [delete]
export const deleteNestInfo = (data) => {
  return service({
    url: '/nestinfo/deleteNestInfo',
    method: 'delete',
    data
  })
}

// @Tags NestInfo
// @Summary 删除NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nestinfo/deleteNestInfo [delete]
export const deleteNestInfoByIds = (data) => {
  return service({
    url: '/nestinfo/deleteNestInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags NestInfo
// @Summary 更新NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NestInfo true "更新NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /nestinfo/updateNestInfo [put]
export const updateNestInfo = (data) => {
  return service({
    url: '/nestinfo/updateNestInfo',
    method: 'put',
    data
  })
}

// @Tags NestInfo
// @Summary 用id查询NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NestInfo true "用id查询NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /nestinfo/findNestInfo [get]
export const findNestInfo = (params) => {
  return service({
    url: '/nestinfo/findNestInfo',
    method: 'get',
    params
  })
}

// @Tags NestInfo
// @Summary 分页获取NestInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取NestInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestinfo/getNestInfoList [get]
export const getNestInfoList = (params) => {
  return service({
    url: '/nestinfo/getNestInfoList',
    method: 'get',
    params
  })
}
