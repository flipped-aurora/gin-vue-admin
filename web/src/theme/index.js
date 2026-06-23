/**
 * 主题预设模块
 * - 内置预设：从 ./preset/*.json 自动发现（按 order 升序）
 * - 自定义预设 CRUD（localStorage）
 * - 整包序列化 / 解析（带 version，兼容旧版扁平「导入配置」）
 *
 * 预设结构：{ order?, version, name, builtin, config:{…全部系统配置…}, ElConf?:{…} }
 *
 * 两类配置的关系（重要）：
 * 1) config —— 指向 gva-theme 的「系统配置」，即 appStore.config 的全部可持久化字段
 *    （颜色 / 菜单 / 圆角 / 卡片 / 顶栏 / 布局模式 / 侧栏尺寸 / 界面开关 / 全局尺寸…）。
 * 2) ElConf —— 注入 Element Plus 的 <el-config-provider> 的额外配置。
 *
 * 二者存在「重叠字段」（例如：尺寸 size）。重叠字段默认【由 gva-theme 的 config 派生】，
 * 例如 config.global_size = 'small'，则 el-config-provider 的 size 也跟着是 'small'；
 * 但只要 ElConf 里【显式声明】了同名字段（如 ElConf.size = 'large'），
 * 则【ElConf 优先级更高】，覆盖派生值。其余非重叠字段（zIndex / button / message）
 * 只来自 ElConf（缺省时回落 BASE_ELCONF）。该优先级在 resolveElConfig 中实现。
 */

const STORAGE_KEY = 'gva-theme-presets'
export const PRESET_VERSION = 2

/**
 * 参与预设的「系统配置」字段：等于 appStore.config 的全部可持久化键
 * （不含 device / drawerSize / operateMinWith / sideCollapse 等纯运行时态）。
 * 切换预设会整包覆盖这些键——未在预设里声明的，回落到 BASE_CONFIG，保证干净不残留。
 */
export const PRESET_KEYS = [
  // 视觉辅助
  'weakness',
  'grey',
  // 颜色（主色 + 语义色）
  'primaryColor',
  'successColor',
  'warningColor',
  'dangerColor',
  'infoColor',
  // 界面
  'showTabs',
  'darkMode',
  // 侧栏尺寸
  'layout_side_width',
  'layout_side_collapsed_width',
  'layout_side_item_height',
  'show_watermark',
  // 布局模式
  'side_mode',
  'transition_type',
  'global_size',
  // 外观细节
  'menu_theme',
  'radius',
  'card_mode',
  // 顶栏可见性 + 配色 + 阴影
  'show_breadcrumb',
  'show_breadcrumb_icon',
  'show_refresh',
  'show_search',
  'show_collapse_btn',
  'header_bg',
  'header_shadow',
  'tabs_bg'
]

/**
 * 系统配置字段的中性默认值（= gin-vue-admin 出厂默认）。
 * 既作为 appStore.config 的初始值与「重置配置」目标，也作为切换预设时的回落值：
 * 预设未显式声明的键回落到这里，保证「干净覆盖」——例如上一个预设设过的顶栏配色会被清空。
 */
export const BASE_CONFIG = {
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
  transition_type: 'slide',
  global_size: 'default',
  menu_theme: 'design',
  radius: 0.75,
  card_mode: 'border',
  show_breadcrumb: true,
  show_breadcrumb_icon: true,
  show_refresh: true,
  show_search: true,
  show_collapse_btn: true,
  header_bg: '',
  header_shadow: 'sm',
  tabs_bg: ''
}

/**
 * <el-config-provider> 专属字段的默认值（与 Element Plus 自身默认对齐，空 ElConf 即原生行为）。
 * 注意：这里【不含 size】——size 默认由 config.global_size 派生（见 resolveElConfig）。
 */
export const BASE_ELCONF = {
  zIndex: 2000,
  button: { autoInsertSpace: false },
  message: { grouping: false }
}

/**
 * 计算最终注入 <el-config-provider> 的 props。
 * 优先级：ElConf 显式值 > 由 gva-theme config 派生 > BASE_ELCONF 默认。
 * - size：缺省派生自 config.global_size；ElConf.size 存在则覆盖。
 * - zIndex / button / message：来自 ElConf，缺省回落 BASE_ELCONF（button/message 做浅合并）。
 */
export const resolveElConfig = (config, elConf = {}) => ({
  size: elConf.size ?? config.global_size,
  zIndex: elConf.zIndex ?? BASE_ELCONF.zIndex,
  button: { ...BASE_ELCONF.button, ...elConf.button },
  message: { ...BASE_ELCONF.message, ...elConf.message }
})

// 内置预设：自动发现 src/theme/preset 下所有 json，按 order 升序排列
const presetModules = import.meta.glob('./preset/*.json', { eager: true })
export const BUILTIN_PRESETS = Object.values(presetModules)
  .map((m) => m.default || m)
  .sort((a, b) => (a.order ?? 0) - (b.order ?? 0))

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

/** 从当前 config + ElConf 构建一个预设对象（整包：全部系统配置 + el-config-provider 覆盖） */
export const buildPresetFromConfig = (config, name, elConf = {}) => ({
  version: PRESET_VERSION,
  name,
  builtin: false,
  config: PRESET_KEYS.reduce((acc, key) => {
    if (key in config) acc[key] = config[key]
    return acc
  }, {}),
  ElConf: { ...elConf }
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
 * - 新版整包：{ version, name, config, ElConf? }
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
