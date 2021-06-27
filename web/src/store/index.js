import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from '@/store/module/user'
import { router } from '@/store/module/router'
import { dictionary } from '@/store/module/dictionary'
import { app } from '@/store/module/app'
Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user']
})
export const store = new Vuex.Store({
  modules: {
    user,
    router,
    dictionary,
    app
  },
  plugins: [vuexLocal.plugin]
})
