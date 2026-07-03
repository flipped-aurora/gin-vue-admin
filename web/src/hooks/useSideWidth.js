import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useAppStore, useThemeStore } from '@/pinia'

/**
 * 侧栏当前宽度（折叠取 sideCollapsedWidth，否则 sideWidth）。
 * 供各布局 mode 与外壳（header / layout）共用，避免宽度公式散落重复。
 * @returns {{ sideWidth: import('vue').ComputedRef<number> }}
 */
export function useSideWidth() {
  const { sideCollapse } = storeToRefs(useAppStore())
  const { settings } = storeToRefs(useThemeStore())

  const sideWidth = computed(() =>
    sideCollapse.value
      ? settings.value.layout.sideCollapsedWidth
      : settings.value.layout.sideWidth
  )

  return { sideWidth }
}
