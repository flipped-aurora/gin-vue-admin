import { login } from '@/api/user'

export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            nickName: "",
            headerImg: "",
            authority: "",
        },
        token: ""
    },
    mutations: {
        setUserInfo(state, userInfo) {
            // 这里的 `state` 对象是模块的局部状态
            state.userInfo = userInfo
        },
        setToken(state, token) {
            // 这里的 `state` 对象是模块的局部状态
            state.token = token
        }
    },
    actions: {
        LoginIn({ commit }, loginInfo) {
            login(loginInfo).then(res => {
                commit('setUserInfo', res.data.user)
                commit('setToken', res.data.token)
                return res
            }).catch(err => {
                console.error(err)
                return Promise.reject(err)
            })
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        }
    }
}