<template>
  <div>
    <normal-mode
      v-if="
        config.side_mode === 'normal' ||
        (device === 'mobile' && config.side_mode == 'head') ||
        (device === 'mobile' && config.side_mode == 'combination') ||
        (device === 'mobile' && config.side_mode == 'sidebar')
      "
    />
    <head-mode v-if="config.side_mode === 'head' && device !== 'mobile'" />
    <combination-mode
      v-if="config.side_mode === 'combination' && device !== 'mobile'"
      :mode="mode"
    />
    <sidebar-mode
      v-if="config.side_mode === 'sidebar' && device !== 'mobile'"
    />
  </div>
</template>

<script setup>
  import NormalMode from './normalMode.vue'
  import HeadMode from './headMode.vue'
  import CombinationMode from './combinationMode.vue'
  import SidebarMode from './sidebarMode.vue'

  defineProps({
    mode: {
      type: String,
      default: 'normal'
    }
  })

  import { storeToRefs } from 'pinia'
  import { useAppStore } from '@/pinia'
  const appStore = useAppStore()
  const { config, device } = storeToRefs(appStore)
</script>
