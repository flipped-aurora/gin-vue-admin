export const User = {
    namespaced: true,
    state: { userName: "QM" },
    mutations: {
        setUserName(state, name) {
            // 这里的 `state` 对象是模块的局部状态
            state.userName = name
        }
    },
    actions: {
        AsyncSetUserName({ commit }, name) {
            setTimeout(() => {
                commit("setUserName", name)
            }, 2000);
        }
    },
    getters: {
        userName(state) {
            return state.userName
        }
    }
}