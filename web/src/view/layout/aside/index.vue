<template>
  <div>
    <normal-mode
      v-if="
        settings.layout.mode === 'normal' ||
        (device === 'mobile' && settings.layout.mode == 'head') ||
        (device === 'mobile' && settings.layout.mode == 'combination') ||
        (device === 'mobile' && settings.layout.mode == 'sidebar') ||
        (device === 'mobile' && settings.layout.mode == 'vertical')
      "
    />
    <head-mode v-if="settings.layout.mode === 'head' && device !== 'mobile'" />
    <combination-mode
      v-if="settings.layout.mode === 'combination' && device !== 'mobile'"
      :mode="mode"
    />
    <sidebar-mode
      v-if="settings.layout.mode === 'sidebar' && device !== 'mobile'"
    />
    <vertical-mode
      v-if="settings.layout.mode === 'vertical' && device !== 'mobile'"
    />
  </div>
</template>

<script setup>
  import NormalMode from './normalMode.vue'
  import HeadMode from './headMode.vue'
  import CombinationMode from './combinationMode.vue'
  import SidebarMode from './sidebarMode.vue'
  import VerticalMode from './verticalMode.vue'

  defineProps({
    mode: {
      type: String,
      default: 'normal'
    }
  })

  import { storeToRefs } from 'pinia'
  import { useAppStore, useThemeStore } from '@/pinia'
  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { device } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
</script>
