/*
                    商用代码公司自用产品无需授权
    若作为代码出售的产品（任何涉及代码交付第三方作为后续开发）必须保留此脚本
                         或标注原作者信息
                          否则将依法维权
*/

var child_process = require('child_process')

var url = 'https://www.gin-vue-admin.com'
var cmd = ''
console.log(process.platform)
switch (process.platform) {
  case 'win32':
    cmd = 'start'
    child_process.exec(cmd + ' ' + url)
    break

  case 'darwin':
    cmd = 'open'
    child_process.exec(cmd + ' ' + url)
    break
}
