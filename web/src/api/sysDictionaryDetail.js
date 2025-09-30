import service from '@/utils/request'
// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "创建SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
export const createSysDictionaryDetail = (data) => {
  return service({
    url: '/sysDictionaryDetail/createSysDictionaryDetail',
    method: 'post',
    data
  })
}

// @Tags SysDictionaryDetail
// @Summary 删除SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "删除SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
export const deleteSysDictionaryDetail = (data) => {
  return service({
    url: '/sysDictionaryDetail/deleteSysDictionaryDetail',
    method: 'delete',
    data
  })
}

// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
export const updateSysDictionaryDetail = (data) => {
  return service({
    url: '/sysDictionaryDetail/updateSysDictionaryDetail',
    method: 'put',
    data
  })
}

// @Tags SysDictionaryDetail
// @Summary 用id查询SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "用id查询SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
export const findSysDictionaryDetail = (params) => {
  return service({
    url: '/sysDictionaryDetail/findSysDictionaryDetail',
    method: 'get',
    params
  })
}

// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysDictionaryDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
export const getSysDictionaryDetailList = (params) => {
  return service({
    url: '/sysDictionaryDetail/getSysDictionaryDetailList',
    method: 'get',
    params
  })
}

// @Tags SysDictionaryDetail
// @Summary 获取层级字典详情树形结构（根据字典ID）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param sysDictionaryID query string true "字典ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getDictionaryTreeList [get]
export const getDictionaryTreeList = (params) => {
  return service({
    url: '/sysDictionaryDetail/getDictionaryTreeList',
    method: 'get',
    params
  })
}

// @Tags SysDictionaryDetail
// @Summary 获取层级字典详情树形结构（根据字典类型）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param dictType query string true "字典类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getDictionaryTreeListByType [get]
export const getDictionaryTreeListByType = (params) => {
  return service({
    url: '/sysDictionaryDetail/getDictionaryTreeListByType',
    method: 'get',
    params
  })
}

// @Tags SysDictionaryDetail
// @Summary 根据父级ID获取字典详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param parentID query string true "父级ID"
// @Param includeChildren query boolean false "是否包含子级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getDictionaryDetailsByParent [get]
export const getDictionaryDetailsByParent = (params) => {
  return service({
    url: '/sysDictionaryDetail/getDictionaryDetailsByParent',
    method: 'get',
    params
  })
}

// @Tags SysDictionaryDetail
// @Summary 获取字典详情的完整路径
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "字典详情ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getDictionaryPath [get]
export const getDictionaryPath = (params) => {
  return service({
    url: '/sysDictionaryDetail/getDictionaryPath',
    method: 'get',
    params
  })
}
