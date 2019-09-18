import Vue from 'vue'
import App from './App.vue'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
// 引入封装的router

import vueParticleLine from 'vue-particle-line'
import 'vue-particle-line/dist/vue-particle-line.css'
Vue.use(vueParticleLine)
    // canvas背景插件

import router from '@/router/index'
import '@/permission'
import { store } from '@/store/index'
Vue.config.productionTip = false


import Bus from '@/utils/bus.js'
Vue.use(Bus)
new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')