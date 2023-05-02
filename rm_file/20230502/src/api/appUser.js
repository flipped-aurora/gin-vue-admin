import service from '@/utils/request'

// @Tags AppUser
// @Summary 创建AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUser true "创建AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/createAppUser [post]
export const createAppUser = (data) => {
  return service({
    url: '/appUser/createAppUser',
    method: 'post',
    data
  })
}

// @Tags AppUser
// @Summary 删除AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUser true "删除AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUser/deleteAppUser [delete]
export const deleteAppUser = (data) => {
  return service({
    url: '/appUser/deleteAppUser',
    method: 'delete',
    data
  })
}

// @Tags AppUser
// @Summary 删除AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUser/deleteAppUser [delete]
export const deleteAppUserByIds = (data) => {
  return service({
    url: '/appUser/deleteAppUserByIds',
    method: 'delete',
    data
  })
}

// @Tags AppUser
// @Summary 更新AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUser true "更新AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUser/updateAppUser [put]
export const updateAppUser = (data) => {
  return service({
    url: '/appUser/updateAppUser',
    method: 'put',
    data
  })
}

// @Tags AppUser
// @Summary 用id查询AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppUser true "用id查询AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUser/findAppUser [get]
export const findAppUser = (params) => {
  return service({
    url: '/appUser/findAppUser',
    method: 'get',
    params
  })
}

// @Tags AppUser
// @Summary 分页获取AppUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/getAppUserList [get]
export const getAppUserList = (params) => {
  return service({
    url: '/appUser/getAppUserList',
    method: 'get',
    params
  })
}
