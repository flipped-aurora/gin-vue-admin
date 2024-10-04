import { login, getUserInfo, setSelfInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'
import { useRouterStore } from './router'
import cookie from 'js-cookie'
import i18n from '@/i18n' // added by mohamed hassan to multilangauge
import {useAppStore} from "@/pinia";


export const useUserStore = defineStore('user', () => {
  const appStore = useAppStore()
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {},
  })
  const token = ref(window.localStorage.getItem('token') || cookie.get('x-token') || '')
  const language = ref(window.localStorage.getItem('language') || cookie.get('language') || 'en') // added by mohamed hassan to allow store selected language for multilanguage support.
  const setUserInfo = (val) => {
    userInfo.value = val
    if(val.originSetting){
      Object.keys(appStore.config).forEach(key => {
        appStore.config[key] = val.originSetting[key]
      })
    }
  }

  const setToken = (val) => {
    token.value = val
  }

  // added by mohame hassan to allow store selected language for multilanguage support.
  const setLanguage = (val) => {
    console.log('setLanguage called with value: ' + val)
    language.value = val
  }

  const getLanguage = () => {
    return language.value
  }

  const NeedInit = async () => {
    token.value = ''
    window.localStorage.removeItem('token')
    await router.push({name: 'Init', replace: true})
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async() => {
    const res = await getUserInfo()
    if (res.code === 0) {
      setUserInfo(res.data.userInfo)
    }
    return res
  }
  /* 登录*/
  const LoginIn = async(loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: i18n.global.t('pinia.modules.user.loggingIn'),
    })

    const res = await login(loginInfo)

    // 登陆失败，直接返回
    if (res.code !== 0) {
      loadingInstance.value.close()
      return false
    }

    // 登陆成功，设置用户信息和权限相关信息
    setUserInfo(res.data.user)
    setToken(res.data.token)

    // 初始化路由信息
    const routerStore = useRouterStore()
    await routerStore.SetAsyncRouter()
    const asyncRouters = routerStore.asyncRouters

    // 注册到路由表里
    asyncRouters.forEach(asyncRouter => {
      router.addRoute(asyncRouter)
    })

    if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
      ElMessage.error(i18n.global.t('pinia.modules.user.connectAdmin'))
    } else {
      await router.replace({ name: userInfo.value.authority.defaultRouter })
    }

    const isWin = ref(/windows/i.test(navigator.userAgent))
    if (isWin.value) {
      window.localStorage.setItem('osType', 'WIN')
    } else {
      window.localStorage.setItem('osType', 'MAC')
    }


    // 全部操作均结束，关闭loading并返回
    loadingInstance.value.close()
    return true
  }
  /* 登出*/
  const LoginOut = async() => {
    const res = await jsonInBlacklist()

    // 登出失败
    if (res.code !== 0) {
      return
    }

    await ClearStorage()

    // 把路由定向到登录页，无需等待直接reload
    router.push({ name: 'Login', replace: true })
    window.location.reload()
  }
  /* 清理数据 */
  const ClearStorage = async() => {
    token.value = ''
    sessionStorage.clear()
    window.localStorage.removeItem('token')
    cookie.remove('x-token')
    localStorage.removeItem('originSetting')
  }

  watch(() => token.value, () => {
    window.localStorage.setItem('token', token.value)
  })

  return {
    userInfo,
    token,
    language, // added by mohame hassan to allow store selected language for multilanguage support.
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    setLanguage, // added by mohame hassan to allow store selected language for multilanguage support.
    getLanguage, // added by mohame hassan to allow store selected language for multilanguage support.
    setToken,
    loadingInstance,
    ClearStorage
  }
})
