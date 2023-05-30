import service from '@/utils/request'

// @Tags SizeList
// @Summary 创建SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SizeList true "创建SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sizeList/createSizeList [post]
export const createSizeList = (data) => {
  return service({
    url: '/sizeList/createSizeList',
    method: 'post',
    data
  })
}

// @Tags SizeList
// @Summary 删除SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SizeList true "删除SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sizeList/deleteSizeList [delete]
export const deleteSizeList = (data) => {
  return service({
    url: '/sizeList/deleteSizeList',
    method: 'delete',
    data
  })
}

// @Tags SizeList
// @Summary 删除SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sizeList/deleteSizeList [delete]
export const deleteSizeListByIds = (data) => {
  return service({
    url: '/sizeList/deleteSizeListByIds',
    method: 'delete',
    data
  })
}

// @Tags SizeList
// @Summary 更新SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SizeList true "更新SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sizeList/updateSizeList [put]
export const updateSizeList = (data) => {
  return service({
    url: '/sizeList/updateSizeList',
    method: 'put',
    data
  })
}

// @Tags SizeList
// @Summary 用id查询SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SizeList true "用id查询SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sizeList/findSizeList [get]
export const findSizeList = (params) => {
  return service({
    url: '/sizeList/findSizeList',
    method: 'get',
    params
  })
}

// @Tags SizeList
// @Summary 分页获取SizeList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SizeList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sizeList/getSizeListList [get]
export const getSizeListList = (params) => {
  return service({
    url: '/sizeList/getSizeListList',
    method: 'get',
    params
  })
}
