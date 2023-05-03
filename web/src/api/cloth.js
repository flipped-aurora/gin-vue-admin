import service from '@/utils/request'

// @Tags Cloth
// @Summary 创建Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cloth true "创建Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloth/createCloth [post]
export const createCloth = (data) => {
  return service({
    url: '/cloth/createCloth',
    method: 'post',
    data
  })
}

// @Tags Cloth
// @Summary 删除Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cloth true "删除Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloth/deleteCloth [delete]
export const deleteCloth = (data) => {
  return service({
    url: '/cloth/deleteCloth',
    method: 'delete',
    data
  })
}

// @Tags Cloth
// @Summary 删除Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloth/deleteCloth [delete]
export const deleteClothByIds = (data) => {
  return service({
    url: '/cloth/deleteClothByIds',
    method: 'delete',
    data
  })
}

// @Tags Cloth
// @Summary 更新Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cloth true "更新Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cloth/updateCloth [put]
export const updateCloth = (data) => {
  return service({
    url: '/cloth/updateCloth',
    method: 'put',
    data
  })
}

// @Tags Cloth
// @Summary 用id查询Cloth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Cloth true "用id查询Cloth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cloth/findCloth [get]
export const findCloth = (params) => {
  return service({
    url: '/cloth/findCloth',
    method: 'get',
    params
  })
}

// @Tags Cloth
// @Summary 分页获取Cloth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Cloth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloth/getClothList [get]
export const getClothList = (params) => {
  return service({
    url: '/cloth/getClothList',
    method: 'get',
    params
  })
}
