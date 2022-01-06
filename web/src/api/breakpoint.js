import { createService } from '@/utils/request'
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]

export const findFile = (params) => {
  const service = createService()
  return service({
    url: '/fileUploadAndDownload/findFile',
    method: 'get',
    params
  })
}

export const breakpointContinue = (data) => {
  const service = createService()
  return service({
    url: '/fileUploadAndDownload/breakpointContinue',
    method: 'post',
    headers: { 'Content-Type': 'multipart/form-data' },
    data
  })
}

export const breakpointContinueFinish = (params) => {
  const service = createService()
  return service({
    url: '/fileUploadAndDownload/breakpointContinueFinish',
    method: 'post',
    params
  })
}

export const removeChunk = (data, params) => {
  const service = createService()
  return service({
    url: '/fileUploadAndDownload/removeChunk',
    method: 'post',
    data,
    params
  })
}
