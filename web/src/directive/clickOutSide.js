export default {
  install: (app) => {
    app.directive('click-outside', {
      mounted(el, binding) {
        const handler = (e) => {
          // 如果绑定的元素包含事件目标，或元素已经被移除，则不触发
          if (!el || el.contains(e.target) || e.target === el) return
          // 支持函数或对象 { handler: fn, exclude: [el1, el2], capture: true }
          const value = binding.value
          if (value && typeof value === 'object') {
            if (
              value.exclude &&
              value.exclude.some(
                (ex) => ex && ex.contains && ex.contains(e.target)
              )
            )
              return
            if (typeof value.handler === 'function') value.handler(e)
          } else if (typeof value === 'function') {
            value(e)
          }
        }

        // 存到 el 上，便于解绑
        el.__clickOutsideHandler__ = handler

        // 延迟注册，避免 mounted 时触发（比如当点击就是触发绑定动作时）
        setTimeout(() => {
          document.addEventListener('mousedown', handler)
          document.addEventListener('touchstart', handler)
        }, 0)
      },
      unmounted(el) {
        const h = el.__clickOutsideHandler__
        if (h) {
          document.removeEventListener('mousedown', h)
          document.removeEventListener('touchstart', h)
          delete el.__clickOutsideHandler__
        }
      }
    })
  }
}
