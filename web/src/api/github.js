import axios from 'axios'
import { ElLoading } from 'element-plus'

let loadingInstance
const service = axios.create()

service.interceptors.request.use((config) => {
  loadingInstance = ElLoading.service({ fullscreen: true })
  return config
})

service.interceptors.response.use((resp) => {
  loadingInstance.close()
  return resp
}, (error) => {
  loadingInstance.close()
  return error
})

export function Commits(page) {
  return service({
    url: 'https://api.github.com/repos/flipped-aurora/gin-vue-admin/commits?page=' + page,
    method: 'get'
  })
}

export function Members() {
  return service({
    url: 'https://api.github.com/orgs/FLIPPED-AURORA/members',
    method: 'get'
  })
}
