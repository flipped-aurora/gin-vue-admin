import { login } from '@/api/user'
import router from '@/router/index'
import { Message } from 'element-ui'
export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            nickName: "",
            headerImg: "",
            authority: "",
        },
        token: "",
        expiresAt: ""
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
        setExpiresAt(state, expiresAt) {
            // 这里的 `state` 对象是模块的局部状态
            state.expiresAt = expiresAt
        },
        clearAll(state) {
            state.userInfo = {}
            state.token = ""
            state.expiresAt = ""
        }

    },
    actions: {
        async LoginIn({ commit }, loginInfo) {
            try {
                const res = await login(loginInfo)
                commit('setUserInfo', res.data.user)
                commit('setToken', res.data.token)
                commit('setExpiresAt', res.data.expiresAt)
                if (res.success) {
                    const redirect = router.history.current.query.redirect
                    if (redirect) {
                        router.push({ path: redirect })
                    } else {
                        router.push({ name: 'dashboard' })
                    }
                }
            } catch (err) {
                Message({
                    type: 'error',
                    message: err,
                    showClose: true
                })
                return Promise.reject(err)
            }
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        },
        expiresAt(state) {
            return state.expiresAt
        }
    }
}