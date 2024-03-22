import service from '@/utils/request'

// @Tags CateMenus
// @Summary 创建CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CateMenus true "创建CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cateMenus/createCateMenus [post]
export const createCateMenus = (data) => {
  return service({
    url: '/cateMenus/createCateMenus',
    method: 'post',
    data
  })
}

// @Tags CateMenus
// @Summary 删除CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CateMenus true "删除CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cateMenus/deleteCateMenus [delete]
export const deleteCateMenus = (data) => {
  return service({
    url: '/cateMenus/deleteCateMenus',
    method: 'delete',
    data
  })
}

// @Tags CateMenus
// @Summary 删除CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cateMenus/deleteCateMenus [delete]
export const deleteCateMenusByIds = (data) => {
  return service({
    url: '/cateMenus/deleteCateMenusByIds',
    method: 'delete',
    data
  })
}

// @Tags CateMenus
// @Summary 更新CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CateMenus true "更新CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cateMenus/updateCateMenus [put]
export const updateCateMenus = (data) => {
  return service({
    url: '/cateMenus/updateCateMenus',
    method: 'put',
    data
  })
}

// @Tags CateMenus
// @Summary 用id查询CateMenus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CateMenus true "用id查询CateMenus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cateMenus/findCateMenus [get]
export const findCateMenus = (params) => {
  return service({
    url: '/cateMenus/findCateMenus',
    method: 'get',
    params
  })
}

// @Tags CateMenus
// @Summary 分页获取CateMenus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CateMenus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cateMenus/getCateMenusList [get]
export const getCateMenusList = (params) => {
  return service({
    url: '/cateMenus/getCateMenusList',
    method: 'get',
    params
  })
}

// 获取模板名称列表
export const getTemplateList = () => {
  return service({
    url: '/cateMenus/getTemplateList',
    method: 'get',
  })
}

// 获取模型数据
export const getModelList = () => {
  return service({
    url: '/cateMenus/getModelsList',
    method: 'get',
  })
}
