import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import router from '@/router/index'
import { ElLoading } from 'element-plus'
import i18n from '@/i18n' // added by mohamed hassan to multilangauge

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let activeAxios = 0
let timer
let loadingInstance
const showLoading = (option = {
  target: null,
}) => {
  const loadDom = document.getElementById('gva-base-load-dom')
  activeAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (activeAxios > 0) {
      if (!option.target) option.target = loadDom
      loadingInstance = ElLoading.service(option)
    }
  }, 400)
}

const closeLoading = () => {
  activeAxios--
  if (activeAxios <= 0) {
    clearTimeout(timer)
    loadingInstance && loadingInstance.close()
  }
}
// http request 拦截器
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading(config.loadingOption)
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      'Accept-Language': userStore.language, // added by mohame hassan to allow store selected language for multilanguage support.
      ...config.headers
    }
    return config
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  response => {
    const userStore = useUserStore()
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      return response.data.msg ? response.data : response
    }
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      ElMessageBox.confirm(
        i18n.global.t('utils.request.requestErrorDetected') + 
        `<p>${error}</p>
        `, i18n.global.t('utils.request.requestError'), {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: i18n.global.t('utils.request.tryAgainLater'),
        cancelButtonText: i18n.global.t('general.cancel')
      })
      return
    }
    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(
          i18n.global.t('utils.request.interfaceErrorDetected') +   
          `<p>${error}</p>` + 
          `<p>` + i18n.global.t('utils.request.errorCode') + `<span style="color:red"> 500 </span>：` + i18n.global.t('utils.request.interfaceErrorNote') + `</p>`, 
          i18n.global.t('utils.request.interfaceError'), {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: i18n.global.t('utils.request.clearCache'),
          cancelButtonText: i18n.global.t('general.cancel')
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.ClearStorage()
            router.push({ name: 'Login', replace: true })
          })
        break
      case 404:
        ElMessageBox.confirm( 
          i18n.global.t('utils.request.interfaceErrorDetected') +   
          `<p>${error}</p>` + 
          `<p>` + i18n.global.t('utils.request.errorCode') + `<span style="color:red"> 404 </span>：` + i18n.global.t('utils.request.interfaceNotRegisteredNote') + `</p>`, 
          i18n.global.t('utils.request.interfaceError'), {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: i18n.global.t('utils.request.iGotIt'),
          cancelButtonText: i18n.global.t('general.cancel')
        })
        break
      case 401:
        ElMessageBox.confirm(
          i18n.global.t('utils.request.invalidToken') +   
          `<p>` + i18n.global.t('utils.request.errorCode') + `<span style="color:red"> 401 </span>` + i18n.global.t('utils.request.errorMessage') + `: ${error}</p>`, 
          i18n.global.t('utils.request.identityInfo'), {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: i18n.global.t('utils.request.loginAgain'),
          cancelButtonText: i18n.global.t('general.cancel')
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.ClearStorage()
            router.push({ name: 'Login', replace: true })
          })
        break
    }

    return error
  }
)
export default service
