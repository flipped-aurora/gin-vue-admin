// 本组件参考 arco-pro 的实现
// https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/hooks/responsive.ts

import { onMounted, onBeforeMount, onBeforeUnmount } from 'vue'
import { useDebounceFn } from '@vueuse/core'
import { useAppStore } from '@/pinia'
import { addEventListen, removeEventListen } from '@/utils/event'

const WIDTH = 992

function queryDevice() {
  const rect = document.body.getBoundingClientRect()
  return rect.width - 1 < WIDTH
}

export default function useResponsive(immediate) {
  const appStore = useAppStore()
  function resizeHandler() {
    if (!document.hidden) {
      const isMobile = queryDevice()
      appStore.toggleDevice(isMobile ? 'mobile' : 'desktop')
      // appStore.toggleDevice(isMobile);
    }
  }
  const debounceFn = useDebounceFn(resizeHandler, 100)
  onMounted(() => {
    if (immediate) debounceFn()
  })
  onBeforeMount(() => {
    addEventListen(window, 'resize', debounceFn)
  })
  onBeforeUnmount(() => {
    removeEventListen(window, 'resize', debounceFn)
  })
}
