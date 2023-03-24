import service from '@/utils/request'

export const getTable = (data) => {
  return service({
    url: '/chatGpt/getTable',
    method: 'post',
    data
  })
}
