<script setup>
import { SelectItem, SelectItemIndicator, SelectItemText } from 'reka-ui'
import { cn } from '../utils'

defineOptions({ name: 'UiSelectItem' })

const props = defineProps({
  value: { type: String, required: true },
  disabled: { type: Boolean, default: false },
  class: { type: null, default: '' }
})
</script>

<template>
  <SelectItem
    :value="value"
    :disabled="disabled"
    :class="cn(
      'relative flex w-full cursor-pointer select-none items-center rounded-sm py-1.5 pl-7 pr-2 text-[13px] text-base-text outline-none transition-colors data-[highlighted]:bg-muted data-[highlighted]:text-primary data-[state=checked]:font-medium data-[state=checked]:text-primary data-[disabled]:pointer-events-none data-[disabled]:opacity-50',
      props.class
    )"
  >
    <span class="absolute left-2 inline-flex items-center justify-center text-primary">
      <SelectItemIndicator>
        <svg-icon icon="lucide:check" class="h-3.5 w-3.5" />
      </SelectItemIndicator>
    </span>
    <SelectItemText>
      <slot />
    </SelectItemText>
  </SelectItem>
</template>
