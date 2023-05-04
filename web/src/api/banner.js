import service from '@/utils/request'

// @Tags Banner
// @Summary 创建Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "创建Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/createBanner [post]
export const createBanner = (data) => {
  return service({
    url: '/banner/createBanner',
    method: 'post',
    data
  })
}

// @Tags Banner
// @Summary 删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [delete]
export const deleteBanner = (data) => {
  return service({
    url: '/banner/deleteBanner',
    method: 'delete',
    data
  })
}

// @Tags Banner
// @Summary 删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [delete]
export const deleteBannerByIds = (data) => {
  return service({
    url: '/banner/deleteBannerByIds',
    method: 'delete',
    data
  })
}

// @Tags Banner
// @Summary 更新Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "更新Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /banner/updateBanner [put]
export const updateBanner = (data) => {
  return service({
    url: '/banner/updateBanner',
    method: 'put',
    data
  })
}

// @Tags Banner
// @Summary 用id查询Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Banner true "用id查询Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /banner/findBanner [get]
export const findBanner = (params) => {
  return service({
    url: '/banner/findBanner',
    method: 'get',
    params
  })
}

// @Tags Banner
// @Summary 分页获取Banner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Banner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/getBannerList [get]
export const getBannerList = (params) => {
  return service({
    url: '/banner/getBannerList',
    method: 'get',
    params
  })
}
