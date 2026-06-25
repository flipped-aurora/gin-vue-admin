<script setup>
import { computed } from 'vue'
import {
  NumberFieldRoot,
  NumberFieldInput,
  NumberFieldIncrement,
  NumberFieldDecrement
} from 'reka-ui'
import { cn, FOCUS_RING_WITHIN } from '../utils'

defineOptions({ name: 'UiNumberField' })

const props = defineProps({
  min: { type: Number, default: undefined },
  max: { type: Number, default: undefined },
  step: { type: Number, default: 1 },
  disabled: { type: Boolean, default: false },
  class: { type: null, default: '' }
})

// 双向绑定统一走 defineModel，取代手写 modelValue prop + defineEmits + @update 三段样板
const modelValue = defineModel({ type: Number, default: 0 })

// 清空输入框时 reka 可能回写 null/NaN，下游会把它拼成 `${val}px` 破坏布局；
// 用可写 computed 拦掉非法值，保留上一个合法值。
const innerValue = computed({
  get: () => modelValue.value,
  set: (val) => {
    if (val == null || Number.isNaN(val)) return
    modelValue.value = val
  }
})
</script>

<template>
  <NumberFieldRoot
    v-model="innerValue"
    :min="min"
    :max="max"
    :step="step"
    :disabled="disabled"
    :class="cn(
      'inline-flex h-8 items-center overflow-hidden rounded-md border border-border bg-container text-base-text transition-colors focus-within:border-primary',
      FOCUS_RING_WITHIN,
      props.class
    )"
  >
    <NumberFieldDecrement
      aria-label="减少"
      class="flex h-full w-7 shrink-0 items-center justify-center text-muted-foreground transition-colors hover:bg-muted hover:text-primary disabled:opacity-50"
    >
      <svg-icon icon="lucide:minus" class="h-3.5 w-3.5" aria-hidden="true" />
    </NumberFieldDecrement>
    <NumberFieldInput
      class="h-full w-full min-w-0 flex-1 border-x border-border bg-transparent text-center text-[13px] text-base-text outline-none [appearance:textfield] [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none"
    />
    <NumberFieldIncrement
      aria-label="增加"
      class="flex h-full w-7 shrink-0 items-center justify-center text-muted-foreground transition-colors hover:bg-muted hover:text-primary disabled:opacity-50"
    >
      <svg-icon icon="lucide:plus" class="h-3.5 w-3.5" aria-hidden="true" />
    </NumberFieldIncrement>
  </NumberFieldRoot>
</template>
