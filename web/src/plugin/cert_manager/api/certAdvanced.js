import service from '@/utils/request'

export const discoverSubdomains = (data) => {
  return service({
    url: '/cert/advanced/discoverSubdomains',
    method: 'post',
    data: data
  })
}

export const batchDiscoverSubdomains = (data) => {
  return service({
    url: '/cert/advanced/batchDiscoverSubdomains',
    method: 'post',
    data: data
  })
}

export const getDomainTree = (params) => {
  return service({
    url: '/cert/advanced/domainTree',
    method: 'get',
    params: params
  })
}

export const reProbeDomainTree = (data) => {
  return service({
    url: '/cert/advanced/reProbeDomainTree',
    method: 'post',
    data: data
  })
}

export const ignoreDomain = (data) => {
  return service({
    url: '/cert/advanced/ignoreDomain',
    method: 'post',
    data: data
  })
}

export const batchReprobe = (data) => {
  return service({
    url: '/cert/advanced/batchReprobe',
    method: 'post',
    data: data
  })
}

export const batchIgnore = (data) => {
  return service({
    url: '/cert/advanced/batchIgnore',
    method: 'post',
    data: data
  })
}

export const getDomainAssetList = (params) => {
  return service({
    url: '/cert/advanced/domainAssetList',
    method: 'get',
    params: params
  })
}

export const exportSubdomainReport = (params) => {
  return service({
    url: '/cert/advanced/exportSubdomainReport',
    method: 'get',
    params: params,
    responseType: 'blob'
  })
}
