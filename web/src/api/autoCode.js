import service from '@/utils/request'

const DEFAULT_LLM_TIMEOUT = 1000 * 60 * 10

const LLM_AUTO_URL = '/autoCode/llmAuto'

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

export const getDB = (params) => {
  return service({
    url: '/autoCode/getDB',
    method: 'get',
    params
  })
}

export const getTable = (params) => {
  return service({
    url: '/autoCode/getTables',
    method: 'get',
    params
  })
}

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

export const llmAuto = (data, options = {}) => {
  return service({
    url: LLM_AUTO_URL,
    method: 'post',
    data: { ...data },
    timeout: options.timeout ?? DEFAULT_LLM_TIMEOUT,
    loadingOption: {
      lock: true,
      fullscreen: true,
      persistLoading: true,
      text: '小淼正在思考，请稍候...'
    },
    donNotShowLoading: options.donNotShowLoading ?? false
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

export const initDictionary = (data) => {
  return service({
    url: '/autoCode/initDictionary',
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

export const mcpStatus = () => {
  return service({
    url: '/autoCode/mcpStatus',
    method: 'post'
  })
}

export const mcpStart = () => {
  return service({
    url: '/autoCode/mcpStart',
    method: 'post'
  })
}

export const mcpStop = () => {
  return service({
    url: '/autoCode/mcpStop',
    method: 'post'
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

export const getPluginList = (params) => {
  return service({
    url: '/autoCode/getPluginList',
    method: 'get',
    params
  })
}

export const removePlugin = (params) => {
  return service({
    url: '/autoCode/removePlugin',
    method: 'post',
    params
  })
}
