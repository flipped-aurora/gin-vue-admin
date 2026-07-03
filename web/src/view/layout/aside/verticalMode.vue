<template>
  <div
    class="fixed left-0 top-0 z-30 flex h-screen flex-col shadow-sider transition-[width] duration-300 ease-in-out"
    :style="{ width: sideWidth + 'px' }"
    :class="[surfaceClass, siderDarkClass]"
  >
    <!-- Logo -->
    <div
      class="flex h-16 flex-shrink-0 cursor-pointer items-center justify-center gap-2 overflow-hidden px-2"
      @click="router.push({ path: '/' })"
    >
      <Logo :dark="sidebarDark" />
      <span v-if="!sideCollapse" class="whitespace-nowrap text-xl font-bold">
        {{ $GIN_VUE_ADMIN.appName }}
      </span>
    </div>

    <!-- 菜单 -->
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

    <!-- 折叠按钮 -->
    <collapse-bar />
  </div>
</template>

<script setup>
  import Logo from '@/components/logo/index.vue'
  import { computed } from 'vue'
  import { useRouter } from 'vue-router'
  import { storeToRefs } from 'pinia'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { useMenuActive, useMenuNavigation } from './composables/useMenu'
  import { useSidebarTheme } from './composables/useSidebarTheme'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import CollapseBar from './CollapseBar.vue'

  defineOptions({ name: 'VerticalMode' })

  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { sideCollapse } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
  const router = useRouter()
  const routerStore = useRouterStore()

  const menus = computed(() => routerStore.rootMenus)
  const { menuTheme, surfaceClass, siderDarkClass, sidebarDark } = useSidebarTheme()
  const { activeKey, openKeys } = useMenuActive(menus, menuTheme)
  const { navigate } = useMenuNavigation()
  const { sideWidth } = useSideWidth()
</script>
