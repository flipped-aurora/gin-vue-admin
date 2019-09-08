import Vue from 'vue'
import App from './App.vue'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
// 引入封装的router
import router from '@/router/index'

Vue.config.productionTip = false

new Vue({
    render: h => h(App),
    router
}).$mount('#app')