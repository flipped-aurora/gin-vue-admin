import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from "@/store/module/user"
Vue.use(Vuex)



const vuexLocal = new VuexPersistence({
    storage: window.localStorage
})
export const store = new Vuex.Store({
    modules: {
        user
    },
    plugins: [vuexLocal.plugin]
})