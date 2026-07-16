/**
 * 运行时主题派发层。
 *
 * themeStore 算出主题上下文后，统一派发给下列「运行时适配器」。每个适配器接收
 * 同一份 context（{ settings, themeColors, isDark }），各自把主题落到对应目标：
 * - applyCssVarsTheme     -> 注入 :root / html.dark 的 CSS 变量（供 UnoCSS / 业务样式消费）
 * - applyElementPlusTheme -> 写 Element Plus 的 --el-* 变量与圆角
 * - applyChromeTheme      -> 顶栏 / 标签栏背景与阴影
 * - applyStructureTheme   -> 菜单风格 / 卡片模式的 <html> class
 *
 * 接入新的 UI 体系（如 shadcn-vue）时，只需新增一个 `(context) => void` 适配器
 * 并追加到 themeRuntimeAdapters，无需改动 themeStore。
 */
import { addThemeVarsToGlobal } from '../shared'
import { createThemeToken } from '../token'
import { applyChromeTheme } from './chrome'
import { applyElementPlusTheme } from './element-plus'
import { applyStructureTheme } from './structure'

const applyCssVarsTheme = ({ settings, themeColors }) => {
  const { themeTokens, darkThemeTokens } = createThemeToken(themeColors, settings.tokens)
  addThemeVarsToGlobal(themeTokens, darkThemeTokens)
}

export const themeRuntimeAdapters = [
  applyCssVarsTheme,
  applyElementPlusTheme,
  applyChromeTheme,
  applyStructureTheme
]

/**
 * 把主题上下文派发给全部运行时适配器。
 * @param {{ settings: object, themeColors: object, isDark: boolean }} context
 */
export const dispatchTheme = (context) => {
  themeRuntimeAdapters.forEach((adapter) => adapter(context))
}
