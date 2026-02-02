import service from '@/utils/request'

export const createCertCategory = (data) => {
  return service({
    url: '/certCategory/createCertCategory',
    method: 'post',
    data
  })
}

export const deleteCertCategory = (data) => {
  return service({
    url: '/certCategory/deleteCertCategory',
    method: 'delete',
    data
  })
}

export const updateCertCategory = (data) => {
  return service({
    url: '/certCategory/updateCertCategory',
    method: 'put',
    data
  })
}

export const getCertCategoryList = (params) => {
  return service({
    url: '/certCategory/getCertCategoryList',
    method: 'get',
    params
  })
}

export const batchUpdateDomainCategory = (data) => {
  return service({
    url: '/certCategory/batchUpdateDomainCategory',
    method: 'post',
    data
  })
}
