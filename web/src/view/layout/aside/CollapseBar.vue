<template>
  <button
    v-if="settings.header.collapseButton.visible"
    type="button"
    :aria-label="collapsed ? '展开侧边栏' : '收起侧边栏'"
    :aria-expanded="!collapsed"
    :class="cn(
      'flex h-10 w-full flex-shrink-0 cursor-pointer appearance-none items-center justify-center border-0 border-t border-solid border-border bg-transparent text-base-text transition-colors hover:bg-muted',
      FOCUS_RING
    )"
    @click="appStore.toggleSideCollapse()"
  >
    <svg-icon
      :icon="collapsed ? 'lucide:chevrons-right' : 'lucide:chevrons-left'"
      class="h-4 w-4"
    />
  </button>
</template>

<script setup>
  import { storeToRefs } from 'pinia'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { cn, FOCUS_RING } from '@/core/componentLibrary/utils'

  defineOptions({ name: 'AsideCollapseBar' })

  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { sideCollapse: collapsed } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
</script>
