import service from '@/utils/request'

// @Tags Style
// @Summary 创建Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Style true "创建Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /style/createStyle [post]
export const createStyle = (data) => {
  return service({
    url: '/style/createStyle',
    method: 'post',
    data
  })
}

// @Tags Style
// @Summary 删除Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Style true "删除Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /style/deleteStyle [delete]
export const deleteStyle = (data) => {
  return service({
    url: '/style/deleteStyle',
    method: 'delete',
    data
  })
}

// @Tags Style
// @Summary 删除Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /style/deleteStyle [delete]
export const deleteStyleByIds = (data) => {
  return service({
    url: '/style/deleteStyleByIds',
    method: 'delete',
    data
  })
}

// @Tags Style
// @Summary 更新Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Style true "更新Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /style/updateStyle [put]
export const updateStyle = (data) => {
  return service({
    url: '/style/updateStyle',
    method: 'put',
    data
  })
}

// @Tags Style
// @Summary 用id查询Style
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Style true "用id查询Style"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /style/findStyle [get]
export const findStyle = (params) => {
  return service({
    url: '/style/findStyle',
    method: 'get',
    params
  })
}

// @Tags Style
// @Summary 分页获取Style列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Style列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /style/getStyleList [get]
export const getStyleList = (params) => {
  return service({
    url: '/style/getStyleList',
    method: 'get',
    params
  })
}
