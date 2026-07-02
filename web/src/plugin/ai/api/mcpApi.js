import service from '@/utils/request'

export const createMcp = (data) => {
  return service({
    url: '/mcpApi/createMcp',
    method: 'post',
    data
  })
}

export const getMcpList = (data) => {
  return service({
    url: '/mcpApi/getMcpList',
    method: 'post',
    data
  })
}

export const getMcpDetail = (data) => {
  return service({
    url: '/mcpApi/getMcpDetail',
    method: 'post',
    data
  })
}

export const updateMcp = (data) => {
  return service({
    url: '/mcpApi/updateMcp',
    method: 'post',
    data
  })
}

export const deleteMcp = (data) => {
  return service({
    url: '/mcpApi/deleteMcp',
    method: 'post',
    data
  })
}

export const addMcpApis = (data) => {
  return service({
    url: '/mcpApi/addMcpApis',
    method: 'post',
    data
  })
}

export const removeMcpApis = (data) => {
  return service({
    url: '/mcpApi/removeMcpApis',
    method: 'post',
    data
  })
}

export const previewMcpManifest = (data) => {
  return service({
    url: '/mcpApi/previewManifest',
    method: 'post',
    data
  })
}

export const previewMcpPrompt = (data) => {
  return service({
    url: '/mcpApi/previewPrompt',
    method: 'post',
    data
  })
}

export const previewApiCommand = (data) => {
  return service({
    url: '/mcpApi/previewApiCommand',
    method: 'post',
    data
  })
}
