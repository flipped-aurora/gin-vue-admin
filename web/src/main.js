import Vue from 'vue'
import App from './App.vue'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 全局配置elementui的dialog不能通过点击遮罩层关闭
ElementUI.Dialog.props.closeOnClickModal.default = false
Vue.use(ElementUI);
// 引入封装的router
import router from '@/router/index'

// canvas背景插件
import vueParticleLine from 'vue-particle-line'
import 'vue-particle-line/dist/vue-particle-line.css'
Vue.use(vueParticleLine)

// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'

// 富文本插件
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

Vue.use(VueQuillEditor)

// markdown插件
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

Vue.use(mavonEditor)

import '@/permission'
import { store } from '@/store/index'
Vue.config.productionTip = false

// 路由守卫
import Bus from '@/utils/bus.js'
Vue.use(Bus)

import APlayer from '@moefe/vue-aplayer';

Vue.use(APlayer, {
    defaultCover: 'https://github.com/u3u.png',
    productionTip: true,
});


new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

//引入echarts
import echarts from 'echarts'
Vue.prototype.$echarts = echarts;



