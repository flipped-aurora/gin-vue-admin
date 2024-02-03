import service from '@/utils/request'

// @Tags DisInfo
// @Summary 创建帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DisInfo true "创建帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /disData/createDisInfo [post]
export const createDisInfo = (data) => {
  return service({
    url: '/disData/createDisInfo',
    method: 'post',
    data
  })
}

// @Tags DisInfo
// @Summary 删除帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DisInfo true "删除帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disData/deleteDisInfo [delete]
export const deleteDisInfo = (params) => {
  return service({
    url: '/disData/deleteDisInfo',
    method: 'delete',
    params
  })
}

// @Tags DisInfo
// @Summary 批量删除帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disData/deleteDisInfo [delete]
export const deleteDisInfoByIds = (params) => {
  return service({
    url: '/disData/deleteDisInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags DisInfo
// @Summary 更新帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DisInfo true "更新帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /disData/updateDisInfo [put]
export const updateDisInfo = (data) => {
  return service({
    url: '/disData/updateDisInfo',
    method: 'put',
    data
  })
}

// @Tags DisInfo
// @Summary 用id查询帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DisInfo true "用id查询帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /disData/findDisInfo [get]
export const findDisInfo = (params) => {
  return service({
    url: '/disData/findDisInfo',
    method: 'get',
    params
  })
}

// @Tags DisInfo
// @Summary 分页获取帖子信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取帖子信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /disData/getDisInfoList [get]
export const getDisInfoList = (params) => {
  return service({
    url: '/disData/getDisInfoList',
    method: 'get',
    params
  })
}
