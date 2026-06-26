import {
  normalizeThemePreset
} from './shared'

const STORAGE_KEY = 'gva-theme-presets'

export {
  buildThemePreset,
  normalizeThemePreset
} from './shared'
export {
  cloneThemeSettings
} from './settings'
export {
  cacheThemeSettings,
  clearCachedThemeSettings,
  readCachedThemeSettings
} from './shared'
export {
  resolveElementPlusConfig
} from './adapters/element-plus'
export {
  GVA_MAIN_VERSION,
  compareMainVersion,
  isPresetCompatible
} from './version'

// Built-in presets are normalized at module boundary so the rest of the app
// consumes a normalized nested `theme` shape.
const presetModules = import.meta.glob('./preset/*.json', { eager: true })
export const BUILTIN_PRESETS = Object.values(presetModules)
  .map((module) => normalizeThemePreset(module.default || module))
  .filter(Boolean)
  .sort((a, b) => (a.order ?? 0) - (b.order ?? 0))

export const loadCustomPresets = () => {
  try {
    const list = JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]')
    return Array.isArray(list) ? list.map(normalizeThemePreset).filter(Boolean) : []
  } catch {
    return []
  }
}

const saveCustomPresets = (list) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(list))
}

export const addCustomPreset = (preset) => {
  const normalized = normalizeThemePreset(preset)
  const list = loadCustomPresets().filter((item) => item.name !== normalized.name)
  list.push(normalized)
  saveCustomPresets(list)
  return list
}

export const removeCustomPreset = (name) => {
  const list = loadCustomPresets().filter((preset) => preset.name !== name)
  saveCustomPresets(list)
  return list
}

export const serializePreset = (preset) => JSON.stringify(normalizeThemePreset(preset), null, 2)

export const parsePreset = (text) => normalizeThemePreset(JSON.parse(text))
