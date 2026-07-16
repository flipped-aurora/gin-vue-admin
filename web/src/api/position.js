import service from '@/utils/request'

// @Router /position/getPositionList [post]
export const getPositionList = (data) => {
  return service({
    url: '/position/getPositionList',
    method: 'post',
    data
  })
}

// @Router /position/createPosition [post]
export const createPosition = (data) => {
  return service({
    url: '/position/createPosition',
    method: 'post',
    data
  })
}

// @Router /position/updatePosition [put]
export const updatePosition = (data) => {
  return service({
    url: '/position/updatePosition',
    method: 'put',
    data
  })
}

// @Router /position/deletePosition [delete]
export const deletePosition = (data) => {
  return service({
    url: '/position/deletePosition',
    method: 'delete',
    data
  })
}

// @Router /position/findPosition [get]
export const findPosition = (params) => {
  return service({
    url: '/position/findPosition',
    method: 'get',
    params
  })
}

// 获取岗位成员ID列表 @Router /position/getPositionUsers [get]
export const getPositionUsers = (params) => {
  return service({
    url: '/position/getPositionUsers',
    method: 'get',
    params
  })
}

// 设置岗位成员(反向分配) @Router /position/setPositionUsers [post]
export const setPositionUsers = (data) => {
  return service({
    url: '/position/setPositionUsers',
    method: 'post',
    data
  })
}
