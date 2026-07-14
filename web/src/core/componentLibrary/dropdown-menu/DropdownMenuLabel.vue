<script setup>
import { DropdownMenuLabel } from 'reka-ui'
import { cn } from '../utils'

defineOptions({ name: 'UiDropdownMenuLabel' })

// 分组标签（非交互项），弱化文本区分于可选中的菜单项
const props = defineProps({
  class: { type: null, default: '' }
})
</script>

<template>
  <DropdownMenuLabel
    :class="cn('block cursor-default px-2.5 py-2 text-[13px] text-muted-foreground outline-none', props.class)"
  >
    <slot />
  </DropdownMenuLabel>
</template>
