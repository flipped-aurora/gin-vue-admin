import { defineStore } from 'pinia'
import { ref, watchEffect, reactive } from 'vue'
import { setBodyPrimaryColor } from '@/utils/format'
import { useDark, usePreferredDark } from '@vueuse/core'

export const useAppStore = defineStore('app', () => {
  const device = ref('')
  const drawerSize = ref('')
  const operateMinWith = ref('240')
  const config = reactive({
    weakness: false,
    grey: false,
    primaryColor: '#3b82f6',
    showTabs: true,
    darkMode: 'auto',
    layout_side_width: 256,
    layout_side_collapsed_width: 80,
    layout_side_item_height: 48,
    show_watermark: true,
    side_mode: 'normal',
    // 页面过渡动画配置
    transition_type: 'slide',
    global_size: 'default'
  })

  const isDark = useDark({
    selector: 'html',
    attribute: 'class',
    valueDark: 'dark',
    valueLight: 'light'
  })

  const preferredDark = usePreferredDark()

  const toggleTheme = (darkMode) => {
    isDark.value = darkMode
  }

  const toggleWeakness = (e) => {
    config.weakness = e
  }

  const toggleGrey = (e) => {
    config.grey = e
  }

  const togglePrimaryColor = (e) => {
    config.primaryColor = e
  }

  const toggleTabs = (e) => {
    config.showTabs = e
  }

  const toggleDevice = (e) => {
    if (e === 'mobile') {
      drawerSize.value = '100%'
      operateMinWith.value = '80'
    } else {
      drawerSize.value = '800'
      operateMinWith.value = '240'
    }
    device.value = e
  }

  const toggleDarkMode = (e) => {
    config.darkMode = e
  }

  // 监听系统主题变化
  watchEffect(() => {
    if (config.darkMode === 'auto') {
      isDark.value = preferredDark.value
      return
    }
    isDark.value = config.darkMode === 'dark'
  })

  const toggleConfigSideWidth = (e) => {
    config.layout_side_width = e
  }

  const toggleConfigSideCollapsedWidth = (e) => {
    config.layout_side_collapsed_width = e
  }

  const toggleConfigSideItemHeight = (e) => {
    config.layout_side_item_height = e
  }

  const toggleConfigWatermark = (e) => {
    config.show_watermark = e
  }

  const toggleSideMode = (e) => {
    config.side_mode = e
  }

  const toggleTransition = (e) => {
    config.transition_type = e
  }

  const toggleGlobalSize = (e) => {
    config.global_size = e
  }

  const baseCoinfg = {
    weakness: false,
    grey: false,
    primaryColor: '#3b82f6',
    showTabs: true,
    darkMode: 'auto',
    layout_side_width: 256,
    layout_side_collapsed_width: 80,
    layout_side_item_height: 48,
    show_watermark: true,
    side_mode: 'normal',
    // 页面过渡动画配置
    transition_type: 'slide',
    global_size: 'default'
  }

  const resetConfig = () => {
    for (let baseCoinfgKey in baseCoinfg) {
      config[baseCoinfgKey] = baseCoinfg[baseCoinfgKey]
    }
  }

  // 监听色弱模式和灰色模式
  watchEffect(() => {
    document.documentElement.classList.toggle('html-weakenss', config.weakness)
    document.documentElement.classList.toggle('html-grey', config.grey)
  })

  // 监听主题色
  watchEffect(() => {
    setBodyPrimaryColor(config.primaryColor, isDark.value ? 'dark' : 'light')
  })

  return {
    isDark,
    device,
    drawerSize,
    operateMinWith,
    config,
    toggleTheme,
    toggleDevice,
    toggleWeakness,
    toggleGrey,
    togglePrimaryColor,
    toggleTabs,
    toggleDarkMode,
    toggleConfigSideWidth,
    toggleConfigSideCollapsedWidth,
    toggleConfigSideItemHeight,
    toggleConfigWatermark,
    toggleSideMode,
    toggleTransition,
    resetConfig,
    toggleGlobalSize
  }
})
