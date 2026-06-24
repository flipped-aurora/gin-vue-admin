import { autoDarkColor } from '../color'

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

const setOrClear = (style, key, value) => {
  if (value === undefined || value === null || value === '') style.removeProperty(key)
  else style.setProperty(key, value)
}

/**
 * 应用顶栏 / 标签栏的背景与阴影到根节点 CSS 变量。
 * @param {{ settings: object, isDark: boolean }} context
 */
export const applyChromeTheme = ({ settings, isDark }) => {
  const style = document.documentElement.style
  const mode = isDark ? 'dark' : 'light'
  const headerBg = isDark ? autoDarkColor(settings.header.bg) : settings.header.bg
  const tabBg = isDark ? autoDarkColor(settings.tab.bg) : settings.tab.bg

  setOrClear(style, '--gva-header-bg', headerBg)
  setOrClear(style, '--gva-tabs-bg', tabBg)
  setOrClear(style, '--gva-header-shadow', HEADER_SHADOWS[mode][settings.header.shadow])
}

