/**
 * 通过 <html> class 切换卡片模式。
 * （菜单风格已由 <g-menu> 的语义 token 驱动，不再需要 gva-menu--* 类。）
 * @param {{ settings: object }} context
 */
export const applyStructureTheme = ({ settings }) => {
  const html = document.documentElement

  html.classList.remove('gva-card--border', 'gva-card--shadow')
  html.classList.add(`gva-card--${settings.card.mode}`)
}

