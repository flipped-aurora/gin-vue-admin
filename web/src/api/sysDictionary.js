import service from '@/utils/request'
// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "创建SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/createSysDictionary [post]
export const createSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/createSysDictionary',
    method: 'post',
    data
  })
}

// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "删除SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionary/deleteSysDictionary [delete]
export const deleteSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/deleteSysDictionary',
    method: 'delete',
    data
  })
}

// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "更新SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionary/updateSysDictionary [put]
export const updateSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/updateSysDictionary',
    method: 'put',
    data
  })
}

// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "用id查询SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionary/findSysDictionary [get]
export const findSysDictionary = (params) => {
  return service({
    url: '/sysDictionary/findSysDictionary',
    method: 'get',
    params
  })
}

// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysDictionary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/getSysDictionaryList [get]
export const getSysDictionaryList = (params) => {
  return service({
    url: '/sysDictionary/getSysDictionaryList',
    method: 'get',
    params
  })
}

// @Tags SysDictionary
// @Summary 导出字典JSON（包含字典详情）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysDictionary true "字典ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导出成功"}"
// @Router /sysDictionary/exportSysDictionary [get]
export const exportSysDictionary = (params) => {
  return service({
    url: '/sysDictionary/exportSysDictionary',
    method: 'get',
    params
  })
}

// @Tags SysDictionary
// @Summary 导入字典JSON（包含字典详情）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body object true "字典JSON数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /sysDictionary/importSysDictionary [post]
export const importSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/importSysDictionary',
    method: 'post',
    data
  })
}
