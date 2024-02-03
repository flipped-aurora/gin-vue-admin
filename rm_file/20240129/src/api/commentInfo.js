import service from '@/utils/request'

// @Tags CommentInfo
// @Summary 创建评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentInfo true "创建评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /commentData/createCommentInfo [post]
export const createCommentInfo = (data) => {
  return service({
    url: '/commentData/createCommentInfo',
    method: 'post',
    data
  })
}

// @Tags CommentInfo
// @Summary 删除评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentInfo true "删除评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentData/deleteCommentInfo [delete]
export const deleteCommentInfo = (params) => {
  return service({
    url: '/commentData/deleteCommentInfo',
    method: 'delete',
    params
  })
}

// @Tags CommentInfo
// @Summary 批量删除评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentData/deleteCommentInfo [delete]
export const deleteCommentInfoByIds = (params) => {
  return service({
    url: '/commentData/deleteCommentInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags CommentInfo
// @Summary 更新评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentInfo true "更新评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commentData/updateCommentInfo [put]
export const updateCommentInfo = (data) => {
  return service({
    url: '/commentData/updateCommentInfo',
    method: 'put',
    data
  })
}

// @Tags CommentInfo
// @Summary 用id查询评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CommentInfo true "用id查询评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commentData/findCommentInfo [get]
export const findCommentInfo = (params) => {
  return service({
    url: '/commentData/findCommentInfo',
    method: 'get',
    params
  })
}

// @Tags CommentInfo
// @Summary 分页获取评论信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commentData/getCommentInfoList [get]
export const getCommentInfoList = (params) => {
  return service({
    url: '/commentData/getCommentInfoList',
    method: 'get',
    params
  })
}
