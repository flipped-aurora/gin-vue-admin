import { setElementPlusColor, setElementPlusPrimaryColor } from '../color'

export const BASE_ELCONF = {
  zIndex: 2000,
  button: { autoInsertSpace: false },
  message: { 
    grouping: false,
    plain: true,
   }
}

/**
 * 解析注入 `<el-config-provider>` 的配置。
 *
 * 注意：Provider 只能下发组件「行为」配置（size / zIndex / button / message 等），
 * 无法覆盖颜色、圆角这类 CSS 变量——后者由 applyElementPlusTheme 直接写 --el-* 变量。
 * 因此这里不再保留独立的 elConf 覆盖通道，size 直接取自 settings，其余为应用级常量。
 */
export const resolveElementPlusConfig = (settings) => ({
  size: settings.size,
  zIndex: BASE_ELCONF.zIndex,
  button: { ...BASE_ELCONF.button },
  message: { ...BASE_ELCONF.message }
})

/**
 * 把主题色与圆角写入 Element Plus 的 CSS 变量（--el-color-* / 圆角）。
 * @param {{ themeColors: object, settings: object, isDark: boolean }} context
 */
export const applyElementPlusTheme = ({ themeColors: colors, settings, isDark }) => {
  const mode = isDark ? 'dark' : 'light'
  const root = document.documentElement.style
  const radius = `${settings.themeRadius}rem`
  // Prism 把品牌色钉在 500（明暗两态一致），故传入的原始种子色即各语义色的 500 步，
  // 与 UnoCSS 的 bg-primary / bg-success 同源，直接写入 --el-* 即可保持一致；mode 仅用于生成明暗变体。

  setElementPlusPrimaryColor(colors.primary, mode)
  setElementPlusColor('success', colors.success, mode)
  setElementPlusColor('warning', colors.warning, mode)
  setElementPlusColor('danger', colors.error, mode)
  setElementPlusColor('error', colors.error, mode)
  setElementPlusColor('info', colors.info, mode)

  root.setProperty('--gva-radius', radius)
  root.setProperty('--el-border-radius-base', radius)
  root.setProperty('--el-card-border-radius', radius)
}
