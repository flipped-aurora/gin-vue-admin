import service from '@/utils/request'

// @Tags JobApply
// @Summary 创建JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JobApply true "创建JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobApply/createJobApply [post]
export const createJobApply = (data) => {
  return service({
    url: '/jobApply/createJobApply',
    method: 'post',
    data
  })
}

// @Tags JobApply
// @Summary 删除JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JobApply true "删除JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jobApply/deleteJobApply [delete]
export const deleteJobApply = (data) => {
  return service({
    url: '/jobApply/deleteJobApply',
    method: 'delete',
    data
  })
}

// @Tags JobApply
// @Summary 删除JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jobApply/deleteJobApply [delete]
export const deleteJobApplyByIds = (data) => {
  return service({
    url: '/jobApply/deleteJobApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags JobApply
// @Summary 更新JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JobApply true "更新JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jobApply/updateJobApply [put]
export const updateJobApply = (data) => {
  return service({
    url: '/jobApply/updateJobApply',
    method: 'put',
    data
  })
}

// @Tags JobApply
// @Summary 用id查询JobApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JobApply true "用id查询JobApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jobApply/findJobApply [get]
export const findJobApply = (params) => {
  return service({
    url: '/jobApply/findJobApply',
    method: 'get',
    params
  })
}

// @Tags JobApply
// @Summary 分页获取JobApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取JobApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobApply/getJobApplyList [get]
export const getJobApplyList = (params) => {
  return service({
    url: '/jobApply/getJobApplyList',
    method: 'get',
    params
  })
}
