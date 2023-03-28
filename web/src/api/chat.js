import service from '@/utils/request'

export const chat = (data) => {
  return service({
    url: '/chat/chat',
    method: 'post',
    data
  })
}
