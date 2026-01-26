import service from '@/utils/request'

export const deleteLoginLog = (data) => {
  return service({
    url: '/sysLoginLog/deleteLoginLog',
    method: 'delete',
    data
  })
}

export const deleteLoginLogByIds = (data) => {
  return service({
    url: '/sysLoginLog/deleteLoginLogByIds',
    method: 'delete',
    data
  })
}

export const getLoginLogList = (params) => {
  return service({
    url: '/sysLoginLog/getLoginLogList',
    method: 'get',
    params
  })
}

export const findLoginLog = (params) => {
    return service({
        url: '/sysLoginLog/findLoginLog',
        method: 'get',
        params
    })
}
