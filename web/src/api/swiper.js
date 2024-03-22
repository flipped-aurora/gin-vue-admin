import service from '@/utils/request'

// @Tags Swiper
// @Summary 创建Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Swiper true "创建Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /swiper/createSwiper [post]
export const createSwiper = (data) => {
  return service({
    url: '/swiper/createSwiper',
    method: 'post',
    data
  })
}

// @Tags Swiper
// @Summary 删除Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Swiper true "删除Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /swiper/deleteSwiper [delete]
export const deleteSwiper = (data) => {
  return service({
    url: '/swiper/deleteSwiper',
    method: 'delete',
    data
  })
}

// @Tags Swiper
// @Summary 删除Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /swiper/deleteSwiper [delete]
export const deleteSwiperByIds = (data) => {
  return service({
    url: '/swiper/deleteSwiperByIds',
    method: 'delete',
    data
  })
}

// @Tags Swiper
// @Summary 更新Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Swiper true "更新Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /swiper/updateSwiper [put]
export const updateSwiper = (data) => {
  return service({
    url: '/swiper/updateSwiper',
    method: 'put',
    data
  })
}

// @Tags Swiper
// @Summary 用id查询Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Swiper true "用id查询Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /swiper/findSwiper [get]
export const findSwiper = (params) => {
  return service({
    url: '/swiper/findSwiper',
    method: 'get',
    params
  })
}

// @Tags Swiper
// @Summary 分页获取Swiper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Swiper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /swiper/getSwiperList [get]
export const getSwiperList = (params) => {
  return service({
    url: '/swiper/getSwiperList',
    method: 'get',
    params
  })
}
