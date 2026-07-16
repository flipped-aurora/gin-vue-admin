<template>
  <div>
    <normal-mode v-if="effectiveMode === 'normal'" />
    <head-mode v-if="effectiveMode === 'head'" />
    <combination-mode
      v-if="effectiveMode === 'combination'"
      :mode="mode"
    />
    <sidebar-mode v-if="effectiveMode === 'sidebar'" />
    <vertical-mode v-if="effectiveMode === 'vertical'" />
  </div>
</template>

<script setup>
  import NormalMode from './normalMode.vue'
  import HeadMode from './headMode.vue'
  import CombinationMode from './combinationMode.vue'
  import SidebarMode from './sidebarMode.vue'
  import VerticalMode from './verticalMode.vue'
  import { useLayoutMode } from '@/hooks/useLayoutMode'

  defineOptions({
    name: 'GvaAside'
  })

  defineProps({
    mode: {
      type: String,
      default: 'normal'
    }
  })

  const { effectiveMode } = useLayoutMode()
</script>
