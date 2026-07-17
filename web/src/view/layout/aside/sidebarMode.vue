<template>
  <div class="flex h-full">
    <!-- 一级图标常驻栏（与二级之间的分隔按卡片样式：border / shadow） -->
    <div
      class="relative h-full"
      :class="[surfaceClass, siderDarkClass, railSepClass]"
      :style="{ width: settings.layout.sideCollapsedWidth + 'px' }"
    >
      <el-scrollbar>
        <nav class="flex flex-col gap-1 p-2">
          <button
            v-for="item in topLevelMenus"
            :key="item.name"
            type="button"
            :class="cn(
              MENU_ICON_BUTTON,
              FOCUS_RING,
              topActive === item.name && 'bg-primary text-white hover:bg-primary'
            )"
            :style="{ height: settings.layout.sideItemHeight + 'px' }"
            :title="item.meta.title"
            @click="selectTopMenuItem(item.name)"
          >
            <component :is="item.meta.icon" v-if="item.meta.icon" class="h-5 w-5" />
            <span v-else class="text-[13px]">{{ item.meta.title[0] }}</span>
          </button>
        </nav>
      </el-scrollbar>
    </div>

    <!-- 二级并列纵向菜单：仅当当前一级菜单存在二级时才出现 -->
    <div
      v-if="secondLevelMenus.length"
      class="relative flex h-full flex-col shadow-sider transition-[width] duration-300 ease-in-out"
      :class="[surfaceClass, siderDarkClass]"
      :style="{ width: layoutSideWidth + 'px' }"
    >
      <el-scrollbar class="flex-1">
        <g-menu
          :items="secondLevelMenus"
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
  import { computed, ref, watchEffect } from 'vue'
  import { useRoute } from 'vue-router'
  import { storeToRefs } from 'pinia'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { resolveActiveTop, visibleItems } from '@/core/componentLibrary/menu/shared'
  import { useMenuActive, useMenuNavigation } from './composables/useMenu'
  import { useSidebarTheme } from './composables/useSidebarTheme'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import { cn, FOCUS_RING } from '@/core/componentLibrary/utils'
  import { MENU_ICON_BUTTON } from '@/core/componentLibrary/menu/variants'
  import CollapseBar from './CollapseBar.vue'

  defineOptions({ name: 'SidebarMode' })

  const { sideCollapse } = storeToRefs(useAppStore())
  const { settings } = storeToRefs(useThemeStore())
  const route = useRoute()
  const routerStore = useRouterStore()

  const topActive = ref('')
  const secondLevelMenus = ref([])

  const topLevelMenus = computed(() => visibleItems(routerStore.rootMenus))
  const { menuTheme, surfaceClass, siderDarkClass } = useSidebarTheme()
  const { activeKey, openKeys } = useMenuActive(secondLevelMenus, menuTheme)
  const { navigate } = useMenuNavigation()
  const { sideWidth: layoutSideWidth } = useSideWidth()

  // 一级与二级之间的分隔：跟随系统「卡片样式」配置
  const railSepClass = computed(() =>
    settings.value.card.mode === 'border'
      ? 'border-0 border-r border-solid border-border'
      : 'shadow-sider'
  )

  // 选择一级：有二级→显示二级并跳首个可见子项；无二级→隐藏二级面板并直接跳
  const selectTopMenuItem = (name) => {
    topActive.value = name
    const top = routerStore.rootMenus.find((i) => i.name === name)
    const children = visibleItems(top?.children)
    if (children.length) {
      secondLevelMenus.value = top.children
      navigate(children[0].name)
    } else {
      secondLevelMenus.value = []
      navigate(name)
    }
  }

  // 路由变化时定位当前所属一级；无二级时清空面板（不回落显示其它一级的子菜单）
  watchEffect(() => {
    const top = resolveActiveTop(
      routerStore.rootMenus,
      route.meta.activeName || route.name
    )
    if (top) {
      topActive.value = top.name
      secondLevelMenus.value = visibleItems(top.children).length ? top.children : []
    } else {
      topActive.value = route.name
      secondLevelMenus.value = []
    }
  })
</script>
