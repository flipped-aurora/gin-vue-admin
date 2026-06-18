import { defineStore } from 'pinia'
import { ref, watchEffect, reactive, computed } from 'vue'
import { applyTheme } from '@/utils/theme'
import { buildPresetFromConfig } from '@/utils/themePreset'
import { useDark, usePreferredDark } from '@vueuse/core'

export const useAppStore = defineStore('app', () => {
  const device = ref('')
  const drawerSize = ref('')
  const operateMinWith = ref('240')
  // 侧栏折叠态（运行时态，不持久化），供 vertical 布局的固定侧栏与外壳共享
  const sideCollapse = ref(false)
  const config = reactive({
    weakness: false,
    grey: false,
    primaryColor: '#5D87FF',
    // Element Plus 语义色
    successColor: '#60C041',
    warningColor: '#F9901F',
    dangerColor: '#F56C6C',
    infoColor: '#909399',
    showTabs: true,
    darkMode: 'auto',
    layout_side_width: 256,
    layout_side_collapsed_width: 80,
    layout_side_item_height: 48,
    show_watermark: true,
    side_mode: 'normal',
    // 页面过渡动画配置
    transition_type: 'slide',
    global_size: 'default',
    // 菜单风格 'design' | 'light' | 'dark'
    menu_theme: 'design',
    // 全局圆角（rem）
    radius: 0.75,
    // 卡片盒子模式 'border' | 'shadow'
    card_mode: 'border',
    // 顶栏按钮可见性
    show_breadcrumb: true,
    show_refresh: true,
    show_search: true,
    show_collapse_btn: true,
    // 顶栏 / 标签栏配色（空字符串 = 跟随主题默认，支持明暗自适应）
    header_bg: '',
    header_border: '',
    tabs_bg: ''
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

  // vertical 布局固定侧栏的当前宽度
  const sideWidth = computed(() =>
    sideCollapse.value
      ? config.layout_side_collapsed_width
      : config.layout_side_width
  )

  const toggleSideCollapse = (v) => {
    sideCollapse.value = typeof v === 'boolean' ? v : !sideCollapse.value
  }

  const toggleMenuTheme = (e) => {
    config.menu_theme = e
  }

  const toggleRadius = (e) => {
    config.radius = e
  }

  const toggleCardMode = (e) => {
    config.card_mode = e
  }

  // key: 'successColor' | 'warningColor' | 'dangerColor' | 'infoColor'
  const toggleSemanticColor = (key, e) => {
    if (key in config) config[key] = e
  }

  // 应用一个主题预设
  const applyPreset = (preset) => {
    if (preset?.config) {
      Object.keys(preset.config).forEach((key) => {
        if (key in config) config[key] = preset.config[key]
      })
    }
  }

  // 以当前配置导出为预设对象
  const exportPreset = (name = 'current') => buildPresetFromConfig(config, name)

  const baseCoinfg = {
    weakness: false,
    grey: false,
    primaryColor: '#5D87FF',
    successColor: '#60C041',
    warningColor: '#F9901F',
    dangerColor: '#F56C6C',
    infoColor: '#909399',
    showTabs: true,
    darkMode: 'auto',
    layout_side_width: 256,
    layout_side_collapsed_width: 80,
    layout_side_item_height: 48,
    show_watermark: true,
    side_mode: 'normal',
    // 页面过渡动画配置
    transition_type: 'slide',
    global_size: 'default',
    menu_theme: 'design',
    radius: 0.75,
    card_mode: 'border',
    show_breadcrumb: true,
    show_refresh: true,
    show_search: true,
    show_collapse_btn: true,
    header_bg: '',
    header_border: '',
    tabs_bg: ''
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

  // 监听主题（主色 + 语义色 + 圆角 + 菜单风格 + 卡片模式）
  watchEffect(() => {
    applyTheme(config, isDark.value)
  })

  return {
    isDark,
    device,
    drawerSize,
    operateMinWith,
    sideCollapse,
    sideWidth,
    toggleSideCollapse,
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
    toggleGlobalSize,
    toggleMenuTheme,
    toggleRadius,
    toggleCardMode,
    toggleSemanticColor,
    applyPreset,
    exportPreset
  }
})
