const colorNames = ['primary', 'info', 'success', 'warning', 'error']
const paletteNumbers = [50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950]

const createPaletteVars = () => {
  const colors = {}

  colorNames.forEach((name) => {
    colors[name] = `rgb(var(--${name}-color))`
    paletteNumbers.forEach((number) => {
      colors[`${name}-${number}`] = `rgb(var(--${name}-${number}-color))`
    })
  })

  return colors
}

export const themeVars = {
  colors: {
    ...createPaletteVars(),
    nprogress: 'rgb(var(--nprogress-color))',
    container: 'rgb(var(--container-bg-color))',
    layout: 'rgb(var(--layout-bg-color))',
    inverted: 'rgb(var(--inverted-bg-color))',
    'base-text': 'rgb(var(--base-text-color))',
    border: 'rgb(var(--border-color))',
    muted: 'rgb(var(--muted-color))',
    'muted-foreground': 'rgb(var(--muted-foreground-color))'
  },
  boxShadow: {
    header: 'var(--header-box-shadow)',
    sider: 'var(--sider-box-shadow)',
    tab: 'var(--tab-box-shadow)',
    card: 'var(--card-box-shadow)'
  }
}

