import { computed, reactive, watch, watchEffect } from 'vue'
import { defineStore } from 'pinia'
import { ElMessage } from 'element-plus'
import { useDark, useDebounceFn, usePreferredDark } from '@vueuse/core'
import { setSelfSetting } from '@/api/user'
import {
  buildThemeEnvelope,
  buildThemePreset,
  cacheThemeSettings,
  mergeThemeSettings,
  normalizeColor,
  normalizeThemePreset,
  normalizeThemeSettings,
  readCachedThemeSettings,
  toggleAuxiliaryColorModes,
  toggleCssDarkMode
} from '@/theme/shared'
import { cloneThemeSettings } from '@/theme/settings'
import { getThemeColors } from '@/theme/token'
import { resolveElementPlusConfig } from '@/theme/adapters/element-plus'
import { dispatchTheme } from '@/theme/adapters/runtime'

// 轻量登录态判断：避免在登录页等未登录场景向 setSelfSetting 发请求。
// 直接读 token，不依赖 userStore，规避 user <-> theme 的循环引用。
const hasAuthToken = () =>
  Boolean(localStorage.getItem('token')) || document.cookie.includes('x-token=')

export const useThemeStore = defineStore('theme', () => {
  const settings = reactive(readCachedThemeSettings() || cloneThemeSettings())

  const isDark = useDark({
    selector: 'html',
    attribute: 'class',
    valueDark: 'dark',
    valueLight: 'light',
    storageKey: null
  })
  const preferredDark = usePreferredDark()

  const darkMode = computed(() => {
    if (settings.themeScheme === 'auto') return preferredDark.value
    return settings.themeScheme === 'dark'
  })

  const themeColors = computed(() => getThemeColors(settings))
  const elConfig = computed(() => resolveElementPlusConfig(settings))
  const primaryColor = computed(() => themeColors.value.primary)
  const settingsJson = computed(() => JSON.stringify(settings))

  // 远端同步基线：记录「已知与后端一致」的 settings 快照。用值比较而非时序标志，
  // 天然规避 watch 的异步触发顺序问题——应用后端 originSetting 时把基线设为该值，
  // 防抖保存触发时发现与基线一致即跳过，从而不会把刚拉取的服务端值再回写给服务端。
  let remoteBaseline = settingsJson.value
  let remoteSaving = false
  let remoteSavePending = false

  // 远端保存串行化：任意时刻最多一个在途请求，完成后若期间又有改动则再保存一次。
  // 这样「最后一次配置」始终最后落库，避免旧请求晚返回覆盖新配置导致的静默丢配置。
  const saveSelfSettingToRemote = async () => {
    if (remoteSaving) {
      remoteSavePending = true // 已有请求在途，记下「完成后需再保存」
      return
    }

    const json = settingsJson.value
    if (json === remoteBaseline) return // 与后端一致，无需保存
    if (!hasAuthToken()) return // 未登录，跳过（本地缓存已落地）

    remoteSaving = true
    try {
      const res = await setSelfSetting(buildThemeEnvelope(settings))
      if (res?.code === 0) {
        remoteBaseline = json
        ElMessage.success({
          message: '保存成功'
        })
      }
    } catch {
      // 静默失败：本地缓存仍是最新，下次任意改动会重新触发保存
    } finally {
      remoteSaving = false
      if (remoteSavePending) {
        remoteSavePending = false
        saveSelfSettingToRemote() // 在途期间又有改动 → 立即用最新值再保存一次
      }
    }
  }
  // 配置项常被滑块、调色板等高频修改，防抖后再保存，避免每次微调都打接口
  const debouncedRemoteSave = useDebounceFn(saveSelfSettingToRemote, 500)

  const setThemeScheme = (themeScheme) => {
    settings.themeScheme = themeScheme
  }

  const toggleTheme = (dark) => {
    settings.themeScheme = dark ? 'dark' : 'light'
  }

  const updateThemeColors = (key, color) => {
    // 颜色统一小写后再写入，保证内部 themeColor / otherColor 始终是小写单一形态
    const value = normalizeColor(color)
    if (key === 'primary') {
      settings.themeColor = value
      return
    }

    if (key in settings.otherColor) settings.otherColor[key] = value
  }

  const applyPreset = (preset) => {
    const normalized = normalizeThemePreset(preset)
    if (!normalized?.theme) return

    mergeThemeSettings(settings, cloneThemeSettings())
    mergeThemeSettings(settings, normalized.theme)
  }

  // 应用后端返回的用户配置：normalizeThemeSettings 内部已对旧/未知数据做防御式回落，
  // 应用后把它设为远端基线，避免把服务端值再次回写。
  const applyRemoteSettings = (raw) => {
    // 防御：本地若有未同步到后端的改动（保存在途 / 与基线不一致），说明本地比服务端新，
    // 跳过覆盖，避免回退用户正在保存的配置。当前时序（登录 / 首次进入各调用一次）下恒不命中，
    // 仅为「GetUserInfo 未来若变成可重入（轮询 / 切回标签页刷新）」时保持 store 自洽。
    if (remoteSaving || remoteSavePending || settingsJson.value !== remoteBaseline) return
    mergeThemeSettings(settings, normalizeThemeSettings(raw))
    remoteBaseline = settingsJson.value
  }

  const exportPreset = (name = 'current') => buildThemePreset(settings, name)

  const resetConfig = () => {
    mergeThemeSettings(settings, cloneThemeSettings())
  }

  watch(
    darkMode,
    (value) => {
      isDark.value = value
      toggleCssDarkMode(value)
    },
    { immediate: true }
  )

  watch(
    [() => settings.grayscale, () => settings.colourWeakness],
    ([grayscale, colourWeakness]) => {
      toggleAuxiliaryColorModes(grayscale, colourWeakness)
    },
    { immediate: true }
  )

  watchEffect(() => {
    dispatchTheme({
      settings,
      themeColors: themeColors.value,
      isDark: isDark.value
    })
  })

  // 统一持久化入口：任意 settings 变化 → 本地缓存（始终）+ 远端保存（登录态 & 与基线不同）。
  // 持久化不再依赖设置抽屉等组件挂载，store 在任何场景下都能自持久化。
  watch(settingsJson, () => {
    cacheThemeSettings(settings)
    debouncedRemoteSave()
  })

  return {
    settings,
    elConfig,
    isDark,
    darkMode,
    themeColors,
    primaryColor,
    settingsJson,
    setThemeScheme,
    toggleTheme,
    updateThemeColors,
    applyPreset,
    applyRemoteSettings,
    exportPreset,
    resetConfig
  }
})
