<template>
  <el-drawer
    v-model="open"
    direction="ltr"
    :size="sideWidth"
    :with-header="false"
    class="gva-mobile-menu-drawer"
  >
    <div class="flex h-full flex-col" :class="[surfaceClass, siderDarkClass]">
      <!-- Logo -->
      <div
        class="flex h-16 flex-shrink-0 cursor-pointer items-center gap-2 px-4"
        @click="goHome"
      >
        <Logo :dark="sidebarDark" />
        <span class="whitespace-nowrap text-xl font-bold">
          {{ $GIN_VUE_ADMIN.appName }}
        </span>
      </div>
      <!-- 菜单 -->
      <el-scrollbar class="flex-1">
        <g-menu
          :items="menus"
          :theme="menuTheme"
          :active="activeKey"
          v-model:open-keys="openKeys"
          :item-height="settings.layout.sideItemHeight"
          @select="onSelect"
        />
      </el-scrollbar>
    </div>
  </el-drawer>
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

  defineOptions({ name: 'MobileMenuDrawer' })

  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { mobileMenuOpen: open } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
  const router = useRouter()
  const routerStore = useRouterStore()

  const menus = computed(() => routerStore.rootMenus)
  const { menuTheme, surfaceClass, siderDarkClass, sidebarDark } = useSidebarTheme()
  const { activeKey, openKeys } = useMenuActive(menus, menuTheme)
  const { navigate } = useMenuNavigation()

  const sideWidth = computed(() => settings.value.layout.sideWidth + 'px')

  const onSelect = (key) => {
    navigate(key)
    appStore.toggleMobileMenu(false)
  }
  const goHome = () => {
    router.push({ path: '/' })
    appStore.toggleMobileMenu(false)
  }
</script>

<!-- 非 scoped：抽屉 body 被 teleport 到 body 外层，scoped :deep 触达不到，用类名限定的全局样式 -->
<style>
  .gva-mobile-menu-drawer .el-drawer__body {
    padding: 0;
  }
</style>
