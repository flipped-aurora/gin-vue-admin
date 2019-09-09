import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { User } from "@/store/module/user"
Vue.use(Vuex)



const vuexLocal = new VuexPersistence({
    storage: window.localStorage
})
export const store = new Vuex.Store({
    modules: {
        User
    },
    plugins: [vuexLocal.plugin]
})