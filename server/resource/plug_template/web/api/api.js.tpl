import service from '@/utils/request'

// @Router /{{ .RouterGroup }}/routerName
export const routerName = (data) => {
  return service({
    url: '/{{ .RouterGroup }}/routerName',
    method: 'post',
    data
  })
}