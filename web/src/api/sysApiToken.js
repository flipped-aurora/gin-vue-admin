import service from '@/utils/request'

export const createApiToken = (data) => {
  return service({
    url: '/sysApiToken/createApiToken',
    method: 'post',
    data
  })
}

export const getApiTokenList = (data) => {
  return service({
    url: '/sysApiToken/getApiTokenList',
    method: 'post',
    data
  })
}

export const deleteApiToken = (data) => {
  return service({
    url: '/sysApiToken/deleteApiToken',
    method: 'post',
    data
  })
}
