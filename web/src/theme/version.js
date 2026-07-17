import { version as pkgVersion } from '../../package.json'

/**
 * 当前 GVA 主版本号，取自 package.json，作为「主题生态」唯一事实来源。
 *
 * 注意区分两类版本：
 * - 预设自身的 `version`：预设的发布迭代号，由预设作者维护，同一 GVA 版本下可反复
 *   发布到几十、几百，与主题生态无关，仅用于展示 / 去重。
 * - 预设的 `minMainVersion`：该预设适配的最低 GVA 主版本，用来判断它能否被当前
 *   GVA 应用——这才是跨版本兼容的判据。
 */
export const GVA_MAIN_VERSION = pkgVersion

/** 旧预设未声明 minMainVersion 时的兜底：视为始终兼容，避免误伤历史导出文件。 */
export const PRESET_FALLBACK_MIN_VERSION = '0.0.0'

const toSegments = (value) =>
  String(value ?? '')
    .split('.')
    .map((part) => Number.parseInt(part, 10) || 0)

/** 语义化主版本比较：a < b → -1，a === b → 0，a > b → 1。缺位段按 0 处理。 */
export const compareMainVersion = (a, b) => {
  const left = toSegments(a)
  const right = toSegments(b)
  const len = Math.max(left.length, right.length)

  for (let i = 0; i < len; i++) {
    const diff = (left[i] || 0) - (right[i] || 0)
    if (diff !== 0) return diff > 0 ? 1 : -1
  }

  return 0
}

/** 预设是否适配当前 GVA：当前主版本 >= 预设要求的最低主版本。 */
export const isPresetCompatible = (preset) => {
  const required = preset?.minMainVersion || PRESET_FALLBACK_MIN_VERSION
  return compareMainVersion(GVA_MAIN_VERSION, required) >= 0
}
