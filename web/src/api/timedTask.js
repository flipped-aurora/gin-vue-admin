import service from '@/utils/request'

export const createTimedTask = (data) => {
  return service({ url: '/timedTask/createTimedTask', method: 'post', data })
}

export const updateTimedTask = (data) => {
  return service({ url: '/timedTask/updateTimedTask', method: 'put', data })
}

export const deleteTimedTask = (data) => {
  return service({ url: '/timedTask/deleteTimedTask', method: 'delete', data })
}

export const toggleTimedTask = (data) => {
  return service({ url: '/timedTask/toggleTimedTask', method: 'post', data })
}

export const triggerTimedTask = (data) => {
  return service({ url: '/timedTask/triggerTimedTask', method: 'post', data })
}

export const getTimedTaskList = (params) => {
  return service({ url: '/timedTask/getTimedTaskList', method: 'get', params })
}

export const getTimedTaskLogList = (params) => {
  return service({ url: '/timedTask/getTimedTaskLogList', method: 'get', params })
}

export const getRegisteredMethods = () => {
  return service({ url: '/timedTask/getRegisteredMethods', method: 'get' })
}
