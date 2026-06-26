<script setup>
import { SelectTrigger, SelectIcon } from 'reka-ui'
import { cn, FOCUS_RING } from '../utils'

defineOptions({ name: 'UiSelectTrigger' })

const props = defineProps({
  class: { type: null, default: '' }
})
</script>

<template>
  <SelectTrigger
    :class="cn(
      'flex h-8 items-center justify-between gap-2 rounded-md border border-border bg-container px-3 text-[13px] text-base-text transition-colors hover:border-primary focus-visible:border-primary disabled:cursor-not-allowed disabled:opacity-50 data-[placeholder]:text-muted-foreground',
      FOCUS_RING,
      props.class
    )"
  >
    <slot />
    <SelectIcon>
      <svg-icon icon="lucide:chevron-down" class="h-4 w-4 shrink-0 opacity-60" />
    </SelectIcon>
  </SelectTrigger>
</template>
