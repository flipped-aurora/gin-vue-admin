// 本组件参考 arco-pro 的实现
// https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/hooks/responsive.ts

import { onMounted, onBeforeMount, onBeforeUnmount } from 'vue'
import { useDebounceFn } from '@vueuse/core'
import { useAppStore } from '@/pinia'
import { addEventListen, removeEventListen } from '@/utils/event'

// 三档断点：
// - < 640：移动端，收起所有内联菜单，改用左侧抽屉
// - 640 ~ 1024：平板，强制「通栏侧边」收缩布局（不改用户所选布局）
// - >= 1024：桌面，使用用户所选布局
const MOBILE_WIDTH = 640
const PAD_WIDTH = 1024

const zoneOf = (width) =>
  width < MOBILE_WIDTH ? 'mobile' : width < PAD_WIDTH ? 'pad' : 'desktop'

export default function useResponsive(immediate) {
  const appStore = useAppStore()
  function resizeHandler() {
    if (document.hidden) return
    const zone = zoneOf(document.body.getBoundingClientRect().width)
    appStore.toggleDevice(zone === 'mobile' ? 'mobile' : 'desktop')
    // 断点区间切换的收缩/恢复编排收敛在 store（sideCollapse 的唯一写入方）
    appStore.setLayoutZone(zone)
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
