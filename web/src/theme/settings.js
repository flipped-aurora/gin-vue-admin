/**
 * Centralized theme settings.
 *
 * settings is the only durable theme source. Component libraries and runtime
 * CSS variables are derived from it through adapters.
 */

/**
 * 主题 settings 的 schema 版本。本次为大版本重构，不兼容旧的扁平主题数据。
 * 持久化（本地缓存 / 后端 originSetting）统一写成 `{ version, settings }` 信封，
 * 读取时只接受当前版本，版本不匹配或旧格式一律回落默认值。
 * 未来 settings 结构再做破坏性调整时，递增此版本即可让旧数据自动失效。
 */
export const THEME_SETTINGS_VERSION = 1

export const themeSettings = {
  themeScheme: 'auto',
  grayscale: false,
  colourWeakness: false,
  recommendColor: false,
  themeColor: '#5d87ff',
  themeRadius: 0.5,
  size: 'default',
  otherColor: {
    info: '#909399',
    success: '#60c041',
    warning: '#f9901f',
    error: '#f56c6c'
  },
  isInfoFollowPrimary: false,
  layout: {
    mode: 'normal',
    sideWidth: 256,
    sideCollapsedWidth: 80,
    sideItemHeight: 48
  },
  page: {
    transition: 'slide'
  },
  header: {
    breadcrumb: {
      visible: true,
      showIcon: true
    },
    refresh: {
      visible: true
    },
    search: {
      visible: true
    },
    collapseButton: {
      visible: true
    },
    bg: '',
    shadow: 'sm',
  },
  tab: {
    visible: true,
    bg: '',
    // 标签栏阴影档位：none / sm / md / lg（与顶栏阴影同机制）
    shadow: 'sm',
    // 标签栏风格：button(描边胶囊) / chrome(浏览器标签) / slider(底部指示条)
    mode: 'button'
  },
  menu: {
    theme: 'design',
    darkSider: false
  },
  card: {
    mode: 'border'
  },
  watermark: {
    visible: false
  },
  tokens: {
    light: {
      colors: {
        container: 'rgb(255, 255, 255)',
        layout: 'rgb(247, 250, 252)',
        inverted: 'rgb(15, 23, 42)',
        'base-text': 'rgb(31, 41, 55)',
        border: 'rgb(229, 231, 235)',
        muted: 'rgb(248, 250, 252)',
        'muted-foreground': 'rgb(100, 116, 139)',
        'control-track': 'rgb(209, 213, 219)'
      },
      boxShadow: {
        header: '0 1px 0 rgba(0, 0, 0, 0.06), 0 1px 3px rgba(0, 0, 0, 0.04)',
        sider: '2px 0 8px 0 rgba(29, 35, 41, 0.05)',
        tab: '0 1px 2px rgba(0, 21, 41, 0.08)',
        card: '0 1px 2px rgba(0, 0, 0, 0.04), 0 6px 18px rgba(0, 0, 0, 0.06)'
      }
    },
    dark: {
      colors: {
        container: 'rgb(15, 23, 42)',
        layout: 'rgb(30, 41, 59)',
        inverted: 'rgb(255, 255, 255)',
        'base-text': 'rgb(226, 232, 240)',
        border: 'rgb(51, 65, 85)',
        muted: 'rgb(30, 41, 59)',
        'muted-foreground': 'rgb(148, 163, 184)',
        'control-track': 'rgb(71, 85, 105)'
      },
      boxShadow: {
        header: '0 1px 0 rgba(0, 0, 0, 0.4), 0 1px 3px rgba(0, 0, 0, 0.35)',
        sider: '2px 0 8px 0 rgba(0, 0, 0, 0.35)',
        tab: '0 1px 0 rgba(0, 0, 0, 0.4), 0 1px 3px rgba(0, 0, 0, 0.35)',
        card: '0 1px 2px rgba(0, 0, 0, 0.28), 0 6px 18px rgba(0, 0, 0, 0.32)'
      }
    }
  }
}

export const cloneThemeSettings = () => structuredClone(themeSettings)
