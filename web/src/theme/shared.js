import { colorToRgbChannels } from './color'
import { cloneThemeSettings, themeSettings, THEME_SETTINGS_VERSION } from './settings'
import { GVA_MAIN_VERSION, PRESET_FALLBACK_MIN_VERSION } from './version'
import { themeVars } from './vars'
import { defu, isObject, liteClone } from 'arcdash'

export const THEME_SETTINGS_STORAGE_KEY = 'gva-theme-settings'

// 把 source 深合并进 target（就地，供 themeStore 直接改写响应式 settings）。
// target / source 都来自完整 schema（默认值或 normalizeThemeSettings 输出），
// 未知字段的过滤已在 normalize 阶段由 pickKnownKeys 完成，这里无需再做白名单。
export const mergeThemeSettings = (target, source) => {
  Object.entries(source || {}).forEach(([key, value]) => {
    if (isObject(value) && isObject(target[key])) {
      mergeThemeSettings(target[key], value)
      return
    }
    target[key] = value
  })
}

// 递归按默认 schema 过滤未知字段：schema 中不存在的键一律丢弃，
// 避免外部数据污染运行态、拖累 adapter 扩展。
const pickKnownKeys = (source, schema) => {
  if (!isObject(source) || !isObject(schema)) return source

  const result = {}
  Object.keys(schema).forEach((key) => {
    if (key in source) result[key] = pickKnownKeys(source[key], schema[key])
  })
  return result
}

const isThemeSettingsLike = (source) =>
  isObject(source) && Boolean(
    source.themeScheme ||
    source.themeColor ||
    source.otherColor ||
    source.layout ||
    source.page ||
    source.header ||
    source.tab ||
    source.menu ||
    source.card ||
    source.watermark ||
    source.tokens
  )

/**
 * 把任意外部主题数据（后端 originSetting、导入预设的内层 theme、本地缓存）规整成
 * 一份完整、可信的 settings。兼容策略（大版本，不迁移旧值，只做防御式回落）：
 * - `{ version, settings }` 信封：版本不等于当前版本 → 回落默认值；否则取 settings 继续。
 * - 旧的扁平主题数据（不含新结构字段）→ isThemeSettingsLike 为 false → 回落默认值。
 * - 合法的新结构 → 先按 schema 过滤未知键（pickKnownKeys），再用 defu 覆盖到默认值之上
 *   （source 优先，缺失 / undefined 自动回落默认；未知键已被剔除，不会进入运行态）。
 */
export const normalizeThemeSettings = (input) => {
  // 带 version 字段的信封：只接受当前大版本，否则视为不兼容数据，回落默认。
  if (isObject(input) && 'version' in input && input.version !== THEME_SETTINGS_VERSION) {
    return cloneThemeSettings()
  }

  const source = input?.theme || input?.settings || input

  if (!isThemeSettingsLike(source)) return cloneThemeSettings()

  return defu(pickKnownKeys(source, themeSettings), cloneThemeSettings())
}

export const normalizeThemePreset = (preset) => {
  if (!preset) return null

  const source = preset.theme || preset.settings
  if (!source) return null

  // 丢弃旧预设里的 ElConf/elConf 与原始 settings 键，统一输出规范化的内层 theme。
  const rest = { ...preset }
  delete rest.theme
  delete rest.settings
  delete rest.ElConf
  delete rest.elConf

  return {
    ...rest,
    // version 是预设自身的发布迭代号，沿用作者声明的值，缺省记为 1。
    version: preset.version ?? 1,
    // minMainVersion 才是与 GVA 主题生态的兼容判据；旧预设未声明时兜底为始终兼容。
    minMainVersion: preset.minMainVersion ?? PRESET_FALLBACK_MIN_VERSION,
    theme: normalizeThemeSettings(source)
  }
}

export const buildThemePreset = (settings, name) => ({
  // 用户导出当前配置 = 一个全新预设，发布号从 1 起；
  // minMainVersion 记为导出时的运行版本，即「能用该预设的最低 GVA 版本」。
  version: 1,
  minMainVersion: GVA_MAIN_VERSION,
  name,
  builtin: false,
  theme: liteClone(settings)
})

/** 主题持久化信封：本地缓存与后端 originSetting 共用同一形态，便于版本校验。 */
export const buildThemeEnvelope = (settings) => ({
  version: THEME_SETTINGS_VERSION,
  settings: liteClone(settings)
})

export const readCachedThemeSettings = () => {
  try {
    const raw = localStorage.getItem(THEME_SETTINGS_STORAGE_KEY)
    if (!raw) return null

    const parsed = JSON.parse(raw)
    // 严格按版本信封读取：版本不匹配（含旧的无版本缓存）一律视为失效，交由调用方回落默认。
    if (parsed?.version !== THEME_SETTINGS_VERSION) return null

    return normalizeThemeSettings(parsed.settings)
  } catch {
    return null
  }
}

export const cacheThemeSettings = (settings) => {
  localStorage.setItem(
    THEME_SETTINGS_STORAGE_KEY,
    JSON.stringify(buildThemeEnvelope(settings))
  )
}

export const clearCachedThemeSettings = () => {
  localStorage.removeItem(THEME_SETTINGS_STORAGE_KEY)
}

const getCssVarByTokens = (tokens) => {
  const styles = []

  Object.entries(themeVars).forEach(([group, vars]) => {
    Object.entries(vars).forEach(([tokenKey, tokenValue]) => {
      let cssVarKey = tokenValue.replace('var(', '').replace(')', '')
      let cssValue = tokens[group]?.[tokenKey]

      if (cssValue === undefined) return

      if (group === 'colors') {
        cssVarKey = cssVarKey.replace('rgb(', '').replace(')', '')
        cssValue = colorToRgbChannels(cssValue)
      }

      styles.push(`${cssVarKey}: ${cssValue}`)
    })
  })

  return styles.join(';')
}

export const addThemeVarsToGlobal = (tokens, darkTokens) => {
  const css = `:root { ${getCssVarByTokens(tokens)} }`
  const darkCss = `html.dark { ${getCssVarByTokens(darkTokens)} }`
  const styleId = 'theme-vars'
  const style = document.querySelector(`#${styleId}`) || document.createElement('style')

  style.id = styleId
  style.textContent = css + darkCss
  document.head.appendChild(style)
}

export const toggleCssDarkMode = (isDark) => {
  document.documentElement.classList.toggle('dark', isDark)
  document.documentElement.classList.toggle('light', !isDark)
}

export const toggleAuxiliaryColorModes = (grayscale, colourWeakness) => {
  document.documentElement.classList.toggle('html-grey', grayscale)
  document.documentElement.classList.toggle('html-weakenss', colourWeakness)
}
