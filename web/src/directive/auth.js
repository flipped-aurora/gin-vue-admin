// 权限按钮展示指令
import { store } from '@/store/index'
const userInfo = store.getters['user/userInfo']
export const auth = (Vue) => {
    Vue.directive('auth', {
        // 当被绑定的元素插入到 DOM 中时……
        bind: function (el, binding) {
          let type = ""
          switch (Object.prototype.toString.call(binding.value)) {
            case "[object Array]":
              type = "Array"
              break;
            case "[object String]":
              type = "String"
              break;
            case "[object Number]":
              type = "Number"
              break;
            default:
              type = ""
              break;
          }
          if (type === "") {
            /* eslint-disable */
            console.error("v-auth必须是Array,Number,String属性,暂不支持其他属性")
            /* eslint-enable */
            return
          }
          const waitUse = binding.value.toString().split(",")
      
          let flag = waitUse.some(item=>item==userInfo.authorityId)
          if (binding.modifiers.not) {
            flag = !flag
          }
          if(!flag){
            el.style.display = 'none'
          }
        }
      })
}
