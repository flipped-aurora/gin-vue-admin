import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useAppStore, useThemeStore } from '@/pinia'

/**
 * 生效的布局模式：640–1024（平板）强制「通栏侧边(vertical)」，
 * 其余用用户所选 settings.layout.mode。仅用于渲染，不改写持久化设置，
 * 因此宽度回到 >= 1024 时会自动恢复用户原布局。
 * 平板区默认收缩（由 useResponsive 在进入区间时置一次），仍支持手动展开。
 * @returns {{ effectiveMode: import('vue').ComputedRef<string> }}
 */
export function useLayoutMode() {
  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { isPad } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)

  const effectiveMode = computed(() =>
    isPad.value ? 'vertical' : settings.value.layout.mode
  )

  return { effectiveMode }
}
