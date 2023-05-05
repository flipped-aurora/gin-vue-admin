import service from '@/utils/request'

// @Tags CompanyApply
// @Summary 创建CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyApply true "创建CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyApply/createCompanyApply [post]
export const createCompanyApply = (data) => {
  return service({
    url: '/companyApply/createCompanyApply',
    method: 'post',
    data
  })
}

// @Tags CompanyApply
// @Summary 删除CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyApply true "删除CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyApply/deleteCompanyApply [delete]
export const deleteCompanyApply = (data) => {
  return service({
    url: '/companyApply/deleteCompanyApply',
    method: 'delete',
    data
  })
}

// @Tags CompanyApply
// @Summary 删除CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyApply/deleteCompanyApply [delete]
export const deleteCompanyApplyByIds = (data) => {
  return service({
    url: '/companyApply/deleteCompanyApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags CompanyApply
// @Summary 更新CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyApply true "更新CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /companyApply/updateCompanyApply [put]
export const updateCompanyApply = (data) => {
  return service({
    url: '/companyApply/updateCompanyApply',
    method: 'put',
    data
  })
}

// @Tags CompanyApply
// @Summary 用id查询CompanyApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CompanyApply true "用id查询CompanyApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /companyApply/findCompanyApply [get]
export const findCompanyApply = (params) => {
  return service({
    url: '/companyApply/findCompanyApply',
    method: 'get',
    params
  })
}

// @Tags CompanyApply
// @Summary 分页获取CompanyApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CompanyApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyApply/getCompanyApplyList [get]
export const getCompanyApplyList = (params) => {
  return service({
    url: '/companyApply/getCompanyApplyList',
    method: 'get',
    params
  })
}
