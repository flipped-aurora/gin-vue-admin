import service from '@/utils/request'

// @Tags ComInfo
// @Summary 创建比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ComInfo true "创建比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /comData/createComInfo [post]
export const createComInfo = (data) => {
  return service({
    url: '/comData/createComInfo',
    method: 'post',
    data
  })
}

// @Tags ComInfo
// @Summary 删除比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ComInfo true "删除比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comData/deleteComInfo [delete]
export const deleteComInfo = (params) => {
  return service({
    url: '/comData/deleteComInfo',
    method: 'delete',
    params
  })
}

// @Tags ComInfo
// @Summary 批量删除比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comData/deleteComInfo [delete]
export const deleteComInfoByIds = (params) => {
  return service({
    url: '/comData/deleteComInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags ComInfo
// @Summary 更新比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ComInfo true "更新比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comData/updateComInfo [put]
export const updateComInfo = (data) => {
  return service({
    url: '/comData/updateComInfo',
    method: 'put',
    data
  })
}

// @Tags ComInfo
// @Summary 用id查询比赛信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ComInfo true "用id查询比赛信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comData/findComInfo [get]
export const findComInfo = (params) => {
  return service({
    url: '/comData/findComInfo',
    method: 'get',
    params
  })
}

// @Tags ComInfo
// @Summary 分页获取比赛信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取比赛信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comData/getComInfoList [get]
export const getComInfoList = (params) => {
  return service({
    url: '/comData/getComInfoList',
    method: 'get',
    params
  })
}
