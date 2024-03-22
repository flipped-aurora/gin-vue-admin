import service from '@/utils/request'

// @Tags Recruit
// @Summary 创建Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recruit true "创建Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recruit/createRecruit [post]
export const createRecruit = (data) => {
  return service({
    url: '/recruit/createRecruit',
    method: 'post',
    data
  })
}

// @Tags Recruit
// @Summary 删除Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recruit true "删除Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recruit/deleteRecruit [delete]
export const deleteRecruit = (data) => {
  return service({
    url: '/recruit/deleteRecruit',
    method: 'delete',
    data
  })
}

// @Tags Recruit
// @Summary 删除Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recruit/deleteRecruit [delete]
export const deleteRecruitByIds = (data) => {
  return service({
    url: '/recruit/deleteRecruitByIds',
    method: 'delete',
    data
  })
}

// @Tags Recruit
// @Summary 更新Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recruit true "更新Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recruit/updateRecruit [put]
export const updateRecruit = (data) => {
  return service({
    url: '/recruit/updateRecruit',
    method: 'put',
    data
  })
}

// @Tags Recruit
// @Summary 用id查询Recruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Recruit true "用id查询Recruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recruit/findRecruit [get]
export const findRecruit = (params) => {
  return service({
    url: '/recruit/findRecruit',
    method: 'get',
    params
  })
}

// @Tags Recruit
// @Summary 分页获取Recruit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Recruit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recruit/getRecruitList [get]
export const getRecruitList = (params) => {
  return service({
    url: '/recruit/getRecruitList',
    method: 'get',
    params
  })
}
