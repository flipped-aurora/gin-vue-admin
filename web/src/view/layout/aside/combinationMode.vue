<template>
  <div class="h-full">
    <!-- head：顶栏一级横向菜单 -->
    <div
      v-if="mode === 'head'"
      class="mx-2 flex h-[calc(100%-4px)] w-[calc(100vw-600px)] items-center text-base-text"
    >
      <g-menu
        orientation="horizontal"
        :items="routerStore.topMenu"
        :active="topActiveKey"
        class="w-full"
        @select="onTopSelect"
      />
    </div>

    <!-- normal：左侧二级纵向菜单（当前一级无二级时整块不渲染） -->
    <div
      v-if="mode === 'normal' && showLeftAside"
      class="relative flex h-full flex-col shadow-sider transition-[width] duration-300 ease-in-out"
      :class="[surfaceClass, siderDarkClass]"
      :style="{ width: layoutSideWidth + 'px' }"
    >
      <el-scrollbar class="flex-1">
        <g-menu
          :items="leftMenus"
          :theme="menuTheme"
          :collapsed="sideCollapse"
          :active="activeKey"
          v-model:open-keys="openKeys"
          :item-height="settings.layout.sideItemHeight"
          @select="navigate"
        />
      </el-scrollbar>
      <collapse-bar />
    </div>
  </div>
</template>

<script setup>
  import { computed } from 'vue'
  import { useRoute } from 'vue-router'
  import { storeToRefs } from 'pinia'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { resolveActiveTop, visibleItems } from '@/core/componentLibrary/menu/shared'
  import { useMenuActive, useMenuNavigation } from './composables/useMenu'
  import { useSidebarTheme } from './composables/useSidebarTheme'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import CollapseBar from './CollapseBar.vue'

  defineOptions({ name: 'AsideCombinationMode' })

  defineProps({
    mode: { type: String, default: 'normal' }
  })

  const { sideCollapse } = storeToRefs(useAppStore())
  const { settings } = storeToRefs(useThemeStore())
  const route = useRoute()
  const routerStore = useRouterStore()

  // 当前路由所属的一级菜单：左栏内容与顶栏高亮都从它派生，
  // 直接读 asyncRouters（rootMenus），刷新 / 直达 / 跨一级切换都正确。
  const activeName = computed(() => route.meta.activeName || route.name)
  const activeTop = computed(() =>
    resolveActiveTop(routerStore.rootMenus, activeName.value)
  )
  const topActiveKey = computed(() => activeTop.value?.name || '')
  const leftMenus = computed(() => visibleItems(activeTop.value?.children))
  const showLeftAside = computed(() => leftMenus.value.length > 0)

  const { menuTheme, surfaceClass, siderDarkClass } = useSidebarTheme()
  const { activeKey, openKeys } = useMenuActive(leftMenus, menuTheme)
  const { navigate } = useMenuNavigation()
  const { sideWidth: layoutSideWidth } = useSideWidth()

  const isHttp = (s) =>
    !!s && (s.indexOf('http://') > -1 || s.indexOf('https://') > -1)

  // 顶栏一级点击：外链新开；否则跳到该一级下第一个可见子项（无子项则跳自身）
  const onTopSelect = (index) => {
    if (isHttp(index)) {
      window.open(index, '_blank')
      return
    }
    const top = routerStore.rootMenus.find((i) => i.name === index)
    const first = visibleItems(top?.children).find((c) => !isHttp(c.path))
    navigate(first ? first.name : index)
  }
</script>
