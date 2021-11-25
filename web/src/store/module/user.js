import { login, getUserInfo, setUserInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElMessage } from 'element-plus'

export const user = {
  namespaced: true,
  state: {
    userInfo: {
      uuid: '',
      nickName: '',
      headerImg: '',
      authority: {},
      sideMode: 'dark',
      activeColor: '#4D70FF',
      baseColor: '#fff'
    },
    token: '',
  },
  mutations: {
    setUserInfo(state, userInfo) {
      // 这里的 `state` 对象是模块的局部状态
      state.userInfo = userInfo
    },
    setToken(state, token) {
      // 这里的 `state` 对象是模块的局部状态
      state.token = token
    },
    NeedInit(state) {
      state.userInfo = {}
      state.token = ''
      sessionStorage.clear()
      router.push({ name: 'Init', replace: true })
    },
    LoginOut(state) {
      state.userInfo = {}
      state.token = ''
      sessionStorage.clear()
      router.push({ name: 'Login', replace: true })
      window.location.reload()
    },
    ResetUserInfo(state, userInfo = {}) {
      state.userInfo = { ...state.userInfo,
        ...userInfo
      }
    },
    ChangeSideMode: (state, val) => {
      state.userInfo.sideMode = val
    },
  },
  actions: {
    async GetUserInfo({ commit }) {
      const res = await getUserInfo()
      if (res.code === 0) {
        commit('setUserInfo', res.data.userInfo)
      }
      return res
    },
    async LoginIn({ commit, dispatch, rootGetters, getters }, loginInfo) {
      const res = await login(loginInfo)
      if (res.code === 0) {
        commit('setUserInfo', res.data.user)
        commit('setToken', res.data.token)
        await dispatch('router/SetAsyncRouter', {}, { root: true })
        const asyncRouters = rootGetters['router/asyncRouters']
        asyncRouters.forEach(asyncRouter => {
          router.addRoute(asyncRouter)
        })
        // const redirect = router.history.current.query.redirect
        // console.log(redirect)
        // if (redirect) {
        //     router.push({ path: redirect })
        // } else {
        router.push({ name: getters['userInfo'].authority.defaultRouter })
        // }
        return true
      }
    },
    async LoginOut({ commit }) {
      const res = await jsonInBlacklist()
      if (res.code === 0) {
        commit('LoginOut')
      }
    },
    async changeSideMode({ commit, state }, data) {
      const res = await setUserInfo({ sideMode: data, ID: state.userInfo.ID })
      if (res.code === 0) {
        commit('ChangeSideMode', data)
        ElMessage({
          type: 'success',
          message: '设置成功'
        })
      }
    },
  },
  getters: {
    userInfo(state) {
      return state.userInfo
    },
    token(state) {
      return state.token
    },
    mode(state) {
      return state.userInfo.sideMode
    },
    sideMode(state) {
      if (state.userInfo.sideMode === 'dark') {
        return '#191a23'
      } else if (state.userInfo.sideMode === 'light') {
        return '#fff'
      } else {
        return state.userInfo.sideMode
      }
    },
    baseColor(state) {
      if (state.userInfo.sideMode === 'dark') {
        return '#fff'
      } else if (state.userInfo.sideMode === 'light') {
        return '#191a23'
      } else {
        return state.userInfo.baseColor
      }
    },
    activeColor(state) {
      if (state.userInfo.sideMode === 'dark' || state.userInfo.sideMode === 'light') {
        return '#4D70FF'
      }
      return state.userInfo.activeColor
    }
  }
}
