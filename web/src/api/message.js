import service from '@/utils/request'

// @Tags Message
// @Summary 创建Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "创建Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /message/createMessage [post]
export const createMessage = (data) => {
  return service({
    url: '/message/createMessage',
    method: 'post',
    data
  })
}

// @Tags Message
// @Summary 删除Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "删除Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /message/deleteMessage [delete]
export const deleteMessage = (data) => {
  return service({
    url: '/message/deleteMessage',
    method: 'delete',
    data
  })
}

// @Tags Message
// @Summary 删除Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /message/deleteMessage [delete]
export const deleteMessageByIds = (data) => {
  return service({
    url: '/message/deleteMessageByIds',
    method: 'delete',
    data
  })
}

// @Tags Message
// @Summary 更新Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Message true "更新Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /message/updateMessage [put]
export const updateMessage = (data) => {
  return service({
    url: '/message/updateMessage',
    method: 'put',
    data
  })
}

// @Tags Message
// @Summary 用id查询Message
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Message true "用id查询Message"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /message/findMessage [get]
export const findMessage = (params) => {
  return service({
    url: '/message/findMessage',
    method: 'get',
    params
  })
}

// @Tags Message
// @Summary 分页获取Message列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Message列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /message/getMessageList [get]
export const getMessageList = (params) => {
  return service({
    url: '/message/getMessageList',
    method: 'get',
    params
  })
}
