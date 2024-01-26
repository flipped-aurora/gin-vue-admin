import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import './style/element_visiable.scss'
import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from '@/router/index'
import { store } from '@/pinia'

const app = createApp(App)

app
    .use(store)
    .use(router)
    .mount('#app')

export default app