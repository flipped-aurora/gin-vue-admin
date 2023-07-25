// 监听 window 的 resize 事件，返回当前窗口的宽高
import { shallowRef } from 'vue'
import { getCurrentInstance, nextTick, onMounted } from 'vue-demi'

export function tryOnMounted(fn, sync = true) {
  if (getCurrentInstance()) { onMounted(fn) } else if (sync) { fn() } else { nextTick(fn) }
}

const width = shallowRef(0)
const height = shallowRef(0)

export const useWindowSize = (cb) => {
  const onResize = () => {
    width.value = window.innerWidth
    height.value = window.innerHeight
    if (cb && typeof cb === 'function') {
      cb(width.value, height.value)
    }
  }

  tryOnMounted(onResize)
  window.addEventListener('resize', onResize, { passive: true })
  return {
    width,
    height,
  }
}
