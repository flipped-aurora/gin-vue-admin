<template>
  <div class="mx-2 flex h-full w-[calc(100vw-600px)] items-center text-base-text">
    <g-menu
      orientation="horizontal"
      :items="menus"
      :active="activeKey"
      class="w-full"
      @select="navigate"
    />
  </div>
</template>

<script setup>
  import { computed } from 'vue'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useMenuActive, useMenuNavigation } from './composables/useMenu'

  defineOptions({ name: 'AsideHeadMode' })

  const routerStore = useRouterStore()
  const menus = computed(() => routerStore.rootMenus)
  const { activeKey } = useMenuActive(menus)
  const { navigate } = useMenuNavigation()
</script>
