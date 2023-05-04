import service from '@/utils/request'

// @Tags MsgBox
// @Summary 创建MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgBox true "创建MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgBox/createMsgBox [post]
export const createMsgBox = (data) => {
  return service({
    url: '/msgBox/createMsgBox',
    method: 'post',
    data
  })
}

// @Tags MsgBox
// @Summary 删除MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgBox true "删除MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgBox/deleteMsgBox [delete]
export const deleteMsgBox = (data) => {
  return service({
    url: '/msgBox/deleteMsgBox',
    method: 'delete',
    data
  })
}

// @Tags MsgBox
// @Summary 删除MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgBox/deleteMsgBox [delete]
export const deleteMsgBoxByIds = (data) => {
  return service({
    url: '/msgBox/deleteMsgBoxByIds',
    method: 'delete',
    data
  })
}

// @Tags MsgBox
// @Summary 更新MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgBox true "更新MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msgBox/updateMsgBox [put]
export const updateMsgBox = (data) => {
  return service({
    url: '/msgBox/updateMsgBox',
    method: 'put',
    data
  })
}

// @Tags MsgBox
// @Summary 用id查询MsgBox
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsgBox true "用id查询MsgBox"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msgBox/findMsgBox [get]
export const findMsgBox = (params) => {
  return service({
    url: '/msgBox/findMsgBox',
    method: 'get',
    params
  })
}

// @Tags MsgBox
// @Summary 分页获取MsgBox列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MsgBox列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgBox/getMsgBoxList [get]
export const getMsgBoxList = (params) => {
  return service({
    url: '/msgBox/getMsgBoxList',
    method: 'get',
    params
  })
}
