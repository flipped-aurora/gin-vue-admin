import service from '@/utils/request'

// @Tags Links
// @Summary 创建Links
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Links true "创建Links"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /links/createLinks [post]
export const createLinks = (data) => {
  return service({
    url: '/links/createLinks',
    method: 'post',
    data
  })
}

// @Tags Links
// @Summary 删除Links
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Links true "删除Links"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /links/deleteLinks [delete]
export const deleteLinks = (data) => {
  return service({
    url: '/links/deleteLinks',
    method: 'delete',
    data
  })
}

// @Tags Links
// @Summary 删除Links
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Links"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /links/deleteLinks [delete]
export const deleteLinksByIds = (data) => {
  return service({
    url: '/links/deleteLinksByIds',
    method: 'delete',
    data
  })
}

// @Tags Links
// @Summary 更新Links
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Links true "更新Links"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /links/updateLinks [put]
export const updateLinks = (data) => {
  return service({
    url: '/links/updateLinks',
    method: 'put',
    data
  })
}

// @Tags Links
// @Summary 用id查询Links
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Links true "用id查询Links"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /links/findLinks [get]
export const findLinks = (params) => {
  return service({
    url: '/links/findLinks',
    method: 'get',
    params
  })
}

// @Tags Links
// @Summary 分页获取Links列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Links列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /links/getLinksList [get]
export const getLinksList = (params) => {
  return service({
    url: '/links/getLinksList',
    method: 'get',
    params
  })
}
