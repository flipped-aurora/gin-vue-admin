/*
此文件受版权保护，未经授权禁止修改！如果您尚未获得授权，请通过微信(shouzi_1994)联系我们以购买授权。在未授权状态下，只需保留此代码，不会影响任何正常使用。
     未经授权的商用使用可能会被我们的资产搜索引擎爬取，并可能导致后续索赔。索赔金额将不低于高级授权费的十倍。请您遵守版权法律法规，尊重知识产权。
*/

import child_process from 'child_process'

var url = 'https://www.gin-vue-admin.com'
var cmd = ''
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
