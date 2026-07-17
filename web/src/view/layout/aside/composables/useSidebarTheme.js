import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'
import { SIDEBAR_SURFACE } from '@/core/componentLibrary/menu/variants'

/**
 * 侧栏深浅的单一事实源：供 5 个布局 mode 与移动抽屉共用。
 * - 解析规则：主题模式优先，浅色（含跟随系统解析为浅色）时才听「深色侧边栏」开关。
 * - 上色机制：浅色 + darkSider → 给侧栏容器加 .gva-sider-dark（作用域 token 翻转），
 *   全局暗色时 html.dark 已生效，无需再叠加；两条路径共用同一套 dark token，观感一致。
 * @returns {{
 *   sidebarDark: import('vue').ComputedRef<boolean>,
 *   siderDarkClass: import('vue').ComputedRef<string>,
 *   menuTheme: import('vue').ComputedRef<string>,
 *   surfaceClass: string
 * }}
 */
export function useSidebarTheme() {
  const themeStore = useThemeStore()
  const { settings, darkMode } = storeToRefs(themeStore)

  // 侧栏是否呈深色（供 logo 等消费）：全局暗色优先，其次听开关
  const sidebarDark = computed(
    () => darkMode.value || settings.value.menu.darkSider
  )
  // 仅「浅色主题 + 开启深色侧边栏」时叠作用域暗色类；全局暗色时留空（html.dark 已覆盖）
  const siderDarkClass = computed(() =>
    !darkMode.value && settings.value.menu.darkSider ? 'gva-sider-dark' : ''
  )
  // 传给 <g-menu> 的风格始终是用户所选（design/light/group），深浅交给 token
  const menuTheme = computed(() => settings.value.menu.theme)

  return { sidebarDark, siderDarkClass, menuTheme, surfaceClass: SIDEBAR_SURFACE }
}
