import service from '@/utils/request'

// @Tags UserProfile
// @Summary 创建UserProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Profile true "创建UserProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "创建UserProfile"
// @Router /profile/createUserProfile [post]
export const createUserProfile = (data) => {
  return service({
    url: '/profile/createUserProfile',
    method: 'post',
    data
  })
}

// @Tags UserProfile
// @Summary 删除UserProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Profile true "删除UserProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "删除UserProfile"
// @Router /profile/deleteUserProfile [delete]
export const deleteUserProfile = (data) => {
  return service({
    url: '/profile/deleteUserProfile',
    method: 'delete',
    data
  })
}

// @Tags UserProfile
// @Summary 批量删除UserProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "批量删除UserProfile"
// @Router /profile/deleteUserProfileByIds [delete]
export const deleteUserProfileByIds = (data) => {
  return service({
    url: '/profile/deleteUserProfileByIds',
    method: 'delete',
    data
  })
}

// @Tags UserProfile
// @Summary 更新UserProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Profile true "更新UserProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "更新UserProfile"
// @Router /profile/updateUserProfile [put]
export const updateUserProfile = (data) => {
  return service({
    url: '/profile/updateUserProfile',
    method: 'put',
    data
  })
}

// @Tags UserProfile
// @Summary 用id查询UserProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Profile true "用id查询UserProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询UserProfile"
// @Router /profile/findUserProfile [get]
export const findUserProfile = (params) => {
  return service({
    url: '/profile/findUserProfile',
    method: 'get',
    params
  })
}

// @Tags UserProfile
// @Summary 分页获取UserProfile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProfileSearch true "分页获取UserProfile列表"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "分页获取UserProfile列表"
// @Router /profile/getUserProfileList [get]
export const getUserProfileList = (params) => {
  return service({
    url: '/profile/getUserProfileList',
    method: 'get',
    params
  })
}

// @Tags CompanyProfile
// @Summary 创建CompanyProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Profile true "创建CompanyProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "创建CompanyProfile"
// @Router /profile/createCompanyProfile [post]
export const createCompanyProfile = (data) => {
  return service({
    url: '/profile/createCompanyProfile',
    method: 'post',
    data
  })
}

// @Tags CompanyProfile
// @Summary 删除CompanyProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Profile true "删除CompanyProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "删除CompanyProfile"
// @Router /profile/deleteCompanyProfile [delete]
export const deleteCompanyProfile = (data) => {
  return service({
    url: '/profile/deleteCompanyProfile',
    method: 'delete',
    data
  })
}

// @Tags CompanyProfile
// @Summary 批量删除CompanyProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "批量删除CompanyProfile"
// @Router /profile/deleteCompanyProfileByIds [delete]
export const deleteCompanyProfileByIds = (data) => {
  return service({
    url: '/profile/deleteCompanyProfileByIds',
    method: 'delete',
    data
  })
}

// @Tags CompanyProfile
// @Summary 更新CompanyProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Profile true "更新CompanyProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "更新CompanyProfile"
// @Router /profile/updateCompanyProfile [put]
export const updateCompanyProfile = (data) => {
  return service({
    url: '/profile/updateCompanyProfile',
    method: 'put',
    data
  })
}

// @Tags CompanyProfile
// @Summary 用id查询CompanyProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Profile true "用id查询CompanyProfile"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询CompanyProfile"
// @Router /profile/findCompanyProfile [get]
export const findCompanyProfile = (params) => {
  return service({
    url: '/profile/findCompanyProfile',
    method: 'get',
    params
  })
}

// @Tags CompanyProfile
// @Summary 分页获取CompanyProfile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProfileSearch true "分页获取CompanyProfile列表"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "分页获取CompanyProfile列表"
// @Router /profile/getCompanyProfileList [get]
export const getCompanyProfileList = (params) => {
  return service({
    url: '/profile/getCompanyProfileList',
    method: 'get',
    params
  })
}
