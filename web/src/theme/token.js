import { createRecommendedPalette, createSimplePalette } from './color'

const colorKeys = ['primary', 'info', 'success', 'warning', 'error']

export const getThemeColors = (settings) => ({
  primary: settings.themeColor,
  info: settings.isInfoFollowPrimary
    ? settings.themeColor
    : settings.otherColor.info,
  success: settings.otherColor.success,
  warning: settings.otherColor.warning,
  error: settings.otherColor.error
})

const createPaletteColors = (colors, recommended = false) => {
  const result = {}
  const createPalette = recommended ? createRecommendedPalette : createSimplePalette

  colorKeys.forEach((key) => {
    const palette = createPalette(colors[key])
    result[key] = palette[500]

    Object.entries(palette).forEach(([number, value]) => {
      result[`${key}-${number}`] = value
    })
  })

  return result
}

export const createThemeToken = (colors, tokens, recommended = false) => {
  const paletteColors = createPaletteColors(colors, recommended)

  const lightColors = {
    ...(tokens?.light?.colors || {})
  }

  const lightBoxShadow = {
    ...(tokens?.light?.boxShadow || {})
  }

  const themeTokens = {
    colors: {
      ...paletteColors,
      nprogress: paletteColors.primary,
      ...lightColors
    },
    boxShadow: lightBoxShadow
  }

  const darkThemeTokens = {
    colors: {
      ...themeTokens.colors,
      ...(tokens?.dark?.colors || {})
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
