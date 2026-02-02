import service from '@/utils/request'

export const createCertCertificate = (data) => {
  return service({
    url: '/certCertificate/createCertCertificate',
    method: 'post',
    data
  })
}

export const deleteCertCertificate = (params) => {
  return service({
    url: '/certCertificate/deleteCertCertificate',
    method: 'delete',
    params
  })
}

export const deleteCertCertificateByIds = (params) => {
  return service({
    url: '/certCertificate/deleteCertCertificateByIds',
    method: 'delete',
    params
  })
}

export const updateCertCertificate = (data) => {
  return service({
    url: '/certCertificate/updateCertCertificate',
    method: 'put',
    data
  })
}

export const findCertCertificate = (params) => {
  return service({
    url: '/certCertificate/findCertCertificate',
    method: 'get',
    params
  })
}

export const getCertCertificateList = (params) => {
  return service({
    url: '/certCertificate/getCertCertificateList',
    method: 'get',
    params
  })
}

export const probeCertificate = (data) => {
  return service({
    url: '/certCertificate/probeCertificate',
    method: 'post',
    data
  })
}

export const updateAllCertificates = () => {
  return service({
    url: '/certCertificate/updateAllCertificates',
    method: 'post'
  })
}
