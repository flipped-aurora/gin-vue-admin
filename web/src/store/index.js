import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from '@/store/module/user'
import { router } from '@/store/module/router'
import { dictionary } from '@/store/module/dictionary'

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user']
})

export const store = createStore({
  modules: {
    user,
    router,
    dictionary
  },
  plugins: [vuexLocal.plugin]
})
