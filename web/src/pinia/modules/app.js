import { defineStore } from 'pinia'
import { ref, watchEffect, reactive } from 'vue'
import { setBodyPrimaryColor } from '@/utils/format'
export const useAppStore = defineStore('app', () => {
  const device = ref('')
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
    side_mode: 'normal'
  })

  const theme = ref('auto')

  const toggleTheme = (dark) => {
    if (dark) {
      theme.value = 'dark'
    } else {
      theme.value = 'light'
    }
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
    device.value = e
  }

  const toggleDarkMode = (e) => {
    config.darkMode = e
  }

  const toggleDarkModeAuto = () => {
    // 处理浏览器主题
    const darkQuery = window.matchMedia('(prefers-color-scheme: dark)')
    const dark = darkQuery.matches
    toggleTheme(dark)
    darkQuery.addEventListener('change', (e) => {
      toggleTheme(e.matches)
    })
  }

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

  const toggleSideModel = (e) => {
    config.side_mode = e
  }

  watchEffect(() => {
    if (theme.value === 'dark') {
      document.documentElement.classList.add('dark')
      document.documentElement.classList.remove('light')
    } else {
      document.documentElement.classList.add('light')
      document.documentElement.classList.remove('dark')
    }
  })
  watchEffect(() => {
    // 色弱模式监听处理
    if (config.weakness) {
      document.documentElement.classList.add('html-weakenss')
    } else {
      document.documentElement.classList.remove('html-weakenss')
    }
  })
  watchEffect(() => {
    // 灰色模式监听处理
    if (config.grey) {
      document.documentElement.classList.add('html-grey')
    } else {
      document.documentElement.classList.remove('html-grey')
    }
  })

  watchEffect(() => {
    if (config.darkMode === 'auto') {
      toggleDarkModeAuto()
    }

    if (config.darkMode === 'dark') {
      toggleTheme(true)
    } else {
      toggleTheme(false)
    }
  })

  watchEffect(() => {
    setBodyPrimaryColor(config.primaryColor, theme.value)
  })

  return {
    theme,
    device,
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
    toggleSideModel
  }
})
