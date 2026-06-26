const clamp = (value, min = 0, max = 255) =>
  Math.min(max, Math.max(min, value))

export const colorToRgb = (color) => {
  if (!color || typeof color !== 'string') return null
  const value = color.trim()

  if (value.startsWith('#')) {
    let hex = value.slice(1)
    if (hex.length === 3) hex = hex.split('').map((c) => c + c).join('')
    if (hex.length !== 6 && hex.length !== 8) return null

    return {
      r: parseInt(hex.slice(0, 2), 16),
      g: parseInt(hex.slice(2, 4), 16),
      b: parseInt(hex.slice(4, 6), 16),
      a: hex.length === 8 ? parseInt(hex.slice(6, 8), 16) / 255 : 1
    }
  }

  const rgb = value.match(/rgba?\(([^)]+)\)/i)
  if (rgb) {
    const parts = rgb[1].split(',').map((item) => parseFloat(item.trim()))
    if (parts.length < 3 || parts.some((item) => Number.isNaN(item))) return null

    return {
      r: parts[0],
      g: parts[1],
      b: parts[2],
      a: parts[3] === undefined ? 1 : parts[3]
    }
  }

  return null
}

export const rgbToHex = (r, g, b) => {
  const parts = [r, g, b].map((item) => {
    const value = clamp(Math.round(item)).toString(16)
    return value.length === 1 ? `0${value}` : value
  })
  return `#${parts.join('')}`
}

export const colorToHex = (color) => {
  const rgb = colorToRgb(color)
  return rgb ? rgbToHex(rgb.r, rgb.g, rgb.b) : color
}

export const colorToRgbChannels = (color) => {
  const rgb = colorToRgb(color)
  return rgb ? `${rgb.r} ${rgb.g} ${rgb.b}` : '0 0 0'
}

export const mixColor = (color, target, amount) => {
  const sourceRgb = colorToRgb(color)
  const targetRgb = colorToRgb(target)
  if (!sourceRgb || !targetRgb) return color

  return rgbToHex(
    sourceRgb.r * (1 - amount) + targetRgb.r * amount,
    sourceRgb.g * (1 - amount) + targetRgb.g * amount,
    sourceRgb.b * (1 - amount) + targetRgb.b * amount
  )
}

export const addOpacityToColor = (color, opacity) => {
  const rgb = colorToRgb(color)
  if (!rgb) return color
  return `rgba(${rgb.r}, ${rgb.g}, ${rgb.b}, ${opacity})`
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

  if (s === 0) {
    const value = Math.round(l * 255)
    return [value, value, value]
  }

  const q = l < 0.5 ? l * (1 + s) : l + s - l * s
  const p = 2 * l - q
  return [
    Math.round(hue2rgb(p, q, h + 1 / 3) * 255),
    Math.round(hue2rgb(p, q, h) * 255),
    Math.round(hue2rgb(p, q, h - 1 / 3) * 255)
  ]
}

export const autoDarkColor = (color) => {
  const rgb = colorToRgb(color)
  if (!rgb) return color

  const [h, s, l] = rgbToHsl(rgb.r, rgb.g, rgb.b)
  const nextLightness = Math.min(0.45, Math.max(0.06, (1 - l) * 0.2 + 0.06))
  const nextSaturation = Math.min(s, 0.6)
  const [r, g, b] = hslToRgb(h, nextSaturation, nextLightness)

  return `rgba(${r}, ${g}, ${b}, ${rgb.a})`
}

const generateDarkColor = (color, amount) => mixColor(color, '#0a0a1e', amount)
const generateLightColor = (color, amount) => mixColor(color, '#f0f8ff', amount)

/**
 * Keep the current GVA/Element Plus palette behavior while moving ownership
 * into the theme package. The adapter can later swap this implementation for
 * a richer 50-950 color palette without touching callers.
 */
export const setElementPlusColor = (name, color, mode) => {
  const createColor = mode === 'light' ? generateLightColor : generateDarkColor
  const root = document.documentElement.style

  root.setProperty(`--el-color-${name}`, color)

  for (let times = 1; times <= 2; times++) {
    root.setProperty(
      `--el-color-${name}-dark-${times}`,
      createColor(color, times / 10)
    )
  }

  for (let times = 1; times <= 10; times++) {
    root.setProperty(
      `--el-color-${name}-light-${times}`,
      createColor(color, times / 10)
    )
  }
}

export const setElementPlusPrimaryColor = (primaryColor, mode) => {
  setElementPlusColor('primary', primaryColor, mode)
  const root = document.documentElement.style
  root.setProperty('--el-color-primary-bg', addOpacityToColor(primaryColor, 0.4))
  root.setProperty('--el-menu-hover-bg-color', addOpacityToColor(primaryColor, 0.2))
}

export const createSimplePalette = (color) => ({
  50: mixColor(color, '#ffffff', 0.92),
  100: mixColor(color, '#ffffff', 0.84),
  200: mixColor(color, '#ffffff', 0.72),
  300: mixColor(color, '#ffffff', 0.56),
  400: mixColor(color, '#ffffff', 0.32),
  500: color,
  600: mixColor(color, '#000000', 0.12),
  700: mixColor(color, '#000000', 0.24),
  800: mixColor(color, '#000000', 0.36),
  900: mixColor(color, '#000000', 0.48),
  950: mixColor(color, '#000000', 0.6)
})

const RECOMMENDED_PALETTE_MIX = {
  50: ['#ffffff', 0.95],
  100: ['#ffffff', 0.9],
  200: ['#ffffff', 0.78],
  300: ['#ffffff', 0.64],
  400: ['#ffffff', 0.38],
  500: [null, 0],
  600: ['#000000', 0.08],
  700: ['#000000', 0.18],
  800: ['#000000', 0.3],
  900: ['#000000', 0.42],
  950: ['#000000', 0.56]
}

export const createRecommendedPalette = (color) => {
  return Object.entries(RECOMMENDED_PALETTE_MIX).reduce((palette, [level, [target, amount]]) => {
    palette[level] = target ? mixColor(color, target, amount) : color
    return palette
  }, {})
}
