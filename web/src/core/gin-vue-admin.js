/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export const run = function(app) {
    register(app)
    console.log(
        `%c gva启动完成 %c 当前版本V2.4.6 Apache%c`,
        'background:#35495e ; padding: 1px; border-radius: 3px 0 0 3px;  color: #fff',
        'background:#007aff ;padding: 1px 5px; border-radius: 0 3px 3px 0;  color: #fff; font-weight: bold;',
        'background:transparent'
    )
    console.log(`欢迎使用 %c Gin-Vue-Admin %c 
%c 微信 %c shouzi_1994 %c 
%c  QQ群 %c 622360840 %c 
  `,
        'background:#007aff ;padding: 1px 5px; border-radius: 3px;  color: #fff; font-weight: bold;margin-bottom:4px;',
        'background:transparent',
        'background:#35495e ; padding: 1px; border-radius: 3px 0 0 3px;  color: #fff',
        'background:#007aff ;padding: 1px 5px; border-radius: 0 3px 3px 0;  color: #fff; font-weight: bold;;margin-bottom:4px;',
        'background:transparent',
        'background:#35495e ; padding: 1px; border-radius: 3px 0 0 3px;  color: #fff',
        'background:#007aff ;padding: 1px 5px; border-radius: 0 3px 3px 0;  color: #fff; font-weight: bold;;margin-bottom:4px;',
        'background:transparent')
}