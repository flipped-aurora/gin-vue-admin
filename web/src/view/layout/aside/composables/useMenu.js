import { ref, watch, watchEffect } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { findMenuPath, visibleItems } from '@/core/componentLibrary/menu/shared'

/**
 * 菜单导航：parameters 预置、外链新开、同名短路。
 * @returns {{ navigate: (key: string) => void }}
 */
export function useMenuNavigation() {
  const route = useRoute()
  const router = useRouter()
  const routerStore = useRouterStore()

  const navigate = (key) => {
    if (!key || key === route.name) return
    if (key.indexOf('http://') > -1 || key.indexOf('https://') > -1) {
      window.open(key, '_blank')
      return
    }
    const query = {}
    const params = {}
    routerStore.routeMap[key]?.parameters?.forEach((item) => {
      if (item.type === 'query') query[item.key] = item.value
      else params[item.key] = item.value
    })
    router.push({ name: key, query, params })
  }

  return { navigate }
}

/**
 * 从路由推导激活键 + 展开键。
 * - activeKey：iframe 用 query.url，否则 meta.activeName || route.name。
 * - openKeys：非 group = 激活项的祖先路径（unique-opened，随路由更新）；
 *   group = 默认展开所有含可见子级的一级分组，仅在切到 group / 菜单变化时重置，
 *   路由跳转不重置，从而保留用户手动收起的分组。
 * @param {import('vue').Ref<Array>|(() => Array)} menus 菜单树（asyncRouters[0].children）
 * @param {import('vue').Ref<string>} [theme] 当前菜单风格
 */
export function useMenuActive(menus, theme) {
  const route = useRoute()
  const activeKey = ref('')
  const openKeys = ref([])
  const readMenus = () => (typeof menus === 'function' ? menus() : menus.value) || []

  watchEffect(() => {
    if (route.name === 'gvaLayoutIframe') {
      activeKey.value = decodeURIComponent(route.query.url)
      return
    }
    activeKey.value = route.meta.activeName || route.name
  })

  // 非 group：openKeys = 激活项祖先路径，随路由 / 菜单变化更新
  watch(
    [activeKey, () => theme?.value, () => readMenus()],
    () => {
      if (theme?.value === 'group') return
      openKeys.value = findMenuPath(readMenus(), activeKey.value)
    },
    { immediate: true }
  )

  // group：默认展开所有含可见子级的一级分组；不随 activeKey 重置，保留手动收起态
  watch(
    [() => theme?.value, () => readMenus()],
    () => {
      if (theme?.value !== 'group') return
      openKeys.value = visibleItems(readMenus())
        .filter((m) => visibleItems(m.children).length > 0)
        .map((m) => m.name)
    },
    { immediate: true }
  )

  return { activeKey, openKeys }
}
