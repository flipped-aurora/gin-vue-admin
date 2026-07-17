import service from '@/utils/request'

// @Router /department/getDepartmentList [post]
export const getDepartmentList = (data) => {
  return service({
    url: '/department/getDepartmentList',
    method: 'post',
    data
  })
}

// @Router /department/createDepartment [post]
export const createDepartment = (data) => {
  return service({
    url: '/department/createDepartment',
    method: 'post',
    data
  })
}

// @Router /department/updateDepartment [put]
export const updateDepartment = (data) => {
  return service({
    url: '/department/updateDepartment',
    method: 'put',
    data
  })
}

// @Router /department/deleteDepartment [delete]
export const deleteDepartment = (data) => {
  return service({
    url: '/department/deleteDepartment',
    method: 'delete',
    data
  })
}

// @Router /department/findDepartment [get]
export const findDepartment = (params) => {
  return service({
    url: '/department/findDepartment',
    method: 'get',
    params
  })
}

// 获取部门成员ID列表 @Router /department/getDepartmentUsers [get]
export const getDepartmentUsers = (params) => {
  return service({
    url: '/department/getDepartmentUsers',
    method: 'get',
    params
  })
}

// 设置部门成员(反向分配) @Router /department/setDepartmentUsers [post]
export const setDepartmentUsers = (data) => {
  return service({
    url: '/department/setDepartmentUsers',
    method: 'post',
    data
  })
}
