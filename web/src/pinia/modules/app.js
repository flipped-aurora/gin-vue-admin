import { defineStore } from 'pinia'
import { ref, watchEffect, reactive, computed } from 'vue'
import { applyTheme } from '@/utils/theme'
import {
  buildPresetFromConfig,
  resolveElConfig,
  PRESET_KEYS,
  BASE_CONFIG
} from '@/theme'
import { useDark, usePreferredDark } from '@vueuse/core'

export const useAppStore = defineStore('app', () => {
  const device = ref('')
  const drawerSize = ref('')
  const operateMinWith = ref('240')
  // 侧栏折叠态（运行时态，不持久化），供 vertical 布局的固定侧栏与外壳共享
  const sideCollapse = ref(false)
  // 系统配置（gva-theme）：初始值来自 BASE_CONFIG（出厂默认），字段说明见 @/theme
  const config = reactive({ ...BASE_CONFIG })

  // <el-config-provider> 的显式覆盖（ElConf），仅存「被显式声明」的字段；空 = 全部派生 / 回落默认。
  // 最终注入用 elConfig 计算（ElConf 显式值 > 由 config 派生 > BASE_ELCONF），见 @/theme resolveElConfig。
  const elConf = reactive({})
  const elConfig = computed(() => resolveElConfig(config, elConf))

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

  // 应用一个主题预设：整包补齐式覆盖——遍历全部系统配置键 PRESET_KEYS，
  // 预设未声明的键回落 BASE_CONFIG（如顶栏配色被清空），保证切换干净、不残留；
  // ElConf 整包替换（未提供则清空 = 全部派生 / 回落默认）。
  const applyPreset = (preset) => {
    if (!preset?.config) return
    PRESET_KEYS.forEach((key) => {
      config[key] =
        preset.config[key] === undefined ? BASE_CONFIG[key] : preset.config[key]
    })
    Object.keys(elConf).forEach((k) => delete elConf[k])
    Object.assign(elConf, preset.ElConf || {})
  }

  // 以当前配置导出为预设对象（整包：全部系统配置 + ElConf 覆盖）
  const exportPreset = (name = 'current') =>
    buildPresetFromConfig(config, name, elConf)

  const resetConfig = () => {
    for (let key in BASE_CONFIG) {
      config[key] = BASE_CONFIG[key]
    }
    Object.keys(elConf).forEach((k) => delete elConf[k])
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
    elConfig,
    applyPreset,
    exportPreset
  }
})
