import service from '@/utils/request'

// @Tags Job
// @Summary 创建Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "创建Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /job/createJob [post]
export const createJob = (data) => {
  return service({
    url: '/job/createJob',
    method: 'post',
    data
  })
}

// @Tags Job
// @Summary 删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /job/deleteJob [delete]
export const deleteJob = (data) => {
  return service({
    url: '/job/deleteJob',
    method: 'delete',
    data
  })
}

// @Tags Job
// @Summary 删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /job/deleteJob [delete]
export const deleteJobByIds = (data) => {
  return service({
    url: '/job/deleteJobByIds',
    method: 'delete',
    data
  })
}

// @Tags Job
// @Summary 更新Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "更新Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /job/updateJob [put]
export const updateJob = (data) => {
  return service({
    url: '/job/updateJob',
    method: 'put',
    data
  })
}

// @Tags Job
// @Summary 用id查询Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Job true "用id查询Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /job/findJob [get]
export const findJob = (params) => {
  return service({
    url: '/job/findJob',
    method: 'get',
    params
  })
}

// @Tags Job
// @Summary 分页获取Job列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Job列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /job/getJobList [get]
export const getJobList = (params) => {
  return service({
    url: '/job/getJobList',
    method: 'get',
    params
  })
}
