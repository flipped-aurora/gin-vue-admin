<script setup>
import { SelectContent, SelectPortal, SelectViewport } from 'reka-ui'
import { cn } from '../utils'

defineOptions({ name: 'UiSelectContent' })

const props = defineProps({
  class: { type: null, default: '' },
  position: { type: String, default: 'popper', validator: (v) => ['popper', 'item-aligned'].includes(v) },
  sideOffset: { type: Number, default: 4 }
})
</script>

<template>
  <SelectPortal>
    <SelectContent
      :position="position"
      :side-offset="sideOffset"
      :class="cn(
        'relative z-[3000] min-w-[8rem] overflow-hidden rounded-md border border-border bg-container text-base-text shadow-lg',
        position === 'popper' && 'w-[var(--reka-select-trigger-width)]',
        props.class
      )"
    >
      <SelectViewport class="max-h-72 overflow-y-auto p-1">
        <slot />
      </SelectViewport>
    </SelectContent>
  </SelectPortal>
</template>
