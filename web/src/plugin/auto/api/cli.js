import service from '@/utils/request'

export const createCli = (data) => {
  return service({
    url: '/cli/createCli',
    method: 'post',
    data
  })
}

export const getCliList = (data) => {
  return service({
    url: '/cli/getCliList',
    method: 'post',
    data
  })
}

export const getCliDetail = (data) => {
  return service({
    url: '/cli/getCliDetail',
    method: 'post',
    data
  })
}

export const updateCli = (data) => {
  return service({
    url: '/cli/updateCli',
    method: 'post',
    data
  })
}

export const deleteCli = (data) => {
  return service({
    url: '/cli/deleteCli',
    method: 'post',
    data
  })
}

export const addCliApis = (data) => {
  return service({
    url: '/cli/addCliApis',
    method: 'post',
    data
  })
}

export const removeCliApis = (data) => {
  return service({
    url: '/cli/removeCliApis',
    method: 'post',
    data
  })
}

export const previewManifest = (data) => {
  return service({
    url: '/cli/previewManifest',
    method: 'post',
    data
  })
}

export const previewApiCommand = (data) => {
  return service({
    url: '/cli/previewApiCommand',
    method: 'post',
    data
  })
}

export const downloadManifest = (data) => {
  return service({
    url: '/cli/downloadManifest',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

export const buildCliBinary = (data) => {
  return service({
    url: '/cli/buildCli',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

export const downloadCliSkill = (data) => {
  return service({
    url: '/cli/downloadSkill',
    method: 'post',
    data,
    responseType: 'blob'
  })
}
