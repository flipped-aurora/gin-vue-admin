import service from '@/utils/request'

// @Tags Company
// @Summary 创建Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "创建Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/createCompany [post]
export const createCompany = (data) => {
  return service({
    url: '/company/createCompany',
    method: 'post',
    data
  })
}

// @Tags Company
// @Summary 删除Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "删除Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
export const deleteCompany = (data) => {
  return service({
    url: '/company/deleteCompany',
    method: 'delete',
    data
  })
}

// @Tags Company
// @Summary 删除Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
export const deleteCompanyByIds = (data) => {
  return service({
    url: '/company/deleteCompanyByIds',
    method: 'delete',
    data
  })
}

// @Tags Company
// @Summary 更新Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "更新Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /company/updateCompany [put]
export const updateCompany = (data) => {
  return service({
    url: '/company/updateCompany',
    method: 'put',
    data
  })
}

// @Tags Company
// @Summary 用id查询Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Company true "用id查询Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /company/findCompany [get]
export const findCompany = (params) => {
  return service({
    url: '/company/findCompany',
    method: 'get',
    params
  })
}

// @Tags Company
// @Summary 分页获取Company列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Company列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/getCompanyList [get]
export const getCompanyList = (params) => {
  return service({
    url: '/company/getCompanyList',
    method: 'get',
    params
  })
}
