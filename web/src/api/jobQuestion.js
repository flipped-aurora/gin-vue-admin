import service from '@/utils/request'

// @Tags JobQuestion
// @Summary 创建JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JobQuestion true "创建JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobQuestion/createJobQuestion [post]
export const createJobQuestion = (data) => {
  return service({
    url: '/jobQuestion/createJobQuestion',
    method: 'post',
    data
  })
}

// @Tags JobQuestion
// @Summary 删除JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JobQuestion true "删除JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jobQuestion/deleteJobQuestion [delete]
export const deleteJobQuestion = (data) => {
  return service({
    url: '/jobQuestion/deleteJobQuestion',
    method: 'delete',
    data
  })
}

// @Tags JobQuestion
// @Summary 删除JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jobQuestion/deleteJobQuestion [delete]
export const deleteJobQuestionByIds = (data) => {
  return service({
    url: '/jobQuestion/deleteJobQuestionByIds',
    method: 'delete',
    data
  })
}

// @Tags JobQuestion
// @Summary 更新JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JobQuestion true "更新JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jobQuestion/updateJobQuestion [put]
export const updateJobQuestion = (data) => {
  return service({
    url: '/jobQuestion/updateJobQuestion',
    method: 'put',
    data
  })
}

// @Tags JobQuestion
// @Summary 用id查询JobQuestion
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JobQuestion true "用id查询JobQuestion"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jobQuestion/findJobQuestion [get]
export const findJobQuestion = (params) => {
  return service({
    url: '/jobQuestion/findJobQuestion',
    method: 'get',
    params
  })
}

// @Tags JobQuestion
// @Summary 分页获取JobQuestion列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取JobQuestion列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobQuestion/getJobQuestionList [get]
export const getJobQuestionList = (params) => {
  return service({
    url: '/jobQuestion/getJobQuestionList',
    method: 'get',
    params
  })
}
