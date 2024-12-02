import { login, getUserInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouterStore } from './router'
import { useCookies } from '@vueuse/integrations/useCookies'
import { useStorage } from '@vueuse/core'

import { useAppStore } from '@/pinia'

export const useUserStore = defineStore('user', () => {
  const appStore = useAppStore()
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {}
  })
  const token = useStorage('token', '')
  const xToken = useCookies('x-token')
  const currentToken = computed(() => token.value || xToken.value || '')

  const setUserInfo = (val) => {
    userInfo.value = val
    if (val.originSetting) {
      Object.keys(appStore.config).forEach((key) => {
        appStore.config[key] = val.originSetting[key]
      })
    }
  }

  const setToken = (val) => {
    token.value = val
    xToken.value = val
  }

  const NeedInit = async () => {
    await ClearStorage()
    await router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async () => {
    const res = await getUserInfo()
    if (res.code === 0) {
      setUserInfo(res.data.userInfo)
    }
    return res
  }
  /* 登录*/
  const LoginIn = async (loginInfo) => {
    try {
      loadingInstance.value = ElLoading.service({
        fullscreen: true,
        text: '登录中，请稍候...'
      })

      const res = await login(loginInfo)
      
      if (res.code !== 0) {
        ElMessage.error(res.message || '登录失败')
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
      asyncRouters.forEach((asyncRouter) => {
        router.addRoute(asyncRouter)
      })

      if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
        ElMessage.error('请联系管理员进行授权')
      } else {
        await router.replace({ name: userInfo.value.authority.defaultRouter })
      }

      const isWindows = /windows/i.test(navigator.userAgent)
      window.localStorage.setItem('osType', isWindows ? 'WIN' : 'MAC')

      // 全部操作均结束，关闭loading并返回
      return true
    } catch (error) {
      console.error('LoginIn error:', error)
      return false
    } finally {
      loadingInstance.value?.close()
    }
  }
  /* 登出*/
  const LoginOut = async () => {
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
  const ClearStorage = async () => {
    token.value = ''
    xToken.value = ''
    sessionStorage.clear()
    localStorage.removeItem('originSetting')
  }

  return {
    userInfo,
    token: currentToken,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    setToken,
    loadingInstance,
    ClearStorage
  }
})
