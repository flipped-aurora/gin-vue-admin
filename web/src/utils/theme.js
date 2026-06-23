import { setBodyPrimaryColor, setElColor } from '@/utils/format'

/**
 * 主题落地层：把 appStore.config 中的主题字段应用到 DOM
 * - 颜色：主色 + 语义色（含 light-1..10 / dark-1..2 阶）
 * - 圆角：--gva-radius 及 Element Plus 圆角变量
 * - 菜单风格 / 卡片盒子模式：通过 <html> 上的类切换
 */

const root = () => document.documentElement

/* ---------- 颜色工具：解析 / HSL 互转 / 暗色自动推导 ---------- */

const parseColor = (str) => {
  if (!str || typeof str !== 'string') return null
  const s = str.trim()
  if (s[0] === '#') {
    let hex = s.slice(1)
    if (hex.length === 3) hex = hex.split('').map((c) => c + c).join('')
    if (hex.length !== 6 && hex.length !== 8) return null
    return {
      r: parseInt(hex.slice(0, 2), 16),
      g: parseInt(hex.slice(2, 4), 16),
      b: parseInt(hex.slice(4, 6), 16),
      a: hex.length === 8 ? parseInt(hex.slice(6, 8), 16) / 255 : 1
    }
  }
  const m = s.match(/rgba?\(([^)]+)\)/i)
  if (m) {
    const p = m[1].split(',').map((x) => parseFloat(x.trim()))
    return { r: p[0], g: p[1], b: p[2], a: p[3] === undefined ? 1 : p[3] }
  }
  return null
}

const rgbToHsl = (r, g, b) => {
  r /= 255
  g /= 255
  b /= 255
  const max = Math.max(r, g, b)
  const min = Math.min(r, g, b)
  const l = (max + min) / 2
  let h = 0
  let s = 0
  if (max !== min) {
    const d = max - min
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min)
    if (max === r) h = (g - b) / d + (g < b ? 6 : 0)
    else if (max === g) h = (b - r) / d + 2
    else h = (r - g) / d + 4
    h *= 60
  }
  return [h, s, l]
}

const hslToRgb = (h, s, l) => {
  h /= 360
  const hue2rgb = (p, q, t) => {
    if (t < 0) t += 1
    if (t > 1) t -= 1
    if (t < 1 / 6) return p + (q - p) * 6 * t
    if (t < 1 / 2) return q
    if (t < 2 / 3) return p + (q - p) * (2 / 3 - t) * 6
    return p
  }
  let r
  let g
  let b
  if (s === 0) {
    r = g = b = l
  } else {
    const q = l < 0.5 ? l * (1 + s) : l + s - l * s
    const p = 2 * l - q
    r = hue2rgb(p, q, h + 1 / 3)
    g = hue2rgb(p, q, h)
    b = hue2rgb(p, q, h - 1 / 3)
  }
  return [Math.round(r * 255), Math.round(g * 255), Math.round(b * 255)]
}

/**
 * 暗色自动推导：保留色相，按「表面色在明暗模式间翻转明度」的规律，
 * 把用户为浅色模式选的颜色映射为合适的深色（浅 → 深）。
 */
export const autoDarkColor = (str) => {
  const c = parseColor(str)
  if (!c) return str
  const [h, s, l] = rgbToHsl(c.r, c.g, c.b)
  const nl = Math.min(0.45, Math.max(0.06, (1 - l) * 0.2 + 0.06))
  const ns = Math.min(s, 0.6)
  const [r, g, b] = hslToRgb(h, ns, nl)
  return `rgba(${r}, ${g}, ${b}, ${c.a})`
}

/** 主色 + 语义色 */
export const applyColors = (config, isDark) => {
  const mode = isDark ? 'dark' : 'light'
  setBodyPrimaryColor(config.primaryColor, mode)
  setElColor('success', config.successColor, mode)
  setElColor('warning', config.warningColor, mode)
  setElColor('danger', config.dangerColor, mode)
  // Element Plus 同时使用 error 作为危险色别名
  setElColor('error', config.dangerColor, mode)
  setElColor('info', config.infoColor, mode)
}

/** 全局圆角（rem） */
export const applyRadius = (radius) => {
  const r = `${radius}rem`
  const s = root().style
  s.setProperty('--gva-radius', r)
  s.setProperty('--el-border-radius-base', r)
  s.setProperty('--el-card-border-radius', r)
}

/** 菜单风格 'design' | 'light' | 'dark' */
export const applyMenuTheme = (theme) => {
  const el = root()
  el.classList.remove('gva-menu--design', 'gva-menu--light', 'gva-menu--dark')
  el.classList.add(`gva-menu--${theme}`)
}

/** 卡片盒子模式 'border' | 'shadow' */
export const applyCardMode = (mode) => {
  const el = root()
  el.classList.remove('gva-card--border', 'gva-card--shadow')
  el.classList.add(`gva-card--${mode}`)
}

/**
 * 顶栏阴影尺寸档位 → box-shadow 预设值（明暗各一套）。
 * none=无；sm/md/lg 依次增强。暗色背景需更强阴影才可见，故单独定义，
 * 与顶栏 / 标签栏配色的「按明暗分治」保持一致。
 */
const HEADER_SHADOWS = {
  light: {
    none: 'none',
    sm: '0 1px 0 rgba(0, 0, 0, 0.06), 0 1px 3px rgba(0, 0, 0, 0.04)',
    md: '0 1px 0 rgba(0, 0, 0, 0.06), 0 4px 12px rgba(0, 0, 0, 0.08)',
    lg: '0 1px 0 rgba(0, 0, 0, 0.06), 0 8px 24px rgba(0, 0, 0, 0.12)'
  },
  dark: {
    none: 'none',
    sm: '0 1px 0 rgba(0, 0, 0, 0.4), 0 1px 3px rgba(0, 0, 0, 0.35)',
    md: '0 1px 0 rgba(0, 0, 0, 0.4), 0 4px 12px rgba(0, 0, 0, 0.45)',
    lg: '0 1px 0 rgba(0, 0, 0, 0.4), 0 8px 24px rgba(0, 0, 0, 0.55)'
  }
}

/**
 * 顶栏 / 标签栏配色 + 顶栏阴影：
 * - 配色（header_bg / tabs_bg）：空值时移除内联变量，回退 theme.scss 默认；有值时覆盖，
 *   暗色模式下对用户所选（浅色）颜色自动推导深色版本。
 * - 阴影（header_shadow）：按尺寸档位取预设 box-shadow（明暗各一套）；
 *   未知档位移除内联变量，回退 theme.scss 默认。
 */
export const applyChrome = (config, isDark) => {
  const s = root().style
  const setOrClear = (name, val) => {
    if (val) s.setProperty(name, isDark ? autoDarkColor(val) : val)
    else s.removeProperty(name)
  }
  setOrClear('--gva-header-bg', config.header_bg)
  setOrClear('--gva-tabs-bg', config.tabs_bg)
  const shadow = HEADER_SHADOWS[isDark ? 'dark' : 'light'][config.header_shadow]
  if (shadow !== undefined) s.setProperty('--gva-header-shadow', shadow)
  else s.removeProperty('--gva-header-shadow')
}

/** 一次性应用全部主题 */
export const applyTheme = (config, isDark) => {
  applyColors(config, isDark)
  applyRadius(config.radius)
  applyMenuTheme(config.menu_theme)
  applyCardMode(config.card_mode)
  applyChrome(config, isDark)
}
