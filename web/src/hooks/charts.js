// 本组件参考 arco-pro 的实现
// https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/hooks/chart-option.ts

import { computed } from 'vue'
import { useAppStore } from '@/pinia'

export default function useChartOption(sourceOption) {
  const appStore = useAppStore()
  const isDark = computed(() => {
    return appStore.isDark
  })
  const chartOption = computed(() => {
    return sourceOption(isDark.value)
  })
  return {
    chartOption
  }
}
