import { generateTheme } from '@simple-prism/core'

const colorKeys = ['primary', 'info', 'success', 'warning', 'error']
const paletteNumbers = [50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950]

export const getThemeColors = (settings) => ({
  primary: settings.themeColor,
  info: settings.isInfoFollowPrimary
    ? settings.themeColor
    : settings.otherColor.info,
  success: settings.otherColor.success,
  warning: settings.otherColor.warning,
  error: settings.otherColor.error
})

// 全部颜色交给 Prism（@simple-prism/core）生成：
// - 语义色板 primary/info/success/warning/error 各取 11 阶；
// - 中性/表面色 container/layout/border/muted/base-text 等取自 neutral 色阶（见 pickSurface）。
// 唯一的字面色是亮色 container 的纯白纸面——neutral 最浅阶仍偏灰，而抬升表面（卡片/侧栏/浮层）需纯白。
// generateTheme 一次算好 light+dark，这里只取 hex 喂回既有 token 管线，由 shared.js 统一
// 转成 `--x-color: r g b` 裸通道。
//
// dispatchTheme 可能因布局等非颜色设置变化被多次触发，故按 5 个基色的组合缓存，避免同色重复求解。
let paletteCache = { key: '', value: null }

// 中性/表面 token 取自 Prism neutral 色阶。neutral 是镜像色阶，同一阶号在明暗两态各给出正确色。
// 抬升表面 container 与页面画布 layout 的关系并非简单镜像：亮色下卡片是纯白、页面是最浅灰；
// 暗色下卡片是最深面、页面抬升一阶——故这两者按模式取值，其余明暗共用同一阶。
const pickSurface = (neutral, isLight) => {
  const at = (step) => neutral.steps[step].hex
  return {
    container: isLight ? '#ffffff' : at(50), // 卡片/侧栏/浮层：亮=纸白，暗=最深面
    layout: at(isLight ? 50 : 100), // 页面画布：位于容器之下一层
    muted: at(100), // 弱化区块底色
    border: at(200), // 描边
    'control-track': at(300), // 开关/滑块未激活轨道
    'base-text': neutral.textContrast.hex, // 主文本（APCA 高对比求解）
    'muted-foreground': neutral.text.hex // 次文本（APCA 低对比求解）
  }
}

const buildPalettes = (colors) => {
  const key = colorKeys.map((name) => colors[name]).join('|')
  if (paletteCache.key === key && paletteCache.value) return paletteCache.value

  const theme = generateTheme({
    primary: colors.primary,
    info: colors.info,
    success: colors.success,
    warning: colors.warning,
    error: colors.error
  })

  const pick = (mode) => {
    const result = {}
    colorKeys.forEach((name) => {
      const scale = theme.scales[name][mode]
      result[name] = scale.steps[500].hex
      paletteNumbers.forEach((number) => {
        result[`${name}-${number}`] = scale.steps[number].hex
      })
    })
    Object.assign(result, pickSurface(theme.scales.neutral[mode], mode === 'light'))
    return result
  }

  const value = { light: pick('light'), dark: pick('dark') }
  paletteCache = { key, value }
  return value
}

export const createThemeToken = (colors, tokens) => {
  const palettes = buildPalettes(colors)

  // 颜色（语义色板 + 中性/表面）已由 Prism 明暗两套色阶完整给出，直接整体铺开即可；
  // settings.tokens 现在只提供与颜色无关的阴影档位。
  const themeTokens = {
    colors: {
      ...palettes.light,
      nprogress: palettes.light.primary
    },
    boxShadow: { ...(tokens?.light?.boxShadow || {}) }
  }

  const darkThemeTokens = {
    colors: {
      ...palettes.dark,
      nprogress: palettes.dark.primary
    },
    boxShadow: {
      ...themeTokens.boxShadow,
      ...(tokens?.dark?.boxShadow || {})
    }
  }

  return {
    themeTokens,
    darkThemeTokens
  }
}
