import variables from '@/style/element_visiable.scss'
export const app = {
  namespaced: true,
  state: {
    theme: variables.colorPrimary,
    sideMode: 'dark'
  },
  mutations: {
    CHANGETHEME: (state, value) => {
      state.theme = value
    },
    CHANGESIDEMODE: (state) => {
      if (state.sideMode === 'dark') {
        state.sideMode = 'light'
      } else {
        state.sideMode = 'dark'
      }
    }
  },
  actions: {
    changeTheme({ commit }, data) {
      commit('CHANGETHEME', data)
    },
    changeSideMode({ commit }) {
      commit('CHANGESIDEMODE')
    }
  },
  getters: {
    getSIdeMode(state) {
      return state.sideMode
    }
  }
}
