import Vue from 'vue'
import App from './App.vue'

//  按需引入element
import {
    Button, 
    Select, 
    Dialog, 
    Form, 
    Input,
    FormItem, 
    Option, 
    Loading, 
    Message, 
    Container, 
    Card,
    Dropdown,
    DropdownMenu,
    DropdownItem,
    Row,
    Col,
    Menu,
    Submenu,
    MenuItem,
    Aside,
    Main,
    Badge,
    Header,
    Tabs,
    Breadcrumb,
    BreadcrumbItem,
    Scrollbar,
    Avatar,
    TabPane,
    Divider,
    Table,
    TableColumn,
    Cascader,
    Checkbox,
    CheckboxGroup,
    Pagination,
    Tag,
    Drawer,
    Tree,
    Popover,
    Switch,
    Collapse,
    CollapseItem,
    Tooltip,
    DatePicker,
    InputNumber,
    Steps,
    Upload,
    Progress
} from 'element-ui';

Vue.use(Button);
Vue.use(Select);
Vue.use(Dialog);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Input);
Vue.use(Option);
Vue.use(Container);
Vue.use(Card);
Vue.use(Dropdown);
Vue.use(DropdownMenu);
Vue.use(DropdownItem);
Vue.use(Row);
Vue.use(Col);
Vue.use(Menu);
Vue.use(Submenu);
Vue.use(MenuItem);
Vue.use(Aside);
Vue.use(Main);
Vue.use(Badge);
Vue.use(Header);
Vue.use(Tabs);
Vue.use(Breadcrumb);
Vue.use(BreadcrumbItem);
Vue.use(Avatar);
Vue.use(TabPane);
Vue.use(Divider);
Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Checkbox);
Vue.use(Cascader);
Vue.use(Tag);
Vue.use(Pagination);
Vue.use(Drawer);
Vue.use(Tree);
Vue.use(CheckboxGroup);
Vue.use(Popover);
Vue.use(InputNumber);
Vue.use(Switch);
Vue.use(Collapse);
Vue.use(CollapseItem);
Vue.use(Tooltip);
Vue.use(DatePicker);
Vue.use(Steps);
Vue.use(Upload);
Vue.use(Progress);
Vue.use(Scrollbar);
Vue.use(Loading.directive);

Vue.prototype.$loading = Loading.service;
Vue.prototype.$message = Message;
Dialog.props.closeOnClickModal.default = false

// 引入封装的router
import router from '@/router/index'

// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'

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


import { auth } from '@/directive/auth'
// 按钮权限指令
auth(Vue)

import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

console.log(`
       欢迎使用 Gin-Vue-Admin
       当前版本:V2.4.0
       加群方式:微信：shouzi_1994 QQ群：622360840
       默认自动化文档地址:http://127.0.0.1:` + buildConfig.goServerPort + `/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:`+ buildConfig.vueClientPort + `
       如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`)