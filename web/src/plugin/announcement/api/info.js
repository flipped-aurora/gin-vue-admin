import service from '@/utils/request'

// @Tags Info
// @Summary 创建公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "创建公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /info/createInfo [post]
export const createInfo = (data) => {
  return service({
    url: '/info/createInfo',
    method: 'post',
    data
  })
}

// @Tags Info
// @Summary 删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "删除公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /info/deleteInfo [delete]
export const deleteInfo = (params) => {
  return service({
    url: '/info/deleteInfo',
    method: 'delete',
    params
  })
}

// @Tags Info
// @Summary 批量删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /info/deleteInfo [delete]
export const deleteInfoByIds = (params) => {
  return service({
    url: '/info/deleteInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags Info
// @Summary 更新公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "更新公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /info/updateInfo [put]
export const updateInfo = (data) => {
  return service({
    url: '/info/updateInfo',
    method: 'put',
    data
  })
}

// @Tags Info
// @Summary 用id查询公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Info true "用id查询公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /info/findInfo [get]
export const findInfo = (params) => {
  return service({
    url: '/info/findInfo',
    method: 'get',
    params
  })
}

// @Tags Info
// @Summary 分页获取公告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/getInfoList [get]
export const getInfoList = (params) => {
  return service({
    url: '/info/getInfoList',
    method: 'get',
    params
  })
}
// @Tags Info
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /info/findInfoDataSource [get]
export const getInfoDataSource = () => {
  return service({
    url: '/info/getInfoDataSource',
    method: 'get'
  })
}
