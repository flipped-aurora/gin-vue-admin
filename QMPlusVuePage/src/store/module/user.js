export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            nickName: ""
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
        // AsyncSetUserInfo({ commit }, loginInfo) {

        // }
    },
    getters: {
        userName(state) {
            return state.userName
        }
    }
}