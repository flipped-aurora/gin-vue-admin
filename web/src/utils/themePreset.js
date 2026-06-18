/**
 * 主题预设模块
 * - 内置预设（含 Azir-清新蓝）
 * - 自定义预设 CRUD（localStorage）
 * - 整包序列化 / 解析（带 version，兼容旧版扁平「导入配置」）
 *
 * 预设结构：{ version, name, builtin, config:{…主题字段…} }
 */

const STORAGE_KEY = 'gva-theme-presets'
export const PRESET_VERSION = 1

// 参与预设整包的「可主题化」config 字段（排除纯尺寸/布局模式等无关项）
export const THEME_KEYS = [
  'primaryColor',
  'successColor',
  'warningColor',
  'dangerColor',
  'infoColor',
  'darkMode',
  'menu_theme',
  'radius',
  'card_mode',
  'transition_type',
  'global_size',
  'show_watermark',
  'weakness',
  'grey',
  'show_breadcrumb',
  'show_refresh',
  'show_search',
  'show_collapse_btn',
  'header_bg',
  'header_border',
  'tabs_bg'
]

export const BUILTIN_PRESETS = [
  {
    version: PRESET_VERSION,
    name: 'Azir-清新蓝',
    builtin: true,
    config: {
      primaryColor: '#5D87FF',
      successColor: '#60C041',
      warningColor: '#F9901F',
      dangerColor: '#F56C6C',
      infoColor: '#909399',
      darkMode: 'auto',
      menu_theme: 'design',
      radius: 0.75,
      card_mode: 'border'
    }
  },
  {
    version: PRESET_VERSION,
    name: 'GVA 经典蓝',
    builtin: true,
    config: {
      primaryColor: '#3b82f6',
      successColor: '#67c23a',
      warningColor: '#e6a23c',
      dangerColor: '#f56c6c',
      infoColor: '#909399',
      darkMode: 'auto',
      menu_theme: 'light',
      radius: 0.125,
      card_mode: 'border'
    }
  },
  {
    version: PRESET_VERSION,
    name: '暗夜深色',
    builtin: true,
    config: {
      primaryColor: '#5D87FF',
      successColor: '#60C041',
      warningColor: '#F9901F',
      dangerColor: '#F56C6C',
      infoColor: '#909399',
      darkMode: 'dark',
      menu_theme: 'dark',
      radius: 0.5,
      card_mode: 'shadow'
    }
  }
]

/** 读取自定义预设列表 */
export const loadCustomPresets = () => {
  try {
    const list = JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]')
    return Array.isArray(list) ? list : []
  } catch {
    return []
  }
}

/** 写入自定义预设列表 */
export const saveCustomPresets = (list) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(list))
}

/** 从当前 config 构建一个预设对象 */
export const buildPresetFromConfig = (config, name) => ({
  version: PRESET_VERSION,
  name,
  builtin: false,
  config: THEME_KEYS.reduce((acc, key) => {
    if (key in config) acc[key] = config[key]
    return acc
  }, {})
})

/** 新增自定义预设（同名覆盖） */
export const addCustomPreset = (preset) => {
  const list = loadCustomPresets().filter((p) => p.name !== preset.name)
  list.push(preset)
  saveCustomPresets(list)
  return list
}

/** 删除自定义预设 */
export const removeCustomPreset = (name) => {
  const list = loadCustomPresets().filter((p) => p.name !== name)
  saveCustomPresets(list)
  return list
}

/** 序列化为可下载文本 */
export const serializePreset = (preset) => JSON.stringify(preset, null, 2)

/**
 * 解析导入文本：
 * - 新版整包：{ version, name, config }
 * - 旧版「导入配置」：扁平 config（无 config 包裹、无 version）
 */
export const parsePreset = (text) => {
  const data = JSON.parse(text)
  if (data && typeof data === 'object' && !('config' in data) && !('version' in data)) {
    return {
      version: 0,
      name: '导入配置',
      builtin: false,
      config: data
    }
  }
  return data
}
