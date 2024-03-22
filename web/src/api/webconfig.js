import service from '@/utils/request'

export const getWebconfig = () => {
    return service({
      url: '/getWebconfig',
      method: 'get'
    })
}

export const setWebconfig = (data) => {
    return service({
      url: '/setWebconfig',
      method: 'post',
      data: data
    })
}

export const createWebconfig = (data) => {
  return service({
    url: '/createWebconfig',
    method: 'post',
    data
  })
}

export const delWebconfig = (data) => {
  return service({
    url: '/delWebconfig',
    method: 'delete',
    data
  })
}

// 切换站点
export const changeWebconfig = (data) => {
  return service({
    url: '/changeWebconfig',
    method: 'post',
    data: data
  })
}


