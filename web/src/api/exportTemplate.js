import service from '@/utils/request'

// @Tags SysExportTemplate
// @Summary 创建导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysExportTemplate true "创建导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysExportTemplate/createSysExportTemplate [post]
export const createSysExportTemplate = (data) => {
  return service({
    url: '/sysExportTemplate/createSysExportTemplate',
    method: 'post',
    data
  })
}

// @Tags SysExportTemplate
// @Summary 删除导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysExportTemplate true "删除导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
export const deleteSysExportTemplate = (data) => {
  return service({
    url: '/sysExportTemplate/deleteSysExportTemplate',
    method: 'delete',
    data
  })
}

// @Tags SysExportTemplate
// @Summary 批量删除导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
export const deleteSysExportTemplateByIds = (data) => {
  return service({
    url: '/sysExportTemplate/deleteSysExportTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags SysExportTemplate
// @Summary 更新导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysExportTemplate true "更新导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysExportTemplate/updateSysExportTemplate [put]
export const updateSysExportTemplate = (data) => {
  return service({
    url: '/sysExportTemplate/updateSysExportTemplate',
    method: 'put',
    data
  })
}

// @Tags SysExportTemplate
// @Summary 用id查询导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysExportTemplate true "用id查询导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysExportTemplate/findSysExportTemplate [get]
export const findSysExportTemplate = (params) => {
  return service({
    url: '/sysExportTemplate/findSysExportTemplate',
    method: 'get',
    params
  })
}

// @Tags SysExportTemplate
// @Summary 分页获取导出模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取导出模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysExportTemplate/getSysExportTemplateList [get]
export const getSysExportTemplateList = (params) => {
  return service({
    url: '/sysExportTemplate/getSysExportTemplateList',
    method: 'get',
    params
  })
}


// ExportExcel 导出表格token
// @Tags SysExportTemplate
// @Summary 导出表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportExcel [get]
export const exportExcel = (params) => {
  return service({
    url: '/sysExportTemplate/exportExcel',
    method: 'get',
    params
  })
}

// ExportTemplate 导出表格模板
// @Tags SysExportTemplate
// @Summary 导出表格模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportTemplate [get]
export const exportTemplate = (params) => {
  return service({
    url: '/sysExportTemplate/exportTemplate',
    method: 'get',
    params
  })
}

// PreviewSQL 预览最终生成的SQL
// @Tags SysExportTemplate
// @Summary 预览最终生成的SQL
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/previewSQL [get]
// @Param templateID query string true  "导出模板ID"
// @Param params     query string false "查询参数编码字符串，参考 ExportExcel 组件"
export const previewSQL = (params) => {
  return service({
    url: '/sysExportTemplate/previewSQL',
    method: 'get',
    params
  })
}
