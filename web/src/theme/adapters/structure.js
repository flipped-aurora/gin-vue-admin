/**
 * 通过 <html> class 切换菜单风格与卡片模式。
 * @param {{ settings: object }} context
 */
export const applyStructureTheme = ({ settings }) => {
  const html = document.documentElement

  html.classList.remove('gva-menu--design', 'gva-menu--light', 'gva-menu--dark')
  html.classList.add(`gva-menu--${settings.menu.theme}`)
  html.classList.remove('gva-card--border', 'gva-card--shadow')
  html.classList.add(`gva-card--${settings.card.mode}`)
}

