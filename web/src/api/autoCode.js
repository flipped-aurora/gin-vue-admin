import service from '@/utils/request'

export const preview = (data) => {
  return service({
    url: '/autoCode/preview',
    method: 'post',
    data
  })
}

export const createTemp = (data) => {
  return service({
    url: '/autoCode/createTemp',
    method: 'post',
    data
  })
}

// @Tags SysApi
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDatabase [get]
export const getDB = (params) => {
  return service({
    url: '/autoCode/getDB',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getTables [get]
export const getTable = (params) => {
  return service({
    url: '/autoCode/getTables',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getColumn [get]
export const getColumn = (params) => {
  return service({
    url: '/autoCode/getColumn',
    method: 'get',
    params
  })
}

export const getSysHistory = (data) => {
  return service({
    url: '/autoCode/getSysHistory',
    method: 'post',
    data
  })
}

export const rollback = (data) => {
  return service({
    url: '/autoCode/rollback',
    method: 'post',
    data
  })
}

export const getMeta = (data) => {
  return service({
    url: '/autoCode/getMeta',
    method: 'post',
    data
  })
}

export const delSysHistory = (data) => {
  return service({
    url: '/autoCode/delSysHistory',
    method: 'post',
    data
  })
}

export const createPackageApi = (data) => {
  return service({
    url: '/autoCode/createPackage',
    method: 'post',
    data
  })
}

export const getPackageApi = () => {
  return service({
    url: '/autoCode/getPackage',
    method: 'post'
  })
}

export const deletePackageApi = (data) => {
  return service({
    url: '/autoCode/delPackage',
    method: 'post',
    data
  })
}

export const getTemplatesApi = () => {
  return service({
    url: '/autoCode/getTemplates',
    method: 'get'
  })
}

export const installPlug = (data) => {
  return service({
    url: '/autoCode/installPlug',
    method: 'post',
    data
  })
}

export const pubPlug = (params) => {
  return service({
    url: '/autoCode/pubPlug',
    method: 'post',
    params
  })
}

export const llmAuto = (data) => {
  return service({
    url: '/autoCode/llmAuto',
    method: 'post',
    data: { ...data, mode: 'ai' },
    timeout: 1000 * 60 * 10,
    loadingOption: {
      lock: true,
      fullscreen: true,
      text: `小淼正在思考，请稍候...`
    }
  })
}

export const butler = (data) => {
  return service({
    url: '/autoCode/llmAuto',
    method: 'post',
    data: { ...data, mode: 'butler' },
    timeout: 1000 * 60 * 10
  })
}

export const eye = (data) => {
  return service({
    url: '/autoCode/llmAuto',
    method: 'post',
    data: { ...data, mode: 'eye' },
    timeout: 1000 * 60 * 10
  })
}


export const createWeb = (data) => {
  return service({
    url: '/autoCode/llmAuto',
    method: 'post',
    data: { ...data, mode: 'painter' },
    timeout: 1000 * 60 * 10
  })
}



export const addFunc = (data) => {
  return service({
    url: '/autoCode/addFunc',
    method: 'post',
    data
  })
}

export const initMenu = (data) => {
  return service({
    url: '/autoCode/initMenu',
    method: 'post',
    data
  })
}

export const initAPI = (data) => {
  return service({
    url: '/autoCode/initAPI',
    method: 'post',
    data
  })
}

export const mcp = (data) => {
  return service({
    url: '/autoCode/mcp',
    method: 'post',
    data
  })
}


export const mcpList = (data) => {
  return service({
    url: '/autoCode/mcpList',
    method: 'post',
    data
  })
}


export const mcpTest = (data) => {
  return service({
    url: '/autoCode/mcpTest',
    method: 'post',
    data
  })
}
