'use strict'

module.exports = {
    title: 'GIN-VUE-ADMIN1',
    baseCdnUrl: '//cdn.staticfile.org',
    cdns: [
        /**
         * 如果设置path属性, { name: 'vue', scope: 'Vue', path: '/vue/2.6.9/vue.min.js' } 即编译出来以[baseCdnUrl][path]
         * 否则自动拼写 [baseCdnUrl]/[name]/[version]/[name].min.js
         * */ 
        { name: 'vue', scope: 'Vue' },
        { name: 'vue-router', scope: 'VueRouter' },
        { name: 'vuex', scope: 'Vuex' },
        { name: 'axios', scope: 'axios' },
        { name: 'element-ui', scope: 'ELEMENT',  path: '/element-ui/2.12.0/index.js'},
    ]
};
