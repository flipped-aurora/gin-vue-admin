<script setup>
import { DropdownMenuSeparator } from 'reka-ui'
import { cn } from '../utils'

defineOptions({ name: 'UiDropdownMenuSeparator' })

const props = defineProps({
  class: { type: null, default: '' }
})
</script>

<template>
  <DropdownMenuSeparator
    :class="cn('mx-2 my-1 h-px bg-border opacity-50', props.class)"
  />
</template>
