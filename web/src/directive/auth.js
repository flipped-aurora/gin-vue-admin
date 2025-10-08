// 权限按钮展示指令
import { useUserStore } from '@/pinia/modules/user'
export default {
  install: (app) => {
    const userStore = useUserStore()
    app.directive('auth', {
      // 当被绑定的元素插入到 DOM 中时……
      mounted: function (el, binding) {
        const userInfo = userStore.userInfo
        if (!binding.value){
          el.parentNode.removeChild(el)
          return
        }
        const waitUse = binding.value.toString().split(',')
        let flag = waitUse.some((item) => Number(item) === userInfo.authorityId)
        if (binding.modifiers.not) {
          flag = !flag
        }
        if (!flag) {
          el.parentNode.removeChild(el)
        }
      }
    })
  }
}
