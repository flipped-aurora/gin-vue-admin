<template>
  <div
    class="relative z-10 flex h-full flex-col shadow-sider transition-[width] duration-300 ease-in-out"
    :class="[surfaceClass, siderDarkClass]"
    :style="{ width: layoutSideWidth + 'px' }"
  >
    <el-scrollbar class="flex-1">
      <g-menu
        :items="menus"
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
</template>

<script setup>
  import { computed } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { useMenuActive, useMenuNavigation } from './composables/useMenu'
  import { useSidebarTheme } from './composables/useSidebarTheme'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import CollapseBar from './CollapseBar.vue'

  defineOptions({ name: 'AsideNormalMode' })

  const { sideCollapse } = storeToRefs(useAppStore())
  const { settings } = storeToRefs(useThemeStore())
  const routerStore = useRouterStore()

  const menus = computed(() => routerStore.rootMenus)
  const { menuTheme, surfaceClass, siderDarkClass } = useSidebarTheme()
  const { activeKey, openKeys } = useMenuActive(menus, menuTheme)
  const { navigate } = useMenuNavigation()
  const { sideWidth: layoutSideWidth } = useSideWidth()
</script>
